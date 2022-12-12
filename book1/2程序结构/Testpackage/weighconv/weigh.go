package weighconv

import "fmt"

const OneKilogram Kilogram = 1
const KPRate float64 = 0.4536

type Pound float64

type Kilogram float64

func (k Kilogram) String() string {
	return fmt.Sprintf("%gKg", k)
}

func (p Pound) String() string {
	return fmt.Sprintf("%gP", p)
}
