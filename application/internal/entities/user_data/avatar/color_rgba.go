package avatar

type ColorRGBA struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewColorRGB(r uint8, g uint8, b uint8) *ColorRGBA {
	return &ColorRGBA{R: r, G: g, B: b}
}

func NewColorRGBA(r uint8, g uint8, b uint8, a uint8) *ColorRGBA {
	return &ColorRGBA{R: r, G: g, B: b, A: a}
}

func NewColorRGBAFromInt32(i int32) *ColorRGBA {
	a := uint8(i & 0x11111111)
	i >>= 8
	b := uint8(i & 0x11111111)
	i >>= 8
	g := uint8(i & 0x11111111)
	i >>= 8
	r := uint8(i & 0x11111111)
	return NewColorRGBA(r, g, b, a)
}

func (c *ColorRGBA) ToInt32() (result int32) {
	result += int32(c.R)
	result <<= 8
	result += int32(c.G)
	result <<= 8
	result += int32(c.B)
	result <<= 8
	result += int32(c.A)
	return
}
