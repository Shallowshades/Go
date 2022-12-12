package lenconv

func MToF(m Metre) Foot {
	return Foot(m * Metre(MFRate))
}

func FToM(f Foot) Metre {
	return Metre(f / Foot(MFRate))
}
