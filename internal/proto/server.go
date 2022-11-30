//go:build pi
// +build pi

package proto

import (
	"context"
	"fmt"
	"github.com/EldoranDev/xmaspi/v2/internal/led"
)

func NewServer() XmasPIServer {
	return &xmaspiServer{}
}

type xmaspiServer struct {
	controller led.Controller
}

func (x *xmaspiServer) ChangeMode(ctx context.Context, request *ModeChangeRequest) (*ModeChangeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (x *xmaspiServer) SetLed(ctx context.Context, request *SetLedRequest) (*SetLedResponse, error) {
	fmt.Printf("Got request: %d -> %x", request.Led, request.Color)

	return &SetLedResponse{}, nil
}

func (x *xmaspiServer) SetAnimation(ctx context.Context, request *SetAnimationRequest) (*SetAnimationResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (x *xmaspiServer) mustEmbedUnimplementedXmasPIServer() {
	//TODO implement me
	panic("implement me")
}
