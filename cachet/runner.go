package cachet

import (
	"github.com/ArthurHlt/cachet-monitor/cachet"
	"github.com/ArthurHlt/microcos/logger"
	"time"
	"fmt"
)

func RunCachetMonitor() {
	componentId := 0
	strictTls := true
	cachet.LoadEmptyConfig()

	cachet.Config.APIUrl = "https://demo.cachethq.io/api"
	cachet.Config.APIToken = "9yMHsdioQosnyVK4iCVR"
	cachet.Config.InsecureAPI = false
	cachet.Config.Monitors = []*cachet.Monitor{
		&cachet.Monitor{
			Name: "nodegear frontend",
			URL: "https://nodegear.io/ping",
			MetricID: 0,
			Threshold: float32(80),
			ComponentID: &componentId,
			ExpectedStatusCode: 200,
			StrictTLS: &strictTls,
		},
	}
	fmt.Println(cachet.Config.SystemName)
	cachet.Logger = logger.GetMinLog().GetLogger()
	log := logger.GetMinLog()

	log.Info("System: %s, API: %s\n", cachet.Config.SystemName, cachet.Config.APIUrl)
	log.Info("Starting %d monitors:\n", len(cachet.Config.Monitors))
	for _, mon := range cachet.Config.Monitors {
		log.Info(" %s: GET %s & Expect HTTP %d\n", mon.Name, mon.URL, mon.ExpectedStatusCode)
		if mon.MetricID > 0 {
			log.Info(" - Logs lag to metric id: %d\n", mon.MetricID)
		}
	}

	log.Info("\n")

	ticker := time.NewTicker(time.Minute)
	for range ticker.C {
		for _, mon := range cachet.Config.Monitors {
			go mon.Run()
		}
	}
}
