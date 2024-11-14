package mqtt

type Set struct {
	Brightness uint8
	Color      Color
	Effect     string
	State      string
}

type Color struct {
	R uint8
	G uint8
	B uint8
}
