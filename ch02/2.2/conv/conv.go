package conv

import "fmt"

type Meter float64
type Feet float64
type Celsius float64
type Fahrenheit float64
type Pound float64
type Kilogram float64

func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (f Feet) String() string       { return fmt.Sprintf("%gft", f) }
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%gkg", k) }

// MToF converts a meter distance to feets.
func MToF(m Meter) Feet { return Feet(m * 3.28084) }

// FToM converts a feet distance to meters.
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// PToK converts a Puonds weigth to Kilograms.
func PToK(p Pound) Kilogram { return Kilogram(p * 0.453592) }

// KToP converts a Kilogram weigth to Pounds.
func KToP(k Kilogram) Pound { return Pound(k * 2.20462) }
