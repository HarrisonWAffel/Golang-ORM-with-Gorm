package handlers

import (
	"github.com/HarrisonWAffel/dbTrain/util"
	"net/http"
)

type Handler struct {
	AppCtx *util.AppCtx
	H      func(*util.AppCtx, http.ResponseWriter, *http.Request)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.H(h.AppCtx, w, r)
}
