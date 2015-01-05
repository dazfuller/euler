package eulerutil

func ToInteger(slice []int) int64 {
	k := int64(0)
	for i := 0; i < len(slice); i++ {
		k = int64(10)*k + int64(slice[i])
	}
	return k
}

func NextPermutation(slice []int) bool {
	var k int
	for k = len(slice) - 2; k >= 0; k-- {
		if slice[k] < slice[k+1] {
			break
		}
	}

	if k == -1 {
		return false
	}

	var l int
	for l = len(slice) - 1; l >= 0; l-- {
		if slice[k] < slice[l] {
			break
		}
	}

	slice[k], slice[l] = slice[l], slice[k]

	for i, j := k+1, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return true
}
