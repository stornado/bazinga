package slicejoin

func JoinStringSlice(a, b []string) ([]string, []string, []string, []string, []string) {
	allA := make(map[string]struct{})
	allB := make(map[string]struct{})
	onlyInA := make(map[string]struct{})           // left join - inner join
	onlyInB := make(map[string]struct{})           // right join - inner join
	bothInAAndB := make(map[string]struct{})       // inner join
	onlyInAOrBNotBoth := make(map[string]struct{}) // a + b - inner join
	eitherInAOrB := make(map[string]struct{})      // full outer join

	for _, aStr := range a {
		if _, ok := eitherInAOrB[aStr]; !ok {
			eitherInAOrB[aStr] = struct{}{}
		}
		if _, ok := allA[aStr]; ok {
			continue
		}
		allA[aStr] = struct{}{}
	}
	for _, bStr := range b {
		if _, ok := eitherInAOrB[bStr]; !ok {
			eitherInAOrB[bStr] = struct{}{}
		}
		if _, ok := allB[bStr]; ok {
			continue
		}
		allB[bStr] = struct{}{}
	}

	for k := range allA {
		if _, ok := allB[k]; ok {
			bothInAAndB[k] = struct{}{}
			continue
		}
		onlyInA[k] = struct{}{}
		onlyInAOrBNotBoth[k] = struct{}{}
	}

	for k := range allB {
		if _, ok := allA[k]; ok {
			bothInAAndB[k] = struct{}{}
			continue
		}
		onlyInB[k] = struct{}{}
		onlyInAOrBNotBoth[k] = struct{}{}
	}

	var elmOnlyInA, elmOnlyInB, elmBothInAAndB, elmOnlyInAOrBNotBoth, elmEitherInAOrB []string

	for k := range onlyInA {
		elmOnlyInA = append(elmOnlyInA, k)
	}
	for k := range onlyInB {
		elmOnlyInB = append(elmOnlyInB, k)
	}
	for k := range bothInAAndB {
		elmBothInAAndB = append(elmBothInAAndB, k)
	}
	for k := range onlyInAOrBNotBoth {
		elmOnlyInAOrBNotBoth = append(elmOnlyInAOrBNotBoth, k)
	}
	for k := range eitherInAOrB {
		elmEitherInAOrB = append(elmEitherInAOrB, k)
	}
	return elmOnlyInA, elmOnlyInB, elmBothInAAndB, elmOnlyInAOrBNotBoth, elmEitherInAOrB
}


func JoinIntSlice(a, b []int) ([]int, []int, []int, []int, []int) {
	allA := make(map[int]struct{})
	allB := make(map[int]struct{})
	onlyInA := make(map[int]struct{})           // left join - inner join
	onlyInB := make(map[int]struct{})           // right join - inner join
	bothInAAndB := make(map[int]struct{})       // inner join
	onlyInAOrBNotBoth := make(map[int]struct{}) // a + b - inner join
	eitherInAOrB := make(map[int]struct{})      // full outer join

	for _, aStr := range a {
		if _, ok := eitherInAOrB[aStr]; !ok {
			eitherInAOrB[aStr] = struct{}{}
		}
		if _, ok := allA[aStr]; ok {
			continue
		}
		allA[aStr] = struct{}{}
	}
	for _, bStr := range b {
		if _, ok := eitherInAOrB[bStr]; !ok {
			eitherInAOrB[bStr] = struct{}{}
		}
		if _, ok := allB[bStr]; ok {
			continue
		}
		allB[bStr] = struct{}{}
	}

	for k := range allA {
		if _, ok := allB[k]; ok {
			bothInAAndB[k] = struct{}{}
			continue
		}
		onlyInA[k] = struct{}{}
		onlyInAOrBNotBoth[k] = struct{}{}
	}

	for k := range allB {
		if _, ok := allA[k]; ok {
			bothInAAndB[k] = struct{}{}
			continue
		}
		onlyInB[k] = struct{}{}
		onlyInAOrBNotBoth[k] = struct{}{}
	}

	var elmOnlyInA, elmOnlyInB, elmBothInAAndB, elmOnlyInAOrBNotBoth, elmEitherInAOrB []int

	for k := range onlyInA {
		elmOnlyInA = append(elmOnlyInA, k)
	}
	for k := range onlyInB {
		elmOnlyInB = append(elmOnlyInB, k)
	}
	for k := range bothInAAndB {
		elmBothInAAndB = append(elmBothInAAndB, k)
	}
	for k := range onlyInAOrBNotBoth {
		elmOnlyInAOrBNotBoth = append(elmOnlyInAOrBNotBoth, k)
	}
	for k := range eitherInAOrB {
		elmEitherInAOrB = append(elmEitherInAOrB, k)
	}
	return elmOnlyInA, elmOnlyInB, elmBothInAAndB, elmOnlyInAOrBNotBoth, elmEitherInAOrB
}
