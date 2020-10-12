// Package unitconv performs meter/feet and pond/kg conversions.
package unitconv

import "fmt"

// length unit
type Meter float64
type Feet float64

// weight unit
type Kg float64
type Pound float64

const (
	// MPerFt means 1[ft] == 0.3048[m]
	MPerFt Meter = 0.3048
	// KgPerPound means 1[ld] == 0.45359237[kg]
	KgPerPound Kg = 0.45359237
)

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}

func (kg Kg) String() string {
	return fmt.Sprintf("%gkg", kg)
}

func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}

func MToF(m Meter) Feet {
	return Feet(m / MPerFt)
}

func FToM(f Feet) Meter {
	return Meter(f) * MPerFt
}

func KgToPound(kg Kg) Pound {
	return Pound(kg / KgPerPound)
}

func PoundToKg(p Pound) Kg {
	return Kg(p) * KgPerPound
}
