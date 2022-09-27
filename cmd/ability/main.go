package main

import (
	"fmt"
	"time"

	"github.com/milobella/ability-sdk-go/pkg/config"
	"github.com/milobella/ability-sdk-go/pkg/model"
	"github.com/milobella/ability-sdk-go/pkg/server"
	"github.com/milobella/ability-sdk-go/pkg/server/conditions"
	"github.com/tkuchiki/go-timezone"
)

const defaultTimezoneLocation = "Europe/Paris"

// fun main()
func main() {

	// Read configuration
	conf := config.Read()

	// Initialize server
	srv := server.New("Clock", conf.Server.Port)
	srv.Register(conditions.IfIntents("GET_TIME"), GetTimeIntentHandler)
	srv.Serve()
}

func GetTimeIntentHandler(request *model.Request, resp *model.Response) {
	tz := timezone.New()
	location := getTimezoneLocation(request)
	now, err := tz.FixedTimezone(time.Now(), location)
	if err != nil {
		resp.Nlg = model.NLG{
			Sentence: "Error",
		}
		return
	}
	timeVal := fmt.Sprintf("%d h %d", now.Hour(), now.Minute())
	resp.Nlg = model.NLG{
		Sentence: "It is {{time}}",
		Params: []model.NLGParam{{
			Name:  "time",
			Value: timeVal,
			Type:  "time",
		}}}
}

func getTimezoneLocation(request *model.Request) string {
	location, ok := request.Device.State["timezone"]
	if !ok {
		return defaultTimezoneLocation
	}
	locationStr, ok := location.(string)
	if !ok {
		return defaultTimezoneLocation
	}
	return locationStr
}
