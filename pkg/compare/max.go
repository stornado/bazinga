package compare

func Max(a, b int, others ...int) int {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxInt8(a, b int8, others ...int8) int8 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxInt16(a, b int16, others ...int16) int16 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxInt32(a, b int32, others ...int32) int32 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxInt64(a, b int64, others ...int64) int64 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxFloat32(a, b float32, others ...float32) float32 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxFloat64(a, b float64, others ...float64) float64 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxUint(a, b uint, others ...uint) uint {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxUint8(a, b uint8, others ...uint8) uint8 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxUint16(a, b uint16, others ...uint16) uint16 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxUint32(a, b uint32, others ...uint32) uint32 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}

func MaxUint64(a, b uint64, others ...uint64) uint64 {
	max := a
	if b > a {
		max = b
	}

	for _, other := range others {
		if other > max {
			max = other
		}
	}

	return max
}
