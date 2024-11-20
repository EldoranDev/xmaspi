//go:build !linux

package led

import "fmt"

var _ Controller = (*loggerController)(nil)

type loggerController struct {
	brightness uint8
	settings   *Settings
}

func NewController(settings *Settings) Controller {
	fmt.Printf("Starting with %d LEDs\n", len(settings.Leds))

	return &loggerController{
		brightness: settings.MaxBrightness,
		settings:   settings,
	}
}

func (*loggerController) Init() error {
	fmt.Println("Initializing Controller")

	return nil
}

func (*loggerController) Apply() error {
	fmt.Println("Applying current controller state")
	return nil
}

func (*loggerController) Close() {
	fmt.Println("Closing Controller")
}

func (l *loggerController) GetLedCount() uint32 {
	return uint32(len(l.settings.Leds))
}

func (l *loggerController) SetLed(led int, color *Color) {
	l.SetLedRaw(led, color.Int())
}

func (*loggerController) SetLedRaw(led int, color uint32) {
	fmt.Printf("Setting LED %d with color %d\n", led, color)
}

func (l *loggerController) Fill(color *Color) {
	l.FillRaw(color.Int())
}

func (l *loggerController) FillRaw(color uint32) {
	//TODO implement me
	fmt.Printf("Filling with %d\n", color)
}

func (l *loggerController) GetBrightness() uint8 {
	return l.brightness
}

func (lc *loggerController) SetBrightness(brightness uint8) {
	lc.brightness = brightness
}
