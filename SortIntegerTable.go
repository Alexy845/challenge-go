package piscine

func SortIntegerTable(table []int) []int {
	for i := 0; i <= len(table)-1; i++ {
		for j := i + 1; j <= len(table)-1; j++ {
			if table[i] > table[j] {
				sauv := table[j]
				table[j] = table[i]
				table[i] = sauv
			}
		}
	}
	return table
}
