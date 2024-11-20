package led

type Color struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

func (c *Color) Int() uint32 {
	out := uint32(c.G)

	out = (out << 8) + uint32(c.R)
	out = (out << 8) + uint32(c.B)

	return out
}

func (c *Color) DimmedInt(factor float32) uint32 {
	out := uint32(float32(c.G) * factor)

	out = (out << 8) + uint32(float32(c.R)*factor)
	out = (out << 8) + uint32(float32(c.B)*factor)

	return out
}
