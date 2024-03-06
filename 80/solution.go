package main


func removeDuplicates(nums []int) int {
    var currNumberIndex int = 0
    var currNumberCount int = 1

    for i := 1; i < len(nums); i++ {

        //If a new number is found
        if nums[i] > nums[currNumberIndex]{
            currNumberIndex++
            nums[currNumberIndex] = nums[i] 
            currNumberCount = 1
        //If I have the same number check if there are less than 2
        }else if currNumberCount < 2{
            currNumberCount++
            currNumberIndex++
            nums[currNumberIndex] = nums[i] 
        }

    }

    return currNumberIndex + 1
}

