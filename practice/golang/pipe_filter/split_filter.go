package filter

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type SplitFilter struct {
	Delimiter string
}

func NewSplitFilter(delimiter string) Filter {
	return &SplitFilter{Delimiter: delimiter,}
}

func (s *SplitFilter) Process(request Request) (response Response, err error) {
	req, ok := request.(string)
	if !ok {
		return nil, errors.New(fmt.Sprintf("Invalid Request for Split Filter: %v", request))
	}
	return strings.Split(req, s.Delimiter), nil
}
