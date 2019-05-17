package filter

import "testing"

func TestNewStraightPipeLine(t *testing.T) {
	splitFilter := NewSplitFilter(",")
	toIntFilter := NewToIntFilter()
	sumFilter := NewSumFilter()
	pipeLine := NewStraightPipeLine("processor", splitFilter, toIntFilter, sumFilter)
	res, err := pipeLine.Process("1,2,3,4")
	if err != nil{
		t.Errorf("error result : %s", err.Error())
		return
	}
	if res != 10 {
		t.Fatalf("the expect is 10, but the result is %d", res)
	}

}
