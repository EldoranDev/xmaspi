//go:build pi
// +build pi

package proto

import (
	"context"
	"github.com/EldoranDev/xmaspi/v2/internal/display/animations"
	xmaspi "github.com/EldoranDev/xmaspi/v2/pkg"
)

func NewServer(ctrl xmaspi.Controller) XmasPIServer {
	return &xmaspiServer{
		controller: ctrl,
	}
}

type xmaspiServer struct {
	controller xmaspi.Controller
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
	x.controller.FillRaw(request.Color)

	return &SetStaticResponse{}, nil
}

func (x *xmaspiServer) SetLed(ctx context.Context, request *SetLedRequest) (*SetLedResponse, error) {
	x.controller.SetLedRaw(
		int(request.Led),
		request.Color,
	)

	return &SetLedResponse{}, nil
}

func (x *xmaspiServer) SetAnimation(ctx context.Context, request *SetAnimationRequest) (*SetAnimationResponse, error) {
	anim, err := animations.GetAnimation(request.Name)

	if err != nil {
		return nil, err
	}

	x.controller.StartAnimation(anim)

	return &SetAnimationResponse{}, nil
}

func (x *xmaspiServer) mustEmbedUnimplementedXmasPIServer() {
	//TODO implement me
	panic("implement me")
}
