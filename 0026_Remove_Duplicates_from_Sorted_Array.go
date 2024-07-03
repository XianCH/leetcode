package leetcode

func removeDouplicate(array []int) int {
	finder, last := 0, 0

	for finder < len(array)-1 {
		for array[finder] == array[last] {
			finder++
			if finder == len(array) {
				return last + 1
			}
		}
		array[last+1] = array[finder]
		last++
	}
	return last + 1
}
