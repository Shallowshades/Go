package lenconv

import "fmt"

type Foot float64
type Metre float64

const OneMetre Metre = 1.0
const MFRate float64 = 3.28083989501

func (f Foot) String() string {
	return fmt.Sprintf("%gF", f)
}

func (m Metre) String() string {
	return fmt.Sprintf("%gm", m)
}
