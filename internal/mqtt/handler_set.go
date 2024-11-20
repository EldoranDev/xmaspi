package mqtt

import (
	"github.com/EldoranDev/xmaspi/v3/internal/rendering"
)

func (h *handler) handleSet(set State) (bool, error) {
	if set.State == "OFF" {
		h.manager.ClearRenderer()
		h.manager.Clear()
	} else {
		if set.Effect != nil {
			renderer, err := rendering.GetRenderer(*set.Effect)
			if err != nil {
				return false, err
			}

			h.manager.Render(h.ctx, renderer)
		} else {
			renderer, err := rendering.GetRenderer(rendering.StaticRenderer)
			if err != nil {
				return false, err
			}

			h.manager.Render(h.ctx, renderer)
		}

		if set.Brightness != nil {
			h.manager.SetBrightness(*set.Brightness)
		}

		if set.Color != nil {
			h.manager.SetColor(*set.Color)
		}
	}

	// Send state Update
	h.stateUpdates <- h.GetState()

	return true, nil
}
