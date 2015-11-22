package eureka_client
import (
	"github.com/ArthurHlt/microserv-helper/config"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/ArthurHlt/microserv-helper/logger"
	"time"
)
var eurekaClient *eureka.Client
func GetEurekaClient() *eureka.Client {
	if eurekaClient != nil {
		return eurekaClient
	}
	var client *eureka.Client
	var err error
	conf := config.GetConfig()
	if conf.Eureka.Config.CertFile == "" || conf.Eureka.Config.KeyFile == "" {
		client = eureka.NewClient(conf.Eureka.Machines)
	}else {
		client, err = eureka.NewTLSClient(
			conf.Eureka.Machines,
			conf.Eureka.Config.CertFile,
			conf.Eureka.Config.KeyFile,
			conf.Eureka.Config.CaCertFiles,
		)
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