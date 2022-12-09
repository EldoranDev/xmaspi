package api

import (
	"context"
	"encoding/json"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"log"
	"net/http"
)

type ledsResponse struct {
	Leds int `json:"leds"`
}

func (a *apiHandler) GetLeds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	response, err := a.client.GetLedCount(context.Background(), &proto.GetLedCountRequest{})

	if err != nil {
		log.Printf("Error requesting led count: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ledsResponse{Leds: int(response.Count)})
}
