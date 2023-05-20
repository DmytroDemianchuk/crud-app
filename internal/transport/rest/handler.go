package rest

import (
	"context"
	"errors"
	"strconv"

	"net/http"

	"github.com/dmytrodemianchuk/crud-app/internal/domain"
	"github.com/gorilla/mux"
)

type Musics interface {
	Create(ctx context.Context, music domain.Music) error
	GetById(ctx context.Context, id int64) (domain.Music, error)
	GetAll(ctx context.Context) ([]domain.Music, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, inp domain.UpdateMusicInput) error
}

type Handler struct {
	musicsService Musics
}

func NewHandler(musics Musics) *Handler {
	return &Handler{
		musicsService: musics,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	musics := r.PathPrefix("/musics").Subrouter()
	{
		musics.HandleFunc("", h.createMusic).Methods(http.MethodPost)
		musics.HandleFunc("", h.getAllMusics).Methods(http.MethodGet)
		musics.HandleFunc("/{id:[0-9]+}", h.getMusicById).Methods(http.MethodGet)
		musics.HandleFunc("/{id:[0-9]+}", h.deleteMusic).Methods(http.MethodDelete)
		musics.HandleFunc("/{id:[0-9]+}", h.updateMusic).Methods(http.MethodPut)
	}

	return r

}

func getIdFromRequest(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return id, nil
}
