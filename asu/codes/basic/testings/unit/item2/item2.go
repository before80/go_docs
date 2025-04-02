package item2

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
)

func Add[T constraints.Integer](a, b T) (T, error) {
	// 溢出检查
	if (b > 0 && a > (math.MaxInt-b)) || (b < 0 && a < (math.MinInt-b)) {
		return 0, fmt.Errorf("integer overflow: %v + %v", a, b)
	}
	return a + b, nil
}
