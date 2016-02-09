package eureka_request
import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"fmt"
	"github.com/ArthurHlt/microcos/eureka_client"
)

type EurekaRequestInterface interface {
	SetRoutes(martini.Router)
}

type EurekaRequest struct {
	server       *martini.ClassicMartini
	eurekaClient *eureka_client.EurekaClient
}

func (this *EurekaRequest) showError(err error, r render.Render) {
	r.Text(500, fmt.Sprintf("%v", err))
}

func (this *EurekaRequest) getAppId(originalAppId string) string {
	return this.eurekaClient.GroupName + "-" + originalAppId
}