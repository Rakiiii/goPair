package pairlib

//IntPair is struct reppresenting a pair of integers
type IntPair struct {
	First  int
	Second int
}

//QuicksortIntPairFirst sort slice of IntPair from low to high value of @IntPair.First with quicksort algorithm
//returns sorted slice
func QuicksortIntPairFirst(a []IntPair) []IntPair {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) >> 1

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].First < a[right].First {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuicksortIntPairFirst(a[:left])
	QuicksortIntPairFirst(a[left+1:])

	return a
}

//QuicksortIntPairFirst sort slice of IntPair from low to high value of @IntPair.Second with quicksort algorithm
//returns sorted slice
func QuicksortIntPairSecond(a []IntPair) []IntPair {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := len(a) >> 1

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i].Second < a[right].Second {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	QuicksortIntPairSecond(a[:left])
	QuicksortIntPairSecond(a[left+1:])

	return a
}

//ReversIntPairSlice returns reversed slice of @a IntPair
func ReversIntPairSlice(a []IntPair) []IntPair {
	res := make([]IntPair, len(a))

	for i, num := range a {
		res[len(a)-i-1] = num
	}

	return res
}
