package main

import (
	"fmt"
	"github.com/milobella/ability-sdk-go/pkg/ability"
	"github.com/sirupsen/logrus"
	"time"
)

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
	location := extractTimeLocation(request)
	now := time.Now().In(location)
	timeVal := fmt.Sprintf("%d h %d", now.Hour(), now.Minute())
	resp.Nlg = ability.NLG{
		Sentence: "It is {{time}}",
		Params: []ability.NLGParam{{
			Name:  "time",
			Value: timeVal,
			Type:  "time",
		}}}
}

func defaultLocation() *time.Location {
	if result, err := time.LoadLocation("Europe/Paris"); err != nil {
		logrus.Errorf("Fatal error trying to parse the default-timezone: %s \n", err)
		return nil
	} else {
		return result
	}
}

func extractTimeLocation(request *ability.Request) *time.Location {
	timezone, ok := request.Device.State["timezone"]
	if !ok {
		return defaultLocation()
	}
	timezoneStr, ok := timezone.(string)
	if !ok {
		return defaultLocation()
	}
	loc, err := time.LoadLocation(timezoneStr)
	if err != nil {
		return defaultLocation()
	}
	return loc
}
