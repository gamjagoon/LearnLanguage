package sorting

func Select_Sort(data []int) {
	l := len(data) - 1
	for i, v := range data {
		now := v
		ptr := i
		for j := l; j > i; j-- {
			if now > data[j] {
				now = data[j]
				ptr = j
			}
		}
		now = data[ptr]
		data[ptr] = data[i]
		data[i] = now
	}
}
