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

func SelectionSort(a flex) []int {
	var size = a.Len()

	var min int
	for i := 0; i < size; i++ {
		min = i

		// en küçük değeri arıyor
		for j := i; j < size; j++ {
			if a.Less(j, min) {
				min = j
			}
		}

		// en küçük değerle i nin yerini değiştiriyor
		a.Swap(i, min)
		//a[i], a[min] = a[min], a[i]
	}

	return a

}

func BubbleSort(a flex) []int {
	arrayLength := a.Len()

	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength-i-1; j++ {

			if a.Less(j+1, j) {
				a.Swap(j, j+1)
			}
		}
	}
	return a
}

func Qsort(a flex) []int {
	if a.Len() < 2 {
		return a
	}

	left, right := 0, a.Len()-1

	piv := rand.Int() % a.Len()

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
			left++
		}
	}

	// son küçük elemanın yeri ayarlanıyor
	a.Swap(left, right)
	//a[left], a[right] = a[right], a[left]

	Qsort(a[:left])   // dump[4, 1, 8, 17]
	Qsort(a[left+1:]) // dump[12, 76, 14, 0, 1, 3]

	return a
}

func MSort(items []int) []int {
	var n = len(items)

	if n == 1 {
		return items
	}

	middle := int(n / 2)
	// iki parça dizi hazırlanıyor
	var (
		l = make([]int, middle)
		r = make([]int, n-middle)
	)

	for i := 0; i < n; i++ {
		if i < middle { // soldaki diziye, index middle a ulaşana kadar değer atanıyor.
			l[i] = items[i]
		} else { // middle dan sonrası için de sağdaki dizi dolduruluyor
			r[i-middle] = items[i]
		}
	}

	return _merge(l, r)
}

func _merge(left, right flex) flex {
	result := make(flex, left.Len()+right.Len())

	i := 0
	for 0 < left.Len() && right.Len() > 0 {
		if left[0] < right[0] {
			result[i] = left[0] // karşılaştırmanın sonucunda left in 0'ı en küçük kabul edilip diziye atanıyor
			left = left[1:]     // dizi baştaki alınmayacak şekilde tekrar şekillendiriliyor.
		} else {
			result[i] = right[0] // aynı işlemler sağdaki dizi için de yapılıyor.
			right = right[1:]
		}
		i++
	}

	// bölünemeyecek hale gelen left deki değeri-leri tekrardan ana diziye yollamaya devam ediyor
	for j := 0; j < left.Len(); j++ {
		result[i] = left[j]
		i++
	}

	// bölünemeyecek hale gelen right daki değeri-leri tekrardan ana diziye yollamaya devam ediyor
	for j := 0; j < right.Len(); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
