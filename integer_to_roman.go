package main

func intToRoman(num int) string {
	roman := ""

	i := 1
	for {
		if num/i < 10 {
			break
		}
		i *= 10
	}

	specials := map[int]string{
		4:   "IV",
		9:   "IX",
		40:  "XL",
		90:  "XC",
		400: "CD",
		900: "CM",
	}

	normals := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	normalsArr := []int{1000, 500, 100, 50, 10, 5, 1}

	for i > 0 {
		n := (num / i) * i
		num -= n
		if v, ok := specials[n]; ok {
			roman += v
			continue
		}

		j := 0
		for n > 0 {
			if n < normalsArr[j] {
				j++
				continue
			}
			n -= normalsArr[j]
			roman += normals[normalsArr[j]]
		}
		i /= 10
	}

	return roman
}
