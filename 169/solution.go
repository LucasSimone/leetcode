func majorityElement(nums []int) int {
    var count int = 0
    var majority = 0

    for i:=0;i<len(nums);i++ {
        
        if count == 0 {
            majority = nums[i]
        }

        if nums[i] == majority {
            count++
        }else{
            count--
        }
    }

    return majority
}
