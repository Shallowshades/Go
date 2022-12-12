package weighconv

// 1 英镑 = 0.4536 公斤
func PToK(p Pound) Kilogram {
	return Kilogram(p / Pound(KPRate))
}

func KToP(k Kilogram) Pound {
	return Pound(k * Kilogram(KPRate))
}
