package server
import (
	"github.com/go-martini/martini"
	"github.com/ArthurHlt/go-eureka-client/eureka"
	"github.com/martini-contrib/render"
	"fmt"
)

type RequestInterface interface {
	SetRoutes(martini.Router)
}

type Request struct {
	Server       *martini.ClassicMartini
	EurekaClient *eureka.Client
}

func (this *Request) showError(err error, r render.Render) {
	r.Text(500, fmt.Sprintf("%v", err))
}