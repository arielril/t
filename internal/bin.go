package internal

import (
	"strconv"
)

type Bin string

func (b Bin) FromString() []byte {
	s := string(b)
	if len(s)%8 != 0 {
		return []byte{}
	}

	res := make([]byte, len(s)/8)
	for i := 0; i < len(s); i += 8 {
		part := s[i : i+8]
		c, err := strconv.ParseUint(part, 2, 8)
		if err != nil {
			return []byte{}
		}
		res = append(res, byte(c))
	}

	return res
}
