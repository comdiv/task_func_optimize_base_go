package implementations

import (
	"fmt"
	"github.com/comdiv/task_func_optimize_base_go/basis"
	"math/rand"
	"testing"
)

func TestAllFunctionality(t *testing.T) {
	for i, fd := range AllImplementations {
		testName := fmt.Sprintf("%d. %s", i, fd.Name)
		t.Run(testName, func(t *testing.T) {
			basis.SuperFuncTestCase(fd.Impl, t)
		})
	}
}

func BenchmarkDefault(b *testing.B) {
	for i, fd := range AllImplementations {
		testName := fmt.Sprintf("%d. %s", i, fd.Name)
		b.Run(testName, func(b *testing.B) {
			basis.SuperFuncBenchmark(fd.Impl, b)
		})
	}
}

func BenchmarkRandom_Low_2_20(b *testing.B) {
	for i, fd := range AllImplementations {
		testName := fmt.Sprintf("%d. %s", i, fd.Name)
		b.Run(testName, func(b *testing.B) {
			basis.SuperFuncTrueRandomBenchmark(fd.Impl, b, 2, 20)
		})
	}
}
func BenchmarkRandom_High_21_40(b *testing.B) {
	for i, fd := range AllImplementations {
		testName := fmt.Sprintf("%d. %s", i, fd.Name)
		b.Run(testName, func(b *testing.B) {
			basis.SuperFuncTrueRandomBenchmark(fd.Impl, b, 21, 40)
		})
	}
}

func BenchmarkRandom_Large_70_77(b *testing.B) {
	for i, fd := range AllImplementations {
		testName := fmt.Sprintf("%d. %s", i, fd.Name)
		b.Run(testName, func(b *testing.B) {
			basis.SuperFuncTrueRandomBenchmark(fd.Impl, b, 70, 77)
		})
	}
}

func BenchmarkFixed_77(b *testing.B) {
	for i, fd := range AllImplementations {
		testName := fmt.Sprintf("%d. %s", i, fd.Name)
		b.Run(testName, func(b *testing.B) {
			SuperFuncBenchmarkN(fd.Impl, 77, b)
		})
	}
}

type funcCase struct {
	x1 float64
	x2 float64
	n  uint8
}

const CASES_SIZE = 8192

func buildCases() []*funcCase {
	// функция вызывется только один раз на запуск - все получат одинаковый набор кейсов
	var seed int64 = rand.Int63()
	rnd := rand.New(rand.NewSource(seed))
	var result []*funcCase
	for i := 0; i < CASES_SIZE; i++ {
		x1 := rnd.Float64()
		x2 := rnd.Float64()
		var n uint8 = 40
		result = append(result, &funcCase{x1, x2, n})
	}
	return result
}

var cases []*funcCase = buildCases()

func SuperFuncBenchmarkN(impl basis.SuperFuncType, n uint8, b *testing.B) {
	for i := 0; i < b.N; i++ {
		fCase := cases[i%CASES_SIZE]
		impl(fCase.x1, fCase.x2, n)
	}
}
