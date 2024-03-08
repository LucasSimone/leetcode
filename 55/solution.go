func canJump(nums []int) bool {
    
    var jump_index int = len(nums)-1

    //Loop from back to front
    for i:=jump_index; i>=0; i-- {
        //Update current position if a further left index can jump to the current position 
        if nums[i] >= jump_index - i{
            jump_index = i
        }
    }

    //Check if jumped to the begining of array
    if jump_index == 0 {
        return true
    }

    return false
}
