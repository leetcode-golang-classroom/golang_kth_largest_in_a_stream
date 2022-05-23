package sol

import "fmt"

func Run(commands *[]string, values *[][]int) []string {
	result := []string{"null"}
	cLen := len(*commands)
	k := (*values)[0][0]
	nums := (*values)[0][1:]
	kLargest := Constructor(k, nums)
	for idx := 1; idx < cLen; idx++ {
		value := (*values)[idx][0]
		result = append(result, fmt.Sprintf("%d", kLargest.Add(value)))
	}
	return result
}
