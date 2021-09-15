package slices

// {
// 	name: "must order values 3...",
// 	args: args{
// 		values: []int{10, 5, 12, 20, 25},
// 		order:  []int{4, 3, 2, 1, 0},
// 	},
// 	want: []int{25, 20, 12, 5, 10},
// }
func orderSlice(values []int, order []int) []int {
	valuesOrdered := []int{}

	for i, v := range order {
		if len(valuesOrdered) > v && valuesOrdered[v] != 0 {
			valuesTemp := append([]int{}, valuesOrdered[0:v]...)

			valuesTemp = append(valuesTemp, values[i])

			valuesTemp = append(valuesTemp, valuesOrdered[v:]...)

			valuesOrdered = valuesOrdered[:0]
			valuesOrdered = append(valuesOrdered, valuesTemp...)
		} else {
			valuesOrdered = append(valuesOrdered, values[v])
		}
	}

	return valuesOrdered
}
