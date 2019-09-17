package ProjectEuler

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
//Find the sum of all the multiples of 3 or 5 below 1000.

func Multiplesof3and5()int{
	sum := 0
	for temp := 3; temp < 1000; temp*=3 {
		sum += temp
	}
	for temp := 5; temp < 1000; temp *= 5 {
		if temp%3 == 0 {
			continue
		}
		sum += temp
	}
	return sum
}
