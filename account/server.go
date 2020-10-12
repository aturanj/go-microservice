package account

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

//NewHTTPServer creates a new server
func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {

	r := mux.NewRouter()

	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoints.CreateUser,
		decodeUserReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoints.GetUser,
		decodeEmailReq,
		encodeResponse,
	))

	return r
}

//verifiying http request/response's type for json
func commonMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}
