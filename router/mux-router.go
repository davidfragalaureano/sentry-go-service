package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type muxRouter struct{}

var (
	router = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(uri, f).Methods(http.MethodGet)
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	router.HandleFunc(uri, f).Methods(http.MethodPost)
}

func (*muxRouter) loadMiddlewares() {
	router.Use(HTTPLogRequest)
}
func (m *muxRouter) SERVE(port string) {
	log.Printf("HTTP server running on port %v", port)
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})
	m.loadMiddlewares()
	http.ListenAndServe(":"+port, corsOpts.Handler(router))
}
