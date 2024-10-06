package transport

import (
	"encoding/json"
	"github.com/Sskrill/TestTaskMusic/internal/domain"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (h *Handler) editSong(w http.ResponseWriter, r *http.Request) {
	go h.logger.PrintInfo("(Transport/HTTP) Editing song")
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.PrintError("(Transport/HTTP) Convert string to int | error:%s | handler:editSong", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.PrintError("((Transport/HTTP) Req Body Read | error:%s | handler:editSong", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	var song domain.UpdateSong
	err = json.Unmarshal(data, &song)
	if err != nil {
		h.logger.PrintError("((Transport/HTTP) Req Body Unmarshal | error:%s | handler:editSong", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	err = h.service.EditSong(id, song)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *Handler) addSong(w http.ResponseWriter, r *http.Request) {
	go h.logger.PrintInfo("(Transport/HTTP) Adding Song")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.PrintError("((Transport/HTTP) Req Body Read | error:%s | handler:addSong", err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	var song domain.Song
	err = json.Unmarshal(data, &song)
	if err != nil {
		h.logger.PrintError("(Transport/HTTP) Req Body Unmarshal | error:%s | handler:addSong", err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	err = h.service.AddSong(song)
	if err != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *Handler) getDetailsSong(w http.ResponseWriter, r *http.Request) {
	go h.logger.PrintInfo("(Transport/HTTP) Getting Details of Song")
	vars := mux.Vars(r)
	songName := vars["song_name"]
	performerName := vars["performer_name"]

	detailsSong, err := h.service.GetSongDetails(songName, performerName)
	if err != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	resp, err := json.Marshal(detailsSong)
	if err != nil {
		h.logger.PrintError("(Transport/HTTP) Marshal to json | error:%s handler:getDetailSong", err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
func (h *Handler) deleteSong(w http.ResponseWriter, r *http.Request) {
	go h.logger.PrintInfo("(Transport/HTTP) Deleting song")
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.PrintError("(Transport/HTTP) Convert string to int | error:%s", err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	err = h.service.DeleteSong(id)
	if err != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	w.WriteHeader(http.StatusOK)
}
func (h *Handler) getSongText(w http.ResponseWriter, r *http.Request) {
	go h.logger.PrintInfo("(Transport/HTTP) Getting song text")
	vars := mux.Vars(r)
	songName := vars["song_name"]
	performerName := vars["performer_name"]

	text, err := h.service.GetSongText(songName, performerName)
	if err != nil {

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	resp, err := json.MarshalIndent(text, "", "  ")
	if err != nil {
		go h.logger.PrintError("(Transport/HTTP) Marshal to jsonIndent | error:%s | handler:getSongText", err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
func (h *Handler) getSongsByFilters(w http.ResponseWriter, r *http.Request) {
	go h.logger.PrintInfo("(Transport/HTTP) Getting songs by filters")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.PrintError("((Transport/HTTP) Req Body Read | error:%s | handler:getSongsByFilters", err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	var songFilters *domain.FiltersForSong
	err = json.Unmarshal(data, &songFilters)
	if err != nil {
		h.logger.PrintError("(Transport/HTTP) Req Body Unmarshal | error:%s | handler:getSongsByFilters", err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	songs, err := h.service.GetSongsByFilters(songFilters)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	resp, err := json.Marshal(songs)
	if err != nil {
		h.logger.PrintError("(Transport/HTTP) Marshal to json | error:%s | handler:getSongsByFilters", err.Error())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(domain.Error{Msg: err.Error()})
		w.Write(resp)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
