import "fmt"
func majorityElement(nums []int) int {

    var left_index = 0
    var right_index = len(nums)-1

    var majority_left int = nums[0]
    var majority_left_count = 0
    var majority_right int = nums[len(nums)-1]
    var majority_right_count = 0

    var prospect_left int = 0
    var prospect_left_count = 0
    var prospect_right int = 0
    var prospect_right_count = 0

    for left_index <= right_index{

        fmt.Printf("Majority Left: %d\n",majority_left)
        fmt.Printf("Majority Left Count: %d\n",majority_left_count)
        fmt.Printf("Majority Right: %d\n",majority_right)
        fmt.Printf("Majority Right Count: %d\n",majority_right_count)
        fmt.Printf("Prospect Left: %d\n",prospect_left)
        fmt.Printf("Prospect Left Count: %d\n",prospect_left_count)
        fmt.Printf("Prospect Right: %d\n",prospect_right)
        fmt.Printf("Prospect Right Count: %d\n",prospect_right_count)


        // Left Check
        if nums[left_index] == majority_left {
            majority_left_count++
        }
        if nums[left_index] == majority_right {
            majority_right_count++
        }

        if nums[left_index] == prospect_left {
            prospect_left_count++
        }
        if nums[left_index] == prospect_right{
            prospect_right_count++
        }
        
        if nums[left_index] != prospect_left{
            prospect_left = nums[left_index]
            prospect_left_count = 1
        }

        // Right Check
        if nums[right_index] == majority_right {
            majority_right_count++
        }
        if nums[right_index] == majority_left {
            majority_left_count++
        }

        if nums[right_index] == prospect_right {
            prospect_right_count++
        }
        if nums[right_index] == prospect_left {
            prospect_left_count++
        }
        
        if nums[right_index] != prospect_right{
            prospect_right = nums[right_index]
            prospect_right_count = 1
        }

        // Check for a new majority
        if prospect_left_count >= majority_left_count{
             majority_left = prospect_left
             majority_left_count = prospect_left_count
        }

        if prospect_right_count >= majority_right_count{
             majority_right = prospect_right
             majority_right_count = prospect_right_count
        }
        

        left_index++
        right_index--
    }

    if majority_left_count > majority_right_count{
        return majority_left
    }
    return majority_right

 }


//  func majorityElement(nums []int) int {

//     var left_index = 0
//     var right_index = len(nums)-1

//     var majority_left int = nums[0]
//     var majority_left_count = 0
//     var majority_right int = nums[len(nums)-1]
//     var majority_right_count = 0

//     var prospect_left int = 0
//     var prospect_left_count = 0
//     var prospect_right int = 0
//     var prospect_right_count = 0

//     for left_index <= right_index{

//         // Left Check
//         if nums[left_index] == majority_left {
//             majority_left_count++
//         }
//         if nums[left_index] == majority_right {
//             majority_right_count++
//         }


//         if nums[left_index] == prospect_left {
//             prospect_left_count++
//         }else
//         if nums[left_index] == prospect_right{
//             prospect_right++
//         }else{
//             prospect_left = nums[left_index]
//             prospect_left_count = 1
//         }

//         // Right Check
//         if nums[right_index] == majority_right {
//             majority_right++
//         }

//         if nums[right_index] == prospect_right {
//             prospect_left_count++
//         }else
//         if nums[right_index] == prospect_left{
//             prospect_right++
//         }else{
//             prospect_right = nums[right_index]
//             prospect_right_count = 1
//         }

//         // Check for a new majority
//         if prospect_left_count >= majority_left_count{
//              majority_left = prospect_left
//              majority_left_count = prospect_left_count
//         }

//         if prospect_right_count >= majority_right_count{
//              majority_right = prospect_right
//              majority_right_count = prospect_right_count
//         }
        

//         left_index++
//         right_index--
//     }

//     if majority_left > majority_right{
//         return majority_left
//     }
//     return majority_right

//  }

// func majorityElement(nums []int) int {

//     var majority int = nums[0]
//     var majority_count = 1

//     var prospect int = nums[1]
//     var prospect_count = 1

//     for i:=2; i<len(nums)-1; i++ {
//         if nums[i] == majority{
//             majority_count++
//         } else
//         if nums[i] == prospect {
//             prospect_count++
//         }else{
//             prospect = nums[i]
//             prospect_count = 1
//         }

//         if prospect_count >= majority_count{
//             majority = prospect
//             majority_count = prospect_count
//         }
//     }
//     return majority
//  }


 // 1 1 4 1 1 4 1 1 4 4

 // 1 2 2 1 1 1 1 1 2 3
