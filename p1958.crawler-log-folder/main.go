package main

func minOperations(logs []string) int {
	stack := []string{}

	for i := 0; i < len(logs); i++ {
		op := logs[i]
		switch op {
		case "./":
		case "../":
			if len(stack) == 0 {
				continue
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, op)
		}
	}
	return len(stack)
}
