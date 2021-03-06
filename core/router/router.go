package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nightlegend/apigateway/core/router/private"
	"github.com/nightlegend/apigateway/core/router/public"
	"github.com/nightlegend/apigateway/middleware"
)

// Varible define to here
var (
	LisAddr string
)

func init() {
	LisAddr = os.Getenv("ADDRESS")
	if LisAddr == "" {
		LisAddr = "0.0.0.0:8080"
	}
}

// Start start application by load self-define router.
func Start(env string) {
	// running mode switcher
	switch env {
	case "development":
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Logger())
	//No Permission Validation
	public.APIRouter(router)
	//Permission Validation
	private.APIRouter(router)

	router.Run(LisAddr)
}
