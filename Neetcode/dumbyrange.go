package main

import "errors"

type Data struct {
	prefix []int
}

func New(raw []int) (*Data, error) {
	if len(raw) == 0 {
		return nil, errors.New("Lens raw = 0")
	}

	prefix := make([]int, len(raw)+1)

	for i := range prefix {
		prefix[i+1] = prefix[i] + raw[i]
	}

	return &Data{
		prefix: prefix,
	}, nil
}

// 1, 2, 3, 4
// 0, 1, 3, 6, 10
func (d *Data) SumByRange(left, right int) (int, error) {
	if left > right || left < 0 || right > len(d.prefix)-2 {
		return -1, errors.ErrUnsupported
	}

	return d.prefix[right+1] - d.prefix[left], nil
}
