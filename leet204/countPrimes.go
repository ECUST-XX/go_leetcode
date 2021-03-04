package leet204

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	res := 1
	for i := 2; i < n; i++ {
		if isPrime(i) {
			res++
		}
	}
	return res
}

func isPrime(n int) bool {
	if n%2==0 {
		return false
	}

	for i := 3; i*i <= n; i=i+2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
