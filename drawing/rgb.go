package drawing

import "fmt"

// RGB color.
type RGB struct {
	Red   byte
	Green byte
	Blue  byte
}

func (c RGB) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.Red, c.Green, c.Blue)
}

func decimal(b int) string {
	switch b {
	case 0:
		return "0"
	case 1:
		return "1"
	case 2:
		return "2"
	case 3:
		return "3"
	case 4:
		return "4"
	case 5:
		return "5"
	case 6:
		return "6"
	case 7:
		return "7"
	case 8:
		return "8"
	case 9:
		return "9"
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	default:
		return "F"
	}
}

// Hex code.
func (c *RGB) Hex() (hex string) {
	hex = "#"
	var code float64
	var nofloat int
	// Red
	code = float64(c.Red) / 16
	nofloat = int(code)
	hex += decimal(nofloat)
	hex += decimal(int((code - float64(nofloat)) * 16))
	// Green
	code = float64(c.Green) / 16
	nofloat = int(code)
	hex += decimal(nofloat)
	hex += decimal(int((code - float64(nofloat)) * 16))
	// Blue
	code = float64(c.Blue) / 16
	nofloat = int(code)
	hex += decimal(nofloat)
	hex += decimal(int((code - float64(nofloat)) * 16))
	return
}

// Reverse returns reverse of color.
func (c *RGB) Reverse() RGB {
	return RGB{
		Red:   255 - c.Red,
		Green: 255 - c.Green,
		Blue:  255 - c.Blue,
	}
}
