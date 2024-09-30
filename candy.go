package main

func candy(ratings []int) int {
	candys := make([]int, len(ratings))
	r := len(ratings)

	for i := 0; i < r-1; i++ {
		if ratings[i] < ratings[i+1] {
			candys[i+1] = candys[i] + 1
		}
	}

	for i := r - 1; i > 0; i-- {
		if ratings[i] < ratings[i-1] {
			candys[i-1] = max(candys[i-1], candys[i]+1)
		}
	}

	for _, c := range candys {
		r += c
	}
	return r
}

func candyV1(ratings []int) int {
	candys := make([]int, len(ratings))

	for i := 0; i < len(ratings)-1; i++ {
		left, right := ratings[i], ratings[i+1]
		diff := left - right
		if diff < 0 {
			candys[i+1] = candys[i] + 1
		} else if diff > 0 {
			if candys[i] == 0 {
				candys[i] = 1
			} else {
				candys[i] = max(candys[i+1]+1, candys[i])
			}
			update(ratings, candys, i-1)
		} else {
			candys[i+1] = 0
		}
	}

	sol := len(ratings)
	for _, c := range candys {
		sol += c
	}

	return sol
}

func update(ratings, candys []int, idx int) {
	if idx < 0 {
		return
	}
	diff := ratings[idx] - ratings[idx+1]
	if diff <= 0 {
		return
	}
	prev := candys[idx]
	candys[idx] = max(candys[idx+1]+1, candys[idx])
	if prev != candys[idx] {
		update(ratings, candys, idx-1)
	}
}
