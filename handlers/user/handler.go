package user

import (
	"encoding/json"
	"github.com/HarrisonWAffel/dbTrain/domain/user"
	"github.com/HarrisonWAffel/dbTrain/util"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Handler struct {
	SrvCtx *util.AppCtx
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Query().Get("by") == "email" {
			var payload struct {
				Email string `json:"email"`
			}

			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			user, err := h.SrvCtx.UserService.GetUserByEmail(payload.Email)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write(user.ToJSON())

		} else if r.URL.Query().Get("by") == "id" {
			var payload struct {
				uuid.UUID `json:"user_id"`
			}

			if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			user, err := h.SrvCtx.UserService.GetUserById(payload.UUID)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write(user.ToJSON())
		}

	case http.MethodPost:
		var payload user.User
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.SrvCtx.UserService.SaveUser(payload); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	case http.MethodDelete:
		var payload struct {
			uuid.UUID `json:"user_id"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user := user.User{}
		user.ID = payload.UUID

		err := h.SrvCtx.UserService.DeleteUser(user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case http.MethodPut:
		var payload user.User
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.SrvCtx.UserService.UpdateUser(payload); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
