package pagination

import "strconv"

func convLimit(limit string) int64 {
	data, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return 0
	}
	return data
}

func convPage(page string) int64 {
	data, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 0
	}
	return data
}

func GetLimitAndPage(limit, page string) (l, p int64) {
	l = convLimit(limit)
	p = convPage(page)
	return l, p
}
