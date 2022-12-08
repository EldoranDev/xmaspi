//go:build pi
// +build pi

package proto

import (
	"context"
	"github.com/EldoranDev/xmaspi/v2/internal/display/animations"
	"github.com/EldoranDev/xmaspi/v2/internal/display/statics"
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

func (x *xmaspiServer) GetLedCount(ctx context.Context, request *GetLedCountRequest) (*GetLedCountResponse, error) {
	return &GetLedCountResponse{Count: uint32(x.controller.LedCount())}, nil
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
	static, err := statics.GetStatic(request.Name)

	if err != nil {
		return nil, err
	}

	x.controller.RenderStatic(static)

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

func (x *xmaspiServer) GetAnimations(ctx context.Context, request *GetAnimationsRequest) (*GetAnimationsResponse, error) {
	response := &GetAnimationsResponse{}

	anims := make([]*GetAnimationsResponse_Animation, 0)

	i := 0
	for name, anim := range animations.GetAll() {
		res := &GetAnimationsResponse_Animation{
			Name: name,
		}

		if desc, ok := anim.(xmaspi.AnimationWithDescription); ok {
			res.Description = desc.Description()
			res.DisplayName = desc.DisplayName()
		}

		anims = append(anims, res)

		i++
	}

	response.Animations = anims

	return response, nil
}

func (x *xmaspiServer) GetStatics(ctx context.Context, request *GetStaticsRequest) (*GetStaticsResponse, error) {
	response := &GetStaticsResponse{}

	stcs := make([]*GetStaticsResponse_Static, 0)

	i := 0
	for name, static := range statics.GetAll() {
		res := &GetStaticsResponse_Static{
			Name: name,
		}

		if desc, ok := static.(xmaspi.StaticWithDescription); ok {
			res.Description = desc.Description()
			res.DisplayName = desc.DisplayName()
		}

		stcs = append(stcs, res)

		i++
	}

	response.Statics = stcs

	return response, nil
}

func (x *xmaspiServer) mustEmbedUnimplementedXmasPIServer() {
	//TODO implement me
	panic("implement me")
}
