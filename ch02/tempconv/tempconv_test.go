package tempconv_test

import (
	"testing"

	"github.com/bunorita/gopl/ch02/tempconv"
)

func TestCtoF(t *testing.T) {
	if tempconv.CToF(tempconv.BoilingC) != 212.0 {
		t.Errorf("212.0 is expected. but %v\n", tempconv.CToF(tempconv.BoilingC))
	}

}

func TestKToC(t *testing.T) {
	actual := tempconv.KToC(0)
	expected := tempconv.AbsoluteZeroC
	if actual != expected {
		t.Errorf("%v is expected. but actual value is %v\n", expected, actual)
	}

}

func TestCtoK(t *testing.T) {
	actual := tempconv.CToK(0)
	expected := -tempconv.Kelvin(tempconv.AbsoluteZeroC)
	if actual != expected {
		t.Errorf("%v is expected. but actual value is %v\n", expected, actual)
	}
}
