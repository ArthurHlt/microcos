package eureka_request

import (
	"github.com/go-martini/martini"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"net/http"
	"strings"
	"github.com/martini-contrib/render"
)

type EurekaRequestDelete struct {
	EurekaRequest
}

func NewEurekaRequestDelete(server *martini.ClassicMartini, eurekaClient *eureka.Client) *EurekaRequestDelete {
	eurekaRequest := &EurekaRequestDelete{}
	eurekaRequest.eurekaClient = eurekaClient
	eurekaRequest.server = server
	return eurekaRequest
}

func (this *EurekaRequestDelete) requestUnregisterInstance(r render.Render, resp http.ResponseWriter, req *http.Request, params martini.Params) {
	values := []string{"apps", params["appId"], params["instanceId"]}
	path := strings.Join(values, "/")
	clientResp, err := this.eurekaClient.Delete(path)
	if err != nil {
		this.showError(err, r)
		return
	}
	resp.WriteHeader(clientResp.StatusCode)
	resp.Write(clientResp.Body)

}

func (this *EurekaRequestDelete) SetRoutes(r martini.Router) {
	r.Delete("/apps/:appId/:instanceId", this.requestUnregisterInstance)
}