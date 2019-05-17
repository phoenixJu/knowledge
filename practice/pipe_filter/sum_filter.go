package filter

import (
	"fmt"
	"github.com/pkg/errors"
)

type SumFilter struct {

}

func NewSumFilter() Filter  {
return &SumFilter{}
}

func (f *SumFilter)Process(request Request)(response Response, err error)  {
	req, ok := request.([]int)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid request: %v for sumFilter.", request))
	}
	sum :=0
	for _, ele := range req{
		sum += ele
	}
	return sum, nil
}
