package eureka_client
import (
	"github.com/ArthurHlt/microcos/config"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/ArthurHlt/microcos/logger"
	"time"
)
type EurekaClient struct {
	*eureka.Client
	GroupName string
}
var eurekaClient *EurekaClient
func GetEurekaClient() *EurekaClient {
	if eurekaClient != nil {
		return eurekaClient
	}
	var client *EurekaClient
	var err error
	conf := config.GetConfig()
	if conf.Eureka.Config.CertFile == "" || conf.Eureka.Config.KeyFile == "" {
		client = &EurekaClient{eureka.NewClient(conf.Eureka.Machines), conf.Instance.Name}
	}else {
		eurekaOriginalClient, err := eureka.NewTLSClient(
			conf.Eureka.Machines,
			conf.Eureka.Config.CertFile,
			conf.Eureka.Config.KeyFile,
			conf.Eureka.Config.CaCertFiles,
		)
		if err != nil {
			panic(err)
		}
		client = &EurekaClient{eurekaOriginalClient, conf.Instance.Name}
	}
	if err != nil {
		panic(err)
	}
	if conf.Eureka.Config.Consistency != "" {
		client.Config.Consistency = conf.Eureka.Config.Consistency
	}
	if conf.Eureka.Config.Timeout != 0 {
		client.Config.DialTimeout = time.Duration(conf.Eureka.Config.Timeout) * time.Second
	}
	eureka.SetLogger(logger.GetMinLog().GetLogger())
	return client
}