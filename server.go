package main

import (
	"net/http"

	"github.com/davidfragalaureano/sentry/controller"
	"github.com/davidfragalaureano/sentry/router"
	"github.com/davidfragalaureano/sentry/service"
)

var (
	sentryServerURI  string                      = "https://ssd-api.jpl.nasa.gov/sentry.api"
	httpRouter       router.Router               = router.NewMuxRouter()
	sentryService    service.SentryService       = service.NewSentryService(&http.Client{}, sentryServerURI)
	sentryController controller.SentryController = controller.NewSentryController(sentryService)
)

func main() {
	// Get a greeting message and print it.
	httpRouter.GET("/hello", sentryController.Hello)
	httpRouter.GET("/spaceObjects", sentryController.GetAllSpaceObjects)
	httpRouter.GET("/spaceObjects/{objectName}", sentryController.GetObjectByName)
	httpRouter.SERVE("8080")
}
