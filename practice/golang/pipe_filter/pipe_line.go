package filter

type StraightPipeLine struct {
	Name    string
	Filters *[]Filter
}

func NewStraightPipeLine(name string, filters ...Filter) *StraightPipeLine {
	return &StraightPipeLine{Name: name, Filters: &filters}
}

func (p *StraightPipeLine) Process(data Request) (response Response, err error) {
	var ret interface{}
	for _, filter := range *p.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}

