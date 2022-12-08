package api

import (
	"context"
	"encoding/json"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"net/http"
)

func (a *apiHandler) SetStatic(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := r.Form["static"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	static := r.Form["static"]

	_, err = a.client.SetStatic(context.Background(), &proto.SetStaticRequest{Name: static[0]})
}

func (a *apiHandler) GetStatics(w http.ResponseWriter, r *http.Request) {
	res, err := a.client.GetStatics(context.Background(), &proto.GetStaticsRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Statics)
}
