package main

import "fmt"

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
func main()  {
	splitFilter := NewSplitFilter(",")
	toIntFilter := NewToIntFilter()
	sumFilter := NewSumFilter()
	pipeLine := NewStraightPipeLine("processor", splitFilter, toIntFilter, sumFilter)
	res, err := pipeLine.Process("1,2,3,4")
	if err != nil{
		fmt.Println("error result : %s", err.Error())
		return
	}
	if res != 10 {
		fmt.Println("the expect is 10, but the result is %d", res)
	}
}
