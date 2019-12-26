package misc
import ( "math/rand" )

type sort rune

const (
	PRE_SORTED sort = iota
	BACK_SORTED
	MIXED
)

func GenerateSlice(genType sort, size int) []int {
	a := make([]int, size)

	switch genType {
	case MIXED:
		for i := range a {
			a[i] = rand.Int() % size
		}
	case BACK_SORTED:
		i := size - 1
		for k := range a {
			a[k] = i
			i--
		}
	case PRE_SORTED:
		i := 0
		for _ = range a {
			a[i] = i
			i++
		}
	}

	return a
}