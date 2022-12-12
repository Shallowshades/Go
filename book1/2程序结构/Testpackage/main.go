package main

import (
	"conv/lenconv"
	"conv/tempconv"
	"conv/weighconv"
	"fmt"
)

func main() {
	{
		fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)

		fmt.Println(tempconv.CToF(tempconv.BoilingC))

		fmt.Println("tempconv.AbsoluteZeroK = ", tempconv.AbsoluteZeroK)

		fmt.Println(tempconv.CToK(tempconv.AbsoluteZeroC))
	}

	{
		fmt.Println("One Kg = ", weighconv.KToP(weighconv.OneKilogram))

		var k weighconv.Kilogram = 20.0
		var p weighconv.Pound = 45.0

		fmt.Println(k, " = ", weighconv.KToP(k))
		fmt.Println(p, " = ", weighconv.PToK(p))
	}

	{
		fmt.Println("One Metre = ", lenconv.MToF(lenconv.OneMetre))

		var m lenconv.Metre = 100
		var f lenconv.Foot = 100

		fmt.Println(m, " = ", lenconv.MToF(m))
		fmt.Println(f, " = ", lenconv.FToM(f))
	}
}
