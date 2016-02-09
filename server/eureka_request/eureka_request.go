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
	eurekaClient *eureka.Client
}

func (this *EurekaRequest) showError(err error, r render.Render) {
	r.Text(500, fmt.Sprintf("%v", err))
}