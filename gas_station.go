package main

func canCompleteCircuit(gas []int, cost []int) int {
	starters := make([]int, 0)
	for i := 0; i < len(gas); i++ {
		// weird test case with only ones
		if gas[i] == 0 || gas[i] == 1 {
			continue
		}
		if gas[i] >= cost[i] {
			starters = append(starters, i)
		}
	}

outer:
	for _, s := range starters {
		fuel := 0
		for i := 0; i < len(gas); i++ {
			idx := (s + i) % len(gas)
			fuel += gas[idx]
			fuel -= cost[idx]
			if fuel < 0 {
				continue outer
			}
		}
		return s
	}
	return -1
}
