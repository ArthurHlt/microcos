package eureka_request

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/ArthurHlt/microcos/eureka_client"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"strings"
)

type EurekaRequestGet struct {
	EurekaRequest
}

func NewEurekaRequestGet(server *martini.ClassicMartini, eurekaClient *eureka_client.EurekaClient) *EurekaRequestGet {
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
	apps = this.filteringApplications(apps)
	this.request(apps, r, req)
}

func (this *EurekaRequestGet) filteringApplications(applications *eureka.Applications) *eureka.Applications {
	filteredApplications := &eureka.Applications{
		AppsHashcode: applications.AppsHashcode,
		VersionsDelta: applications.VersionsDelta,
		Applications: make([]eureka.Application, 0),
	}
	for _, application := range applications.Applications {
		if !strings.HasPrefix(application.Name, this.eurekaClient.GroupName + "-") {
			continue
		}
		application.Name = this.filteringAppName(application.Name)
		filteredApplications.Applications = append(filteredApplications.Applications, application)
	}
	return filteredApplications
}
func (this *EurekaRequestGet) filteringAppName(appId string) {
	return strings.TrimPrefix(appId, this.eurekaClient.GroupName + "-")
}
func (this *EurekaRequestGet) requestApplication(r render.Render, req *http.Request, params martini.Params) {
	var err error
	app, err := this.eurekaClient.GetApplication(this.getAppId(params["appId"]))
	if err != nil {
		this.showError(err, r)
		return
	}
	app.Name = this.filteringAppName(app.Name)
	this.request(app, r, req)
}

func (this *EurekaRequestGet) requestInstance(r render.Render, req *http.Request, params martini.Params) {
	var err error
	instance, err := this.eurekaClient.GetInstance(this.getAppId(params["appId"]), params["instanceId"])
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