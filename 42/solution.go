func trap(height []int) int {

    var left_index int = 0
    var right_index int = len(height)-1
    var max_left int = height[0]
    var max_right int = height[len(height)-1]
    var trapped_water int = 0

    //Water[i] = min_height(left, right) - height[i]
    for i:=0; i<len(height)-2; i++{
        if(max_left < max_right){ 
            left_index++

            var water_level int = max_left - height[left_index]
            if water_level > 0 {
                trapped_water += water_level
            }


            if(height[left_index] > max_left){ max_left = height[left_index]}
        }else{
            right_index--

            var water_level int = max_right - height[right_index]
            if water_level > 0 {
                trapped_water += water_level
            }


            if(height[right_index] > max_right){ max_right = height[right_index]}
        } 
    }
    return trapped_water
}

