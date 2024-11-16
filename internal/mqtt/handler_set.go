package mqtt

import "github.com/EldoranDev/xmaspi/v3/internal/rendering"

func (h *handler) handleSet(set State) (bool, error) {
	if set.State == "OFF" {
		h.manager.ClearRenderer()
	} else {
		if set.Effect != nil {
			renderer, err := rendering.GetRenderer(*set.Effect)
			if err != nil {
				return false, err
			}

			h.manager.Render(h.ctx, renderer)
		}

		h.manager.SetBrightness(set.Brightness)
	}

	// Send state Update
	h.stateUpdates <- h.GetState()

	return true, nil
}
