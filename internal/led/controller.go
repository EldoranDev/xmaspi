package led

type Controller interface {
	Init()
	Close()
	Apply()
}
