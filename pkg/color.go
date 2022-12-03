package pkg

type Color struct {
	B uint8
	R uint8
	G uint8
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
