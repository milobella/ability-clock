package main

import (
	"fmt"
	"github.com/milobella/ability-sdk-go/pkg/ability"
	"github.com/tkuchiki/go-timezone"
	"time"
)

const defaultTimezoneLocation = "Europe/Paris"

// fun main()
func main() {

	// Read configuration
	conf := ability.ReadConfiguration()

	// Initialize server
	server := ability.NewServer("Clock Ability", conf.Server.Port)
	server.RegisterIntentRule("GET_TIME", GetTimeIntentHandler)
	server.Serve()
}

func GetTimeIntentHandler(request *ability.Request, resp *ability.Response) {
	tz := timezone.New()
	location := getTimezoneLocation(request)
	now, err := tz.FixedTimezone(time.Now(), location)
	if err != nil {
		resp.Nlg = ability.NLG{
			Sentence: "Error",
		}
		return
	}
	timeVal := fmt.Sprintf("%d h %d", now.Hour(), now.Minute())
	resp.Nlg = ability.NLG{
		Sentence: "It is {{time}}",
		Params: []ability.NLGParam{{
			Name:  "time",
			Value: timeVal,
			Type:  "time",
		}}}
}

func getTimezoneLocation(request *ability.Request) string {
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
