package unitconv_test

import (
	"testing"

	"github.com/bunorita/gopl/ch02/ex02/unitconv"
)

func TestMToF(t *testing.T) {
	actual := unitconv.FToM(1)
	expected := unitconv.MPerFt
	if actual != expected {
		t.Errorf("%v is expected. but actual value is %v\n", expected, actual)
	}

}

func TestFToM(t *testing.T) {
	actual := unitconv.MToF(1)
	expected := unitconv.Feet(1 / unitconv.MPerFt)
	if actual != expected {
		t.Errorf("%v is expected. but actual value is %v\n", expected, actual)
	}
}

func TestKgToPound(t *testing.T) {
	actual := unitconv.KgToPound(1)
	expected := unitconv.Pound(1 / unitconv.KgPerPound)
	if actual != expected {
		t.Errorf("%v is expected. but actual value is %v\n", expected, actual)
	}
}

func TestPoundToKg(t *testing.T) {
	actual := unitconv.PoundToKg(1)
	expected := unitconv.KgPerPound
	if actual != expected {
		t.Errorf("%v is expected. but actual value is %v\n", expected, actual)
	}
}
