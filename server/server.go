package server

import (
	"github.com/go-martini/martini"
	"github.com/ArthurHlt/microcos/server/eureka_request"
	"github.com/ArthurHlt/microcos/eureka_client"
	"github.com/martini-contrib/cors"
	"github.com/martini-contrib/render"
	"github.com/ArthurHlt/microcos/server/jobs_request"
	"github.com/martini-contrib/auth"
	"github.com/ArthurHlt/microcos/config"
	"net"
	"github.com/ArthurHlt/gominlog"
	"github.com/ArthurHlt/microcos/logger"
)
var loggerServer *gominlog.MinLog = logger.GetMinLog()
type Server struct {
	Martini      *martini.ClassicMartini
	EurekaClient *eureka_client.EurekaClient
}

func NewServer() *Server {
	return &Server{
		Martini: martini.Classic(),
		EurekaClient: eureka_client.GetEurekaClient(),
	}
}

func (this *Server) Run(withRandomPort ...bool) {
	this.Martini.Use(cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "GET", "OPTIONS", "POST"},
		AllowCredentials: true,
	}))
	this.Martini.Use(render.Renderer())

	this.Martini.Group("/eureka", func(r martini.Router) {
		for _, request := range this.getRequestsEureka() {
			request.SetRoutes(r)
		}
	})
	this.registerJobsRoutes()
	if len(withRandomPort) > 0 && withRandomPort[0] {
		listener, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			loggerServer.Severe("Error when getting a free random port: %v", err.Error())
		}
		host := listener.Addr().String()
		listener.Close()
		this.Martini.RunOnAddr(host)
	}else if config.GetConfig().Host != "" {
		this.Martini.RunOnAddr(config.GetConfig().Host)
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
