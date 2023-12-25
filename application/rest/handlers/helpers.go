package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bersennaidoo/farmstyle/foundation/emsg"
)

type MiddlewareFn func(http.ResponseWriter, *http.Request)

func ErrorResponse(prob *emsg.ProblemJson) MiddlewareFn {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(prob.Error())
		prob := emsg.Absolutify(prob, "/probs", "/")
		writeJson(prob.Status, prob)(w, r)
	}
}

func writeJson(status int, msg interface{}) MiddlewareFn {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("%s", msg)
		w.Header().Add("Content-type", "application/json")
		msgBytes, _ := json.Marshal(msg)
		w.WriteHeader(status)
		w.Write([]byte(msgBytes))
	}
}
