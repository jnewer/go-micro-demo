package service

func GetSum(input ...int64) (ret int64) {
	for _, v := range input {
		ret += v
	}
	return ret
}
