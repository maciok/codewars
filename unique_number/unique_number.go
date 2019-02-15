package unique_number

func FindUniq(arr []float32) float32 {

	var common float32
	if arr[0] == arr[1] {
		common = arr[0]
	} else if arr[0] == arr[2] {
		common = arr[0]
	} else {
		return arr[0]
	}

	for i := 1; i < len(arr); i++ {
		if arr[i] != common {
			return arr[i]
		}
	}

	panic("Cannot find solution")
}
