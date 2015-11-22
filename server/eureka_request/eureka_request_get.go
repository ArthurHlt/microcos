package eureka_request

import (
	"github.com/go-martini/martini"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/martini-contrib/render"
	"net/http"
)

type EurekaRequestGet struct {
	EurekaRequest
}

func NewEurekaRequestGet(server *martini.ClassicMartini, eurekaClient *eureka.Client) *EurekaRequestGet {
	eurekaRequest := &EurekaRequestGet{}
	eurekaRequest.eurekaClient = eurekaClient
	eurekaRequest.server = server
	return eurekaRequest
}

func (this *EurekaRequestGet) request(data interface{}, r render.Render, req *http.Request) {
	if req.URL.Query().Get("json") != "" {
		r.JSON(200, data)
	}else {
		r.XML(200, data)
	}
}
func (this *EurekaRequestGet) requestApplications(r render.Render, req *http.Request) {
	var err error
	apps, err := this.eurekaClient.GetApplications()
	if err != nil {
		this.showError(err, r)
		return
	}
	this.request(apps, r, req)
}

func (this *EurekaRequestGet) requestApplication(r render.Render, req *http.Request, params martini.Params) {
	var err error
	app, err := this.eurekaClient.GetApplication(params["appId"])
	if err != nil {
		this.showError(err, r)
		return
	}
	this.request(app, r, req)
}

func (this *EurekaRequestGet) requestInstance(r render.Render, req *http.Request, params martini.Params) {
	var err error
	instance, err := this.eurekaClient.GetInstance(params["appId"], params["instanceId"])
	if err != nil {
		this.showError(err, r)
		return
	}
	this.request(instance, r, req)
}

func (this *EurekaRequestGet) SetRoutes(r martini.Router) {
	r.Get("/apps", this.requestApplications)
	r.Get("/apps/:appId", this.requestApplication)
	r.Get("/apps/:appId/:instanceId", this.requestInstance)
}