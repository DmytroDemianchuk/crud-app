package rest

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dmytrodemianchuk/crud-app/internal/domain"
)

func (h *Handler) createMusic(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var music domain.Music
	if err = json.Unmarshal(reqBytes, &music); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.musicsService.Create(context.TODO(), music)
	if err != nil {
		log.Println("createMusic() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getMusicById(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("getMusicById() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	music, err := h.musicsService.GetById(context.TODO(), id)
	if err != nil {
		if errors.Is(err, domain.ErrorMusicNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Println("getMusicById() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(music)
	if err != nil {
		log.Println("getMusicById() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)

}

func (h *Handler) getAllMusics(w http.ResponseWriter, r *http.Request) {
	musics, err := h.musicsService.GetAll(context.TODO())
	if err != nil {
		log.Println("getAllMusics() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(musics)
	if err != nil {
		log.Println("getAllMusics() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func (h *Handler) deleteMusic(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("deleteMusic() error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.musicsService.Delete(context.TODO(), id)
	if err != nil {
		log.Println("deleteMusic() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateMusic(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromRequest(r)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	var inp domain.UpdateMusicInput
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.musicsService.Update(context.TODO(), id, inp)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
