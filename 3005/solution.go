func maxFrequencyElements(nums []int) int {

    frequency := make(map[int]int)
    //Map the frequency
    for i:= range nums {
        frequency[nums[i]]++
    }

    //Loop through frequency count how many ties the max frequency appears
    var max_freq int = 0
    var count int = 0
    for k:= range frequency{
        if frequency[k] > max_freq {
            max_freq = frequency[k]
            count = 0
        }

        if frequency[k] == max_freq {
            count += frequency[k]
        }
    }   

    return count
}
