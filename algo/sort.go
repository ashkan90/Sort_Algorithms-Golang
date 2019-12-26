package algo

import (
	"math/rand"
)

type flex []int

func (d flex) Len() int {
	return len(d)
}

func (d flex) Less(i, j int) bool {
	return d[i] < d[j]
}

func (d *flex) Swap(i, j int) {
	cop := *d
	cop[i], cop[j] = cop[j], cop[i]
}

func BubbleSort(array flex) []int {
	arrayLength := array.Len()

	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength-i-1; j++ {

			if array.Less(j + 1, j) {
				array.Swap(j, j + 1)
			}
		}
	}
	return array
}



func Qsort(a flex) []int {
	if a.Len() < 2 { return a }

	left, right := 0, a.Len() - 1

	piv := rand.Int() % a.Len() // assume that this is '5'

	// pivot ile right ın değerleri değişecek.
	a.Swap(piv, right)
	//a[piv], a[right] = a[right], a[piv]


	// dump = [4, 1, 8, 17, 12, 76, 14, 0, 1, 3]
	for k := range a { // k starts from 0
		// a[0](4) < a[5](76)
		if a.Less(k, right) {
			// bulunan indexin değerini left ile değiştiriyor
			a.Swap(k, left)
			//a[k], a[left] = a[left], a[k]
			// a[0] = a[0]
			// a[0] = a[0]
			left++
		}
	}

	// son küçük elemanın yeri ayarlanıyor
	a.Swap(left, right)
	//a[left], a[right] = a[right], a[left]

	// assume left is 4
	Qsort(a[:left]) // unsorted v. dump[4, 1, 8, 17]
	Qsort(a[left + 1:]) // unsorted v. dump[12, 76, 14, 0, 1, 3]

	return a
}