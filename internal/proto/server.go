//go:build pi
// +build pi

package proto

import (
	"context"
	"github.com/EldoranDev/xmaspi/v2/internal/animations"
	"github.com/EldoranDev/xmaspi/v2/internal/controller"
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
)

func NewServer(ctrl controller.Controller) XmasPIServer {
	return &xmaspiServer{
		controller: ctrl,
	}
}

type xmaspiServer struct {
	controller controller.Controller
}

func (x *xmaspiServer) Render(ctx context.Context, request *RenderRequest) (*RenderResponse, error) {
	_ = x.controller.Render()

	return &RenderResponse{}, nil
}

func (x *xmaspiServer) GetControllerInfo(ctx context.Context, request *ControllerInfoRequest) (*ControllerInfoResponse, error) {
	return &ControllerInfoResponse{
		LedCount: int64(x.controller.LedCount()),
	}, nil
}

func (x *xmaspiServer) SetStatic(ctx context.Context, request *SetStaticRequest) (*SetStaticResponse, error) {
	x.controller.Fill(request.Color)

	return &SetStaticResponse{}, nil
}

func (x *xmaspiServer) SetLed(ctx context.Context, request *SetLedRequest) (*SetLedResponse, error) {
	x.controller.SetLed(
		int(request.Led),
		request.Color,
	)

	return &SetLedResponse{}, nil
}

func (x *xmaspiServer) SetAnimation(ctx context.Context, request *SetAnimationRequest) (*SetAnimationResponse, error) {
	x.controller.StartAnimation(
		&animations.Stars{
			Color: leds.Color{
				B: 255,
				R: 3,
				G: 200,
			},
		},
	)

	return &SetAnimationResponse{}, nil
}

func (x *xmaspiServer) mustEmbedUnimplementedXmasPIServer() {
	//TODO implement me
	panic("implement me")
}
