package config

import (
	"coffeebeans-people-backend/constants"
	"encoding/json"
	"flag"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/log"
	"os"
)

const defaultLogLevel = "info"

type Configuration struct {
	PORT           string
	BASE_URL       string
	MONGO_SERVER   string
	MONGO_DATABASE string
	LogLevel       string `json:"log_level"`
	SECRET_KEY     string
}

var (
	configuration *Configuration = nil
	configFile    *string        = nil
)

//defined all the required flags
func init() {
	configFile = flag.String(constants.FILE, constants.DEFAULT_CONFIG, constants.CONFIG_FILE_USAGE)
}

func ResetConfiguration() {
	configuration = nil
}

func LoadAppConfiguration() {

	if configuration == nil {

		flag.Parse()
		if len(*configFile) == 0 {
			log.Error("Mandatory arguments not provided for executing the App")
			StopService("Mandatory arguments not provided for executing the App")
		}
		configuration = loadConfiguration(*configFile)
	}
}

func loadConfiguration(filename string) *Configuration {
	if configuration == nil {
		configFile, err := os.Open(filename)
		defer configFile.Close()
		if err != nil {
			StopService(err.Error())
		}
		jsonParser := json.NewDecoder(configFile)
		err1 := jsonParser.Decode(&configuration)
		if err1 != nil {
			log.Error("Failed to parse configuration file")
			StopService(err1.Error())
		}
		setDefaultConfig()
	}
	return configuration
}

func GetAppConfiguration() *Configuration {
	if configuration == nil {
		log.Info("Unable to get the app configuration. Loading freshly. \t")
		LoadAppConfiguration()
	}
	return configuration
}

func StopService(message string) {
	log.Fatal(message)
	p, _ := os.FindProcess(os.Getpid())
	p.Signal(os.Kill)
}

func setDefaultConfig() {
	if configuration.LogLevel == "" {
		configuration.LogLevel = defaultLogLevel
	}
}
