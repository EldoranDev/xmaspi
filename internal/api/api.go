package api

import (
	"fmt"
	"github.com/EldoranDev/xmaspi/v2/internal/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var _ ApiHandler = (*apiHandler)(nil)

type ApiHandler interface {
	GetLeds(w http.ResponseWriter, r *http.Request)

	SetAnimation(w http.ResponseWriter, r *http.Request)
	GetAnimations(w http.ResponseWriter, r *http.Request)

	SetStatic(w http.ResponseWriter, r *http.Request)
	GetStatics(w http.ResponseWriter, r *http.Request)

	Close() error
}

type apiHandler struct {
	conn   *grpc.ClientConn
	client proto.XmasPIClient
}

func NewApiHandler() ApiHandler {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", viper.GetString("grpc.host"), viper.GetInt("grpc.port")),
		opts...,
	)

	if err != nil {
		log.Fatalf("Error connecting: %d", err)
	}

	return &apiHandler{
		conn:   conn,
		client: proto.NewXmasPIClient(conn),
	}
}

func (a *apiHandler) Close() error {
	return a.conn.Close()
}
