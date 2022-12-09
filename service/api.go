// api отвечает за обработку запроса, вызов сервиса и подготовку его ответа

package service

import (
	"errors"
	"io"
	"net/http"

	"github.com/go-logr/logr"
	jsoniter "github.com/json-iterator/go"
)

// Package debug level
var DL = 1

type APIService struct {
	userHeader string
	store      Storage
}

func NewAPIService(userHeader string, store Storage) APIService {
	return APIService{userHeader, store}
}

func (srv APIService) SetupRoutes(mux *http.ServeMux, privPrefix string) {
	mux.HandleFunc("/api/item", srv.Item)
	mux.HandleFunc(privPrefix + "api/new", srv.ItemCreate)
	mux.HandleFunc(privPrefix + "api/items", srv.Items)
	mux.HandleFunc(privPrefix + "api/stat", srv.Stats)
}

/*
Вариант сбора статистики в реальном времени

получаем события
new - пристегнуть к автору, обновить стату
show - удалить у автора, сохранить время показа, обновить стату
expire? - удалить у автора, сохранить статус
timeout - expire? обновить стату
*/

func (srv APIService) ItemCreate(w http.ResponseWriter, r *http.Request) {
	log := logr.FromContextOrDiscard(r.Context())
	// read body
	// TODO: why not use // json.NewDecoder(resp.Body).Decode(&result)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.V(DL).Info("Error reading body", "error", err.Error())
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	// parse request json
	var req NewItemRequest
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.V(DL).Info("JSON parse error", "error", err.Error())
		http.Error(w, "JSON parse error", http.StatusBadRequest)
		return
	}
	// Fetch X-Username
	user := r.Header.Get(srv.userHeader)
	if user == "" {
		log.V(DL).Info("Username must be set")
		http.Error(w, "Username must be set", http.StatusUnauthorized)
		return
	}
	id, err := srv.store.SetMeta(user, req)
	response(w, log, id, err)
}

func (srv APIService) Item(w http.ResponseWriter, r *http.Request) {
	log := logr.FromContextOrDiscard(r.Context())
	id := r.URL.Query().Get("id")
	if id == "" {
		log.V(DL).Info("ID must be set")
		http.Error(w, "ID must be set", http.StatusBadRequest)
		return
	}
	if r.Method == "POST" {
		resp, err := srv.store.GetData(id)
		response(w, log, resp, err)
	} else {
		resp, err := srv.store.GetMeta(id)
		response(w, log, resp, err)
	}
}

func (srv APIService) Items(w http.ResponseWriter, r *http.Request) {
	log := logr.FromContextOrDiscard(r.Context())
	// Fetch X-Username
	user := r.Header.Get(srv.userHeader)
	if user == "" {
		log.V(DL).Info("Username must be set")
		http.Error(w, "Username must be set", http.StatusUnauthorized)
		return
	}
	resp, err := srv.store.Items(user)
	response(w, log, resp, err)
}

func (srv APIService) Stats(w http.ResponseWriter, r *http.Request) {
	log := logr.FromContextOrDiscard(r.Context())
	// Fetch X-Username
	user := r.Header.Get(srv.userHeader)
	if user == "" {
		log.V(DL).Info("Username must be set")
		http.Error(w, "Username must be set", http.StatusUnauthorized)
		return
	}
	resp, err := srv.store.Stats(user)
	response(w, log, resp, err)
}

func response(w http.ResponseWriter, log logr.Logger, data interface{}, err error) {
	if err != nil {
		// if ErrNotFound - return 404
		if errors.Is(err, ErrNotFound) {
			http.Error(w, "Item not found", http.StatusNotFound)
		} else {
			log.V(DL).Info("Call error", "error", err.Error())
			http.Error(w, "Call error", http.StatusInternalServerError)
		}
		return
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.V(DL).Info("JSON marshal error", "error", err.Error())
		http.Error(w, "JSON marshal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
