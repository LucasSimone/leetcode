func jump(nums []int) int {
    
    var jump_index int = len(nums)-1
    var num_jumps int = 0

    for jump_index > 0 {

        var left_index int = jump_index
      
        //Find the max left jump
        for i:=jump_index; i>=0; i-- {
            if nums[i] >= jump_index - i{
                left_index = i
            }
        }

        //If no new left value is found 
        if left_index == jump_index{
            return 0
        }

        //Update the new index to jump to
        jump_index = left_index
        num_jumps++
    }
    
    return num_jumps
}

