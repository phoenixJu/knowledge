package filter

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

type ToIntFilter struct {
}

func NewToIntFilter() Filter {
	return &ToIntFilter{}
}

func (f *ToIntFilter) Process(request Request) (response Response, err error) {
	req, ok := request.([]string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid request for toIntFilter: %v", request))
	}
	res := []int{}
	for _, i := range req {
		r, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}
