package eulermath

import (
	"reflect"
	"testing"
)

func TestGetTriangularNumbers(t *testing.T) {
	expected := map[int]int{
		1: 1,
		2: 3,
		3: 6,
		4: 10,
		5: 15,
		6: 21,
	}

	actual := make(map[int]int)

	for x := 1; x <= 6; x++ {
		actual[x] = GetTriangular(x)
	}

	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestIsTriangular(t *testing.T) {
	expected := map[int]bool{
		3:  true,
		6:  true,
		7:  false,
		10: true,
		15: true,
		20: false,
		21: true,
	}

	for k, v := range expected {
		chk := IsTriangular(k)
		if v != chk {
			t.Errorf("Expected %t for %d but got %t", v, k, chk)
		}
	}
}

func TestGetPentagonalNumbers(t *testing.T) {
	expected := map[int]int{
		1: 1,
		2: 5,
		3: 12,
		4: 22,
		5: 35,
		6: 51,
	}

	actual := make(map[int]int)

	for x := 1; x <= 6; x++ {
		actual[x] = GetPentagonal(x)
	}

	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestIsPentagonal(t *testing.T) {
	expected := map[int]bool{
		1:  true,
		5:  true,
		12: true,
		22: true,
		24: false,
		31: false,
		35: true,
		51: true,
	}

	for k, v := range expected {
		chk := IsPentagonal(k)
		if v != chk {
			t.Errorf("Expected %t for %d but got %t", v, k, chk)
		}
	}
}
