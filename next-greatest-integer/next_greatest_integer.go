package next_greatest_integer

import (
	"strconv"
	"sort"
	"math"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func GetNextGreatestInt(num int) int {
	//Convert num into []int so that it's easier to manipulate
	length := len(strconv.Itoa(num))
	tmp := num
	i := length - 1
	digits := make([]int, length)
	for tmp > 0 {
		digits[i] = tmp % 10
		i--
		tmp = tmp / 10
	}

	//Start from the right-most digit and find the first digit that is smaller than the digit to its right
	//Example: Given the number "534976", we would stop at 4 because it is smaller than 9
	i = length - 1
	for ; i > 0; i-- {
		if digits[i] > digits[i-1] {
			break
		}
	}

	//If no such digit was found, then the digits are in descending order and no larger number can be found
	//Return 0 for false
	if i == 0 {
		return 0
	}

	//Find the smallest digit on right side of (i-1)'th digit that is greater than digits[i-1]
	//Example: Given the same number, "534976", this would be "6" because it is the smallest digit to the right of 4
	number := digits[i - 1]
	smallest := i
	for j := i + 1; j < length; j++ {
		if digits[j] > number && digits[j] < digits[smallest] {
			smallest = j
		}
	}

	//Swap the above found smallest digit with digits[i-1]
	swap(&digits[smallest], &digits[i - 1])

	//Sort the digits after (i-1) in ascending order
	digitsToSort := digits[i:]
	sort.Ints(digitsToSort)
	l := 0
	for k := i; k < length; k++ {
		digits[k] = digitsToSort[l]
		l++
	}

	//Format the result as a number
	numToReturn := 0
	pow := 0
	for k := length - 1; k >= 0; k-- {
		numToReturn = int(float64(numToReturn) + float64(digits[k]) * (math.Pow(float64(10), float64(pow))))
		pow++
	}

	return numToReturn
}