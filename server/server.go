package server

import (
	"github.com/go-martini/martini"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/ArthurHlt/microserv-helper/server/eureka_request"
	"github.com/ArthurHlt/microserv-helper/eureka_client"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"github.com/ArthurHlt/microserv-helper/server/jobs_request"
	"github.com/martini-contrib/auth"
	"github.com/ArthurHlt/microserv-helper/config"
	"os"
)

type Server struct {
	Martini      *martini.ClassicMartini
	EurekaClient *eureka.Client
}

func NewServer() *Server {
	return &Server{
		Martini: martini.Classic(),
		EurekaClient: eureka_client.GetEurekaClient(),
	}
}

func (this *Server) Run() {
	this.Martini.Use(render.Renderer())

	this.Martini.Group("/eureka", func(r martini.Router) {
		for _, request := range this.getRequestsEureka() {
			request.SetRoutes(r)
		}
	})
	this.registerJobsRoutes()
	this.Martini.Get("/test", func() {
		client := eureka.NewClient([]string{
			"http://127.0.0.1:3000/eureka",
		})
		instance := eureka.NewInstanceInfo("titi.org", "tutu", "12.12.12.12", 80, 30, false)
		client.UnregisterInstance(instance.App, instance.HostName)
	}, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "GET", "OPTIONS", "POST"},
		AllowCredentials: true,
	}))
	if config.GetConfig().Host != "" {
		this.Martini.RunOnAddr(config.GetConfig().Host)
	}else if os.Getenv("PORT") != "" {
		this.Martini.RunOnAddr(":" + os.Getenv("PORT"))
	}else {
		this.Martini.Run()
	}

}
func (this *Server) registerJobsRoutes() {
	if config.GetConfig().Admin.UserName == "" {
		this.Martini.Group("/jobs", func(r martini.Router) {
			jobs_request.CreateJobsRoutes(r)
		})
	}else {
		this.Martini.Group("/jobs", func(r martini.Router) {
			jobs_request.CreateJobsRoutes(r)
		}, auth.Basic(config.GetConfig().Admin.UserName, config.GetConfig().Admin.Password))
	}

}
func (this *Server) getRequestsEureka() []eureka_request.EurekaRequestInterface {
	return []eureka_request.EurekaRequestInterface{
		eureka_request.NewEurekaRequestGet(this.Martini, this.EurekaClient),
		eureka_request.NewEurekaRequestPost(this.Martini, this.EurekaClient),
		eureka_request.NewEurekaRequestPut(this.Martini, this.EurekaClient),
		eureka_request.NewEurekaRequestDelete(this.Martini, this.EurekaClient),
	}
}