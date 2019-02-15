package pick_peaks

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func PickPeaks(array []int) PosPeaks {
	if len(array) < 3 {
		return PosPeaks{[]int{}, []int{}}
	}

	peaks := PosPeaks{[]int{}, []int{}}
	var position int
	var value int
	falling := array[0] >= array[1]

	for i := 1; i < len(array)-1; i++ {
		if array[i-1] < array[i] {
			position = i
			value = array[i]
			falling = false
		}

		if array[i] > array[i+1] && !falling {
			peaks.Peaks = append(peaks.Peaks, value)
			peaks.Pos = append(peaks.Pos, position)
			falling = true
		}
	}

	return peaks
}
