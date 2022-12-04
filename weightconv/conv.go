package weightconv

// Converts pound to kilogram
func PToK (p Pound) Kilogram {return Kilogram(p * Conversion)}
func KToP (k Kilogram) Pound {return Pound(k / Conversion)}

