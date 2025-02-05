package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	costs := [3]int{}
	total := 0
	for i, m := range messages {
		costs[i] = len(m) + total
		total += len(m)
	}
	return messages, costs
}
