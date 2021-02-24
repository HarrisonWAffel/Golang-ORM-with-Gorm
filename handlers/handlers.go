package handlers

import (
	"encoding/json"
	"github.com/HarrisonWAffel/dbTrain/models"
	"github.com/HarrisonWAffel/dbTrain/util"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func User(srvCtx *util.AppCtx, w http.ResponseWriter, r *http.Request) {

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

			user, err := srvCtx.UserService.GetUserByEmail(payload.Email)
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

			user, err := srvCtx.UserService.GetUserById(payload.UUID)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write(user.ToJSON())
		}

	case http.MethodPost:
		var payload models.User
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := srvCtx.UserService.SaveUser(payload); err != nil {
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

		user := models.User{}
		user.ID = payload.UUID

		err := srvCtx.UserService.DeleteUser(user)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case http.MethodPut:
		var payload models.User
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := srvCtx.UserService.UpdateUser(payload); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}


func Post(srvCtx *util.AppCtx, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var payload models.User

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		posts, err := srvCtx.UserPostsService.GetUserPostsByUserId(payload)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		j, err := json.Marshal(posts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(j)

	case http.MethodPost:
		var payload models.Post
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		email := r.URL.Query().Get("email")
		if  email == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = srvCtx.UserPostsService.CreateNewPost(payload, email)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		log.Println("post created")


	case http.MethodPut:

		var payload models.Post
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = srvCtx.PostsService.UpdatePost(payload)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case http.MethodDelete:
		var payload models.Post
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = srvCtx.UserPostsService.RemoveUserPostByPostId(payload.ID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}


		err = srvCtx.PostsService.DeletePost(payload)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}