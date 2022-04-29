package webapp

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Router() *mux.Router {
	return mux.NewRouter()
}

// ListenAndServe listens all TCP network and passed address,
// calls Serve to handle requests on incoming connections.
func ListenAndServe(address string, r http.Handler) {
	log.Print("Listening on tcp://" + address)
	http.ListenAndServe(address, r)
}
