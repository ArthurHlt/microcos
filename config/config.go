package config
import (
	"io/ioutil"
	"encoding/json"
	"github.com/ArthurHlt/microserv-helper/logger"
	"github.com/kardianos/osext"
	"github.com/ArthurHlt/gominlog"
)
const CONFIG_PATH = "/config.json"

type ConfigMsh struct {
	Host     string `json="host,omitempty"`
	Admin    struct {
				 UserName string `json="username"`
				 Password string `json="password"`
			 } `json="admin,omitempty"`
	Db       struct {
				 Driver     string `json:"driver"`
				 DataSource string `json:"data_source"`
			 } `json:"db"`
	Eureka   struct {
				 Config   struct {
							  CertFile    string `json:"cert_file,omitempty"`
							  KeyFile     string `json:"key_file,omitempty"`
							  CaCertFiles []string `json:"ca_cert_files,omitempty"`
							  Timeout     int `json:"timeout,omitempty"`
							  Consistency string `json:"consistency,omitempty"`
						  } `json:"config,omitempty"`

				 Machines []string `json:"machines"`
			 } `json:"eureka"`
	Instance struct {
				 Name   string `json:"name"`
				 Cachet struct {
							APIURL      string `json:"api_url"`
							APIToken    string `json:"api_token"`
							InsecureAPI bool `json:"insecure_api"`
						} `json:"cachet"`
			 } `json:"instance"`
}
var config *ConfigMsh
var loggerConfig *gominlog.MinLog = logger.GetMinLog()
func newConfigFromFile(file string) (*ConfigMsh, error) {
	var configFile *ConfigMsh = &ConfigMsh{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, configFile)
	if err != nil {
		return nil, err
	}
	return configFile, nil
}
func GetConfig() *ConfigMsh {
	if config != nil {
		return config
	}
	err := loadConfigFromFile()
	if err != nil {
		panic(err)
	}
	return config
}
func loadConfigFromFile() error {
	var err error
	loggerConfig.Info("Load config from file: " + CONFIG_PATH)
	execFolder, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}
	config, err = newConfigFromFile(execFolder + CONFIG_PATH)
	if err != nil {
		return err
	}
	return nil
}