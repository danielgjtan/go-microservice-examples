package main

import (
	"encoding/base64"
	"errors"
	_ "expvar"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"time"

	uuid "github.com/nu7hatch/gouuid"

	"context"
	"golang.org/x/net/trace"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	merr "github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"go-microservice-examples/api/hotel/proto"
	"go-microservice-examples/srv/auth/proto"
	"go-microservice-examples/srv/geo/proto"
	"go-microservice-examples/srv/profile/proto"
	"go-microservice-examples/srv/rate/proto"
)

const (
	BASIC_SCHEMA  string = "Basic "
	BEARER_SCHEMA string = "Bearer "
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Method())
	return l.Client.Call(ctx, req, rsp)
}

// Implements client.Wrapper as logWrapper
func logWrap(c client.Client) client.Client {
	return &logWrapper{c}
}

// Implements the server.HandlerWrapper
func logHandlerWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		fmt.Printf("[Log Wrapper] Before serving request method: %v\n", req.Method())
		err := fn(ctx, req, rsp)
		fmt.Println("[Log Wrapper] After serving request")
		return err
	}
}

type profileResults struct {
	hotels []*profile.Hotel
	err    error
}

type rateResults struct {
	ratePlans []*rate.RatePlan
	err       error
}

type Hotel struct {
	Client client.Client
}

func (s *Hotel) Rates(ctx context.Context, req *hotel.Request, rsp *hotel.Response) error {
	// tracing
	tr := trace.New("api.v1", "Hotel.Rates")
	defer tr.Finish()

	// context
	ctx = trace.NewContext(ctx, tr)

	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}

	// add a unique request id to context
	if traceID, err := uuid.NewV4(); err == nil {
		// make copy
		tmd := metadata.Metadata{}
		for k, v := range md {
			tmd[k] = v
		}

		tmd["traceID"] = traceID.String()
		tmd["fromName"] = "api.v1"
		ctx = metadata.NewContext(ctx, tmd)
	}

	// token from request headers
	token, err := getToken(md)
	if err != nil {
		return merr.Forbidden("api.hotel.rates", err.Error())
	}

	// verify token w/ auth service
	authClient := auth.NewAuthService("go.micro.srv.auth", s.Client)
	if _, err = authClient.VerifyToken(ctx, &auth.Request{AuthToken: token}); err != nil {
		return merr.Unauthorized("api.hotel.rates", "Unauthorized")
	}

	// checkin and checkout date query params
	inDate, outDate := req.InDate, req.OutDate
	if inDate == "" || outDate == "" {
		return merr.BadRequest("api.hotel.rates", "Please specify inDate/outDate params")
	}

	// finds nearby hotels
	// TODO(hw): use lat/lon from request params
	geoClient := geo.NewGeoService("go.micro.srv.geo", s.Client)
	nearby, err := geoClient.Nearby(ctx, &geo.Request{
		Lat: 51.502973,
		Lon: -0.114723,
	})
	if err != nil {
		return merr.InternalServerError("api.hotel.rates", err.Error())
	}

	// make reqeusts for profiles and rates
	profileCh := getHotelProfiles(s.Client, ctx, nearby.HotelIds)
	rateCh := getRatePlans(s.Client, ctx, nearby.HotelIds, inDate, outDate)

	// wait on profiles reply
	profileReply := <-profileCh
	if err := profileReply.err; err != nil {
		return merr.InternalServerError("api.hotel.rates", err.Error())
	}

	// wait on rates reply
	rateReply := <-rateCh
	if err := rateReply.err; err != nil {
		return merr.InternalServerError("api.hotel.rates", err.Error())
	}

	rsp.Hotels = profileReply.hotels
	rsp.RatePlans = rateReply.ratePlans
	return nil
}

func getToken(md metadata.Metadata) (string, error) {
	// Grab the raw Authoirzation header
	authHeader := md["Authorization"]
	if authHeader == "" {
		return "", errors.New("Authorization header required")
	}

	// Confirm the request is sending Basic Authentication credentials.
	if !strings.HasPrefix(authHeader, BASIC_SCHEMA) && !strings.HasPrefix(authHeader, BEARER_SCHEMA) {
		return "", errors.New("Authorization requires Basic/Bearer scheme")
	}

	// Get the token from the request header
	// The first six characters are skipped - e.g. "Basic ".
	if strings.HasPrefix(authHeader, BASIC_SCHEMA) {
		str, err := base64.StdEncoding.DecodeString(authHeader[len(BASIC_SCHEMA):])
		if err != nil {
			return "", errors.New("Base64 encoding issue")
		}
		creds := strings.Split(string(str), ":")
		return creds[0], nil
	}

	return authHeader[len(BEARER_SCHEMA):], nil
}

func getRatePlans(c client.Client, ctx context.Context, hotelIDs []string, inDate string, outDate string) chan rateResults {
	rateClient := rate.NewRateService("go.micro.srv.rate", c)
	ch := make(chan rateResults, 1)

	go func() {
		res, err := rateClient.GetRates(ctx, &rate.Request{
			HotelIds: hotelIDs,
			InDate:   inDate,
			OutDate:  outDate,
		})
		ch <- rateResults{res.RatePlans, err}
	}()

	return ch
}

func getHotelProfiles(c client.Client, ctx context.Context, hotelIDs []string) chan profileResults {
	profileClient := profile.NewProfileService("go.micro.srv.profile", c)
	ch := make(chan profileResults, 1)

	go func() {
		res, err := profileClient.GetProfiles(ctx, &profile.Request{
			HotelIds: hotelIDs,
			Locale:   "en",
		})
		ch <- profileResults{res.Hotels, err}
	}()

	return ch
}

func main() {
	// trace library patched for demo purposes.
	// https://github.com/golang/net/blob/master/trace/trace.go#L94
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	service := micro.NewService(
		micro.Client(client.NewClient(client.Wrap(logWrap))),
		micro.Server(server.NewServer(server.WrapHandler(logHandlerWrapper))),
		micro.Name("go.micro.api.hotel"),
		micro.RegisterTTL(time.Second*3),
		micro.RegisterInterval(time.Second*3),
	)

	service.Init()
	hotel.RegisterHotelHandler(service.Server(), &Hotel{service.Client()})
	service.Run()
}
