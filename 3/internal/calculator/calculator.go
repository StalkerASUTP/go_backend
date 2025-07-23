package calculator

type SimpleCalculator struct{}

func NewCalculator() *SimpleCalculator {
	return &SimpleCalculator{}
}

func (c *SimpleCalculator) Sum(numbers []int64) int64 {
	var sum int64
	for _, num := range numbers {
		sum += num
	}
	return sum
}
