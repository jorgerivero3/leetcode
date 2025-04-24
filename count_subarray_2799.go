// nums - positive ints
//subarray is complete if the distinct elements in subarray is equal to the num of distinct elements in whole array
//return the number of complete arrays


func countCompleteSubarrays(nums []int) int {
    //find the set, amount of distinct numbers in nums
    distinct := map[int]bool{}

    for _, num := range nums {
        distinct[num] = true
    }

    target := len(distinct)
    //count subarrays
    var result int
    for i, _ := range nums{
        freq := make(map[int]int)
        distinctCount := 0

        for j := i; j < len(nums); j++ {
            if freq[nums[j]] == 0{
                distinctCount++ 
            }
            freq[nums[j]]++

            if distinctCount == target{
                result++
            }
        }
    }


    return result
}
