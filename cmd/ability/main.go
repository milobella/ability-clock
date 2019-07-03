package main

import (
	"gitlab.milobella.com/milobella/ability-sdk-go/pkg/ability"
	"os"
	"time"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var additionalConfigPath string

var defaultLocation *time.Location

//TODO: try to put some common stuff into a separate repository
func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	additionalConfigPath = os.Getenv("ADDITIONAL_CONFIG_PATH")
	if len(additionalConfigPath) != 0 {
		viper.AddConfigPath(additionalConfigPath)
	}

	viper.AddConfigPath(".")
	viper.SetDefault("server.log-level", "info")
	viper.SetDefault("default-timezone", "Europe/Paris")

	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// TODO: read it in the config when move to viper
	logrus.SetLevel(logrus.DebugLevel)

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		logrus.Errorf("Fatal error config file: %s \n", err)
	}

	if level, err := logrus.ParseLevel(viper.GetString("server.log-level")); err == nil {
		logrus.SetLevel(level)
	} else {
		logrus.Warn("Failed to parse the log level. Keeping the logrus default level.")
	}

	logrus.Debugf("Configuration -> %+v", viper.AllSettings())

	defaultLocation, err = time.LoadLocation(viper.GetString("default-timezone"))
	if err != nil {
		logrus.Errorf("Fatal error trying to parse the default-timezone: %s \n", err)
	}
}

// fun main()
func main() {

	// Initialize server
	server := ability.NewServer("clock", viper.GetInt("server.port"))
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

func extractTimeLocation(request *ability.Request) *time.Location {
	timezone, ok := request.Device.State["timezone"]
	if !ok {
		return defaultLocation
	}
	timezoneStr, ok := timezone.(string)
	if !ok {
		return defaultLocation
	}
	loc, err := time.LoadLocation(timezoneStr)
	if err != nil {
		return defaultLocation
	}
	return loc
}
