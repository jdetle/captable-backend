package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jdetle/captable-backend/captable"
	"github.com/jdetle/captable-backend/config"
	"github.com/jdetle/captable-backend/dal/model"
	"github.com/jdetle/captable-backend/httputils"

	log "github.com/sirupsen/logrus"
)

func CreateChunkHandler(cfg *config.Config, captable *captable.CapTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload model.CreateCapTableRequest
		err := httputils.ValidateJSONPayload(w, r.Body, &payload)
		if err != nil {
			return
		}
	}
}

func GetChunkHandler(cfg *config.Config, captable *captable.CapTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return
		}
		log.Debugf("%#v", id)
	}
}

func UpdateChunkHandler(cfg *config.Config, captable *captable.CapTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		var payload model.UpdateCapTableRequest
		err = httputils.ValidateJSONPayload(w, r.Body, &payload)
		if err != nil {
			return
		}
		log.Debugf("%#v", id)
	}
}

func DeleteChunkHandler(cfg *config.Config, captable *captable.CapTable) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		var payload model.UpdateCapTableRequest
		err = httputils.ValidateJSONPayload(w, r.Body, &payload)
		if err != nil {
			return
		}
		log.Debugf("%#v", id)
	}
}
