package fibonacci

// Fibonacci calculate Fibonacci regular version
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

// TailRecursive calculate Fibonacci tail recursive version
func TailRecursive(n int) int {
	return fibonacciTailRecInner(n, 1, 0)
}

// fibonacciTailRecInner calculate Fibonacci tail recursive version using an accumulator
func fibonacciTailRecInner(n int, a int, b int) int {
	if n == 0 {
		return b
	}
	return fibonacciTailRecInner(n-1, a+b, a)
}
