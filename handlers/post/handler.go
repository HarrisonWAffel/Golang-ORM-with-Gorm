package post

import (
	"encoding/json"
	"github.com/HarrisonWAffel/dbTrain/domain/posts"
	"github.com/HarrisonWAffel/dbTrain/domain/user"
	"github.com/HarrisonWAffel/dbTrain/util"
	"log"
	"net/http"
)

type Handler struct {
	SrvCtx *util.AppCtx
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var payload user.User
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		posts, err := h.SrvCtx.UserPostsService.GetUserPostsByUserId(payload)
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
		var payload posts.Post
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		email := r.URL.Query().Get("email")
		if email == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.SrvCtx.UserPostsService.CreateNewPost(payload, email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Println("post created")

	case http.MethodPut:

		var payload posts.Post
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.SrvCtx.PostsService.UpdatePost(payload)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case http.MethodDelete:
		var payload posts.Post
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.SrvCtx.UserPostsService.RemoveUserPostByPostId(payload)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
