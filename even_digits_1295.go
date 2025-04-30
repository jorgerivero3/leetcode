// nums - slice of positive integers
// return how many numbers have an even number of digits

func findNumbers(nums []int) int {
    var result int

    for _, num := range nums {
        // count digits
        digitCount := 0
        for n := num; n > 0; n /= 10 {
            digitCount++
        }

        // check if digit count is even
        if digitCount%2 == 0 {
            result++
        }
    }

    return result
}
