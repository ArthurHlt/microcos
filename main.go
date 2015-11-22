package main

import (
//"github.com/castawaylabs/cachet-monitor/cachet"
	"github.com/ArthurHlt/go-eureka-client/eureka"
//"time"
//"fmt"
//"io/ioutil"
	"fmt"
	"github.com/ArthurHlt/microserv-helper/server"
//"github.com/ArthurHlt/microserv-helper/cachet"
	"github.com/ArthurHlt/microserv-helper/db"
	"github.com/ArthurHlt/microserv-helper/eureka_client"
//"github.com/ArthurHlt/microserv-helper/logger"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//logger.SetAppName("microserv-helper")
	//go cachet.RunCachetMonitor()
	_, err := db.GetDb()
	check(err)
	client := eureka_client.GetEurekaClient()
	//d1, err := client.MarshalJSON()
	//check(err)
	//ioutil.WriteFile("config.json", d1, 0644)
	instance := eureka.NewInstanceInfo("test.com", "test", "69.172.200.235", 80, 30, false)
	/*instance.Metadata = &eureka.MetaData{
		Map: make(map[string]string),
	}
	instance.Metadata.Map["toto"] = "titi"*/
	client.RegisterInstance("toto", instance)
	fmt.Println(client.GetApplication(instance.App))
	fmt.Println(client.GetInstance(instance.App, instance.HostName))
	//applications, err := client.GetApplications()
	//fmt.Println(applications)
	//fmt.Println(err)*/

	server := server.NewServer()
	server.Run()
}