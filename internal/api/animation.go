package api

import (
	"context"
	"encoding/json"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"net/http"
)

func (a *apiHandler) SetAnimation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := r.Form["name"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	anim := r.Form["name"]

	_, err = a.client.SetAnimation(context.Background(), &proto.SetAnimationRequest{Name: anim[0]})
}

func (a *apiHandler) GetAnimations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	res, err := a.client.GetAnimations(context.Background(), &proto.GetAnimationsRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Animations)
}
