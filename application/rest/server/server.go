package server

import (
	"net/http"

	"github.com/bersennaidoo/farmstyle/application/rest/handlers"
	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

var PROBS_URL string = "/probs"
var BASE_URL string = "/"

const BASE_PATH string = "/v1"

type HttpServer struct {
	router         *mux.Router
	userHandler    *handlers.UserHandler
	reviewshandler *handlers.ReviewsHandler
	config         *viper.Viper
	log            *golog.Logger
}

func New(userHandler *handlers.UserHandler, reviewshandler *handlers.ReviewsHandler, config *viper.Viper, log *golog.Logger) *HttpServer {
	return &HttpServer{
		router:         mux.NewRouter(),
		userHandler:    userHandler,
		reviewshandler: reviewshandler,
		config:         config,
		log:            log,
	}
}

func (s *HttpServer) InitRouter() {

	api := s.router.PathPrefix(BASE_PATH).Subrouter()
	//api.Use(mid.ValidateRequestMiddleware)

	api.HandleFunc("/users", s.userHandler.AddUser).Methods(http.MethodPost)
	//api.HandleFunc("/tokens", s.userHandler.CreateToken).Methods(http.MethodPost)

	api.HandleFunc("/reviews", s.reviewshandler.GetReviews).Methods(http.MethodGet)
	api.HandleFunc("/reviews", s.reviewshandler.AddReview).Methods(http.MethodPost)

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
