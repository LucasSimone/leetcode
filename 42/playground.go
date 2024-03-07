package main

import "fmt"

func main() {
    height := []int{0,1,0,2,1,0,1,3,1,1,2,1}

    var ret int = trap(height)

    fmt.Printf("Trapped Water: %d",ret)
}

func trap(height []int) int {

    var left_index int = 0
    var right_index int = len(height)-1
    var max_left int = height[0]
    var max_right int = height[len(height)-1]
    var trapped_water int = 0

    //Water[i] = min_height(left, right) - height[i]
    for i:=0; i<len(height)-2; i++{
        fmt.Printf("%d\n",i)
        fmt.Printf("Max Left: %d\n",max_left)
        fmt.Printf("Max Right: %d\n",max_right)
        if(max_left < max_right){ 
            fmt.Printf("Left Index: %d\n",left_index+1)
            left_index++

            var water_level int = max_left - height[left_index]
            if water_level > 0 {
                trapped_water += water_level
            }

            fmt.Printf("Water Level: %d\n",water_level)

            if(height[left_index] > max_left){ max_left = height[left_index]}
        }else{
            fmt.Printf("Right Index: %d\n",right_index-1)
            right_index--

            fmt.Printf("Right Index Height: %d\n",height[right_index])
            var water_level int = max_right - height[right_index]
            if water_level > 0 {
                trapped_water += water_level
            }

            fmt.Printf("Water Level: %d\n",water_level)

            if(height[right_index] > max_right){ max_right = height[right_index]}
        } 
        fmt.Printf("Trapped water: %d\n",trapped_water)
    }
    return trapped_water
}
