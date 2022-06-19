package sorting

func Insert_Sort(data []int) {
	for i, v := range data {
		now := v
		for j := i - 1; j >= 0; j-- {
			if data[j] > now {
				data[j+1] = data[j]
				data[j] = now
			} else {
				break
			}
		}
	}
}
