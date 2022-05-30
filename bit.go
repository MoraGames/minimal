package minimal

import "fmt"

type bit uint8

func SetBit (value int)(*bit, error) {
	var b bit

	if value == 0 || value == 1 {
		b = bit(value)
		return &b, nil
	}
	return nil, fmt.Errorf("'bit' type cannot assume values other than 0 and 1")
}