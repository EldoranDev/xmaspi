//go:build pi
// +build pi

package proto

import (
	"context"
	"github.com/EldoranDev/xmaspi/v2/internal/led"
)

func NewServer(ctrl led.Controller) XmasPIServer {
	return &xmaspiServer{
		controller: ctrl,
	}
}

type xmaspiServer struct {
	controller led.Controller
}

func (x *xmaspiServer) GetControllerInfo(ctx context.Context, request *ControllerInfoRequest) (*ControllerInfoResponse, error) {
	return &ControllerInfoResponse{
		LedCount: int64(x.controller.LedCount()),
	}, nil
}

func (x *xmaspiServer) SetStatic(ctx context.Context, request *SetStaticRequest) (*SetStaticResponse, error) {
	x.controller.Fill(
		request.Color,
		true,
	)

	return &SetStaticResponse{}, nil
}

func (x *xmaspiServer) SetLed(ctx context.Context, request *SetLedRequest) (*SetLedResponse, error) {
	x.controller.SetLed(
		int(request.Led),
		request.Color,
		true,
	)

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
