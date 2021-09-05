package drawing

import "errors"

// HexToRGB parses hex to RGB.
func HexToRGB(hex string) (c RGB, err error) {
	if hex[0] != '#' {
		return c, errors.New("invalid format")
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errors.New("invalid format")
		return 0
	}
	switch len(hex) {
	case 7:
		c.Red = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		c.Green = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		c.Blue = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		c.Red = hexToByte(hex[1]) * 17
		c.Green = hexToByte(hex[2]) * 17
		c.Blue = hexToByte(hex[3]) * 17
	default:
		err = errors.New("invalid format")
	}
	return
}
