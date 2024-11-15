//go:build !linux

package led

import "fmt"

var _ Controller = (*loggerController)(nil)

type loggerController struct {
	brightness uint8
}

func NewController() Controller {
	return &loggerController{}
}

func (*loggerController) Init() {
	fmt.Println("Initializing Controller")
}

func (*loggerController) Apply() {
	fmt.Println("Applying current controller state")
}

func (*loggerController) Close() {
	fmt.Println("Closing Controller")
}

func (lc *loggerController) GetBrightness() uint8 {
	return lc.brightness
}

func (lc *loggerController) SetBrightness(brightness uint8) {
	lc.brightness = brightness
}
