package api

import (
	"context"
	"encoding/json"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"net/http"
)

func (a *apiHandler) SetStatic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := r.Form["name"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	static := r.Form["name"]

	_, err = a.client.SetStatic(context.Background(), &proto.SetStaticRequest{Name: static[0]})
}

func (a *apiHandler) GetStatics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	res, err := a.client.GetStatics(context.Background(), &proto.GetStaticsRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Statics)
}
