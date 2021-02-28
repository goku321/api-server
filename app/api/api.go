package api

import (
	"encoding/json"
	"log"
	"net/http"

	geo "github.com/goku321/geolocation/geolocation"
	"github.com/gorilla/mux"
)

// Router represents an object to deal with all http calls.
type Router struct {
	Root  *mux.Router
	store geo.GeoDataProvider
}

// New creates a new server instance to serve all the registered endpoints.
func New(dao geo.GeoDataProvider) *Router {
	r := mux.NewRouter()
	return &Router{
		Root:  r,
		store: dao,
	}
}

// Register registers all the routes.
func (s *Router) Register() {
	s.Root.HandleFunc("/geolocation/{ip}", CreateGetHandler(s.store))
}

// CreateGetHandler creates a http handler to serve GET requests.
func CreateGetHandler(dao geo.GeoDataProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ip := vars["ip"]
		var err error
		var g *geo.GeoData
		if g, err = dao.Get(ip); err == nil {
			resp, err := json.Marshal(g)
			if err != nil {
				w.WriteHeader(400)
				return
			}
			w.WriteHeader(200)
			w.Write(resp)
			return
		}
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("geolocation does not exists"))
	}
}
