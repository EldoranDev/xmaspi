package api

import (
	"context"
	"encoding/json"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"net/http"
)

func (a *apiHandler) SetAnimation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, ok := r.Form["animation"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
	}

	anim := r.Form["animation"]

	_, err = a.client.SetAnimation(context.Background(), &proto.SetAnimationRequest{Name: anim[0]})
}

func (a *apiHandler) GetAnimations(w http.ResponseWriter, r *http.Request) {
	res, err := a.client.GetAnimations(context.Background(), &proto.GetAnimationsRequest{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(res.Animations)
}
