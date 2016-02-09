package config
import (
	"io/ioutil"
	"encoding/json"
	"github.com/ArthurHlt/microcos/logger"
	"github.com/kardianos/osext"
	"github.com/ArthurHlt/gominlog"
)
const CONFIG_PATH = "/config.json"
type Admin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
type Db struct {
	Driver     string `json:"driver"`
	DataSource string `json:"data_source"`
}
type ConfigEureka struct {
	CertFile    string   `json:"cert_file,omitempty"`
	KeyFile     string   `json:"key_file,omitempty"`
	CaCertFiles []string `json:"ca_cert_files,omitempty"`
	Timeout     int      `json:"timeout,omitempty"`
	Consistency string   `json:"consistency,omitempty"`
}
type Eureka struct {
	Config   ConfigEureka  `json:"config,omitempty"`
	Machines []string      `json:"machines"`
}
type ConfigCachet struct {
	APIURL      string `json:"api_url"`
	APIToken    string `json:"api_token"`
	InsecureAPI bool   `json:"insecure_api"`
}
type Instance struct {
	Name string       `json:"name"`
}
type ConfigMsc struct {
	Host     string       `json:"host,omitempty"`
	Admin    Admin        `json:"admin,omitempty"`
	Db       Db           `json:"db"`
	Eureka   Eureka       `json:"eureka"`
	Instance Instance     `json:"instance"`
	Cachet   ConfigCachet `json:"cachet"`
}
var config *ConfigMsc
var loggerConfig *gominlog.MinLog = logger.GetMinLog()
func newConfigFromFile(file string) (*ConfigMsc, error) {
	var configFile *ConfigMsc = &ConfigMsc{}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, configFile)
	if err != nil {
		loggerConfig.Warning("Error during unmarshalling json: %v", err.Error())
		return nil, err
	}
	return configFile, nil
}
func GetConfig() *ConfigMsc {
	if config != nil {
		return config
	}
	err := loadConfigFromFile()
	if err != nil {
		loadDefaultConfig()
		b, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			panic(err)
		}
		loggerConfig.Info("No config file found, default config has been loaded: " + string(b))
	}
	return config
}
func loadDefaultConfig() {
	config = &ConfigMsc{
		Host: ":3000",
		Admin: Admin{
			UserName: "admin",
			Password: "admin",
		},
		Db: Db{
			Driver: "sqlite3",
			DataSource: "./microcos.db",
		},
		Eureka: Eureka{
			Config: ConfigEureka{
				CertFile: "",
				KeyFile: "",
				CaCertFiles: make([]string, 0),
				Timeout: 100000,
				Consistency: "",
			},
			Machines: []string{"http://127.0.0.1:8761/eureka"},
		},
		Instance: Instance{
			Name: "default",
		},
		Cachet: ConfigCachet{
			APIToken: "",
			APIURL: "",
			InsecureAPI: false,
		},
	}
}
func loadConfigFromFile() error {
	var err error
	loggerConfig.Info("Load config from file: " + CONFIG_PATH)
	execFolder, err := osext.ExecutableFolder()
	if err != nil {
		return err
	}
	loggerConfig.Debug("Folder: " + execFolder)
	config, err = newConfigFromFile(execFolder + CONFIG_PATH)
	if err != nil {
		return err
	}
	return nil
}