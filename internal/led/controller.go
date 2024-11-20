package led

type Controller interface {
	Init() error
	Close()
	Apply() error

	GetLedCount() uint32

	SetBrightness(uint8)
	GetBrightness() uint8

	SetLed(led int, color *Color)
	SetLedRaw(led int, color uint32)

	Fill(color *Color)
	FillRaw(color uint32)
}
