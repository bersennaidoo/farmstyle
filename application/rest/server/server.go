package server

import (
	"net/http"

	"github.com/bersennaidoo/farmstyle/application/rest/handlers"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

type HttpServer struct {
	router      *mux.Router
	userHandler *handlers.UserHandler
	config      *viper.Viper
	log         *golog.Logger
}

func New(userHandler *handlers.UserHandler, config *viper.Viper, log *golog.Logger) *HttpServer {
	return &HttpServer{
		router:      mux.NewRouter(),
		userHandler: userHandler,
		config:      config,
		log:         log,
	}
}

func (s *HttpServer) InitRouter() {

	fileServer := http.FileServer(http.Dir("./hci/static/"))
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	s.router.HandleFunc("/", s.snipHandler.Home).Methods("GET")
	s.router.HandleFunc("/snip/view", s.snipHandler.SnipView).Methods("GET")
	s.router.HandleFunc("/snip/create", s.snipHandler.SnipCreate).Methods("POST")
}

func (s *HttpServer) Start() {

	addr := s.config.GetString("http.http_addr")
	srv := &http.Server{
		Addr:    addr,
		Handler: s.router,
	}

	s.log.Debugf("Server Starting on :3000")

	err := srv.ListenAndServe()
	if err != nil {
		s.log.Fatal(err)
	}
}
