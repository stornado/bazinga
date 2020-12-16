package compare

func Min(a, b int, others ...int) int {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinInt8(a, b int8, others ...int8) int8 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinInt16(a, b int16, others ...int16) int16 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinInt32(a, b int32, others ...int32) int32 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinInt64(a, b int64, others ...int64) int64 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinFloat32(a, b float32, others ...float32) float32 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinFloat64(a, b float64, others ...float64) float64 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinUint(a, b uint, others ...uint) uint {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinUint8(a, b uint8, others ...uint8) uint8 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinUint16(a, b uint16, others ...uint16) uint16 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinUint32(a, b uint32, others ...uint32) uint32 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinUint64(a, b uint64, others ...uint64) uint64 {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinByte(a, b byte, others ...byte) byte {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}

func MinRune(a, b rune, others ...rune) rune {
	min := a
	if b < a {
		min = b
	}

	for _, other := range others {
		if other < min {
			min = other
		}
	}

	return min
}
