package led

type Controller interface {
	Init()
	Close()
	Apply()

	SetBrightness(uint8)
	GetBrightness() uint8
}
