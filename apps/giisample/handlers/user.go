package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type (
	// artistService specifies the interface for the artist service needed by artistResource.
	// 在多层架构中 可以同时适配dao层和service层
	artistService interface {
		Get(id int) (*models.Artist, error)
		Query(offset, limit int) ([]models.Artist, error)
		Count() (int, error)
		Create(model *models.Artist) (*models.Artist, error)
		Update(id int, model *models.Artist) (*models.Artist, error)
		Delete(id int) (*models.Artist, error)
	}

	// artistResource defines the handlers for the CRUD APIs.
	artistResource struct {
		service artistService
	}
)

// ServeArtist sets up the routing of artist endpoints and the corresponding handlers.
func ServeArtistResource(r *mux.Router, service artistService) {
	h := &artistResource{service}
	rg.Get("/artists/<id>", r.get)
	rg.Get("/artists", r.query)
	rg.Post("/artists", r.create)
	rg.Put("/artists/<id>", r.update)
	rg.Delete("/artists/<id>", r.delete)

	// r := mux.NewRouter()
	// AccessTraces
	r.HandleFunc("access-trace/{id:[0-9]+}", h.get).Methods("GET")        // Get model by id
	r.HandleFunc("access-trace", h.query).Methods("GET")                  // list models
	r.HandleFunc("access-trace", h.create).Methods("POST")                // create model
	r.HandleFunc("access-trace/{id}", h.update).Methods("PUT", "POST")    // update model
	r.HandleFunc("access-trace/{id}", h.delete).Methods("DELETE", "POST") // delete model
}

func (r *artistResource) get() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("get user!" + mux.Vars(request)["id"]))
	}
}

func (r *artistResource) query() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (r *artistResource) create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (r *artistResource) update() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}

func (r *artistResource) delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

	}
}
