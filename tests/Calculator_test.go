package tests

import "calculator/machine"
import "testing"

var mockedInput1 string = "1+2+3+4" // 10
var mockedInput2 string = "1+2/3+4" // 5
var mockedInput3 string = "2+3*1-2" // 3

func TestProcess(t *testing.T) {
	t.Log("Trying mocked inputs ...")

	if result := machine.Process(mockedInput1); result != 10 {
		t.Errorf("Expected 10, but got %v", result)
	}
	if result := machine.Process(mockedInput2); result != 5 {
		t.Errorf("Expected 5, but got %v", result)
	}

	if result := machine.Process(mockedInput3); result != 3 {
		t.Errorf("Expected 3, but got %v", result)
	}
}
