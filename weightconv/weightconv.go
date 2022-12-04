// Package weightconv perform pound and kilogram conversion
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const (
	Conversion = 0.45359237
)

func (p Pound) String() string {return fmt.Sprintf("%glb", p)} 
func (k Kilogram) String() string {return fmt.Sprintf("%gkg", k)}