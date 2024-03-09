func rotate(nums []int, k int)  {

    //reduce rotations greater than the array length
    k = k % len(nums)
    //Reverse the whole array
    reverse(nums, 0 ,len(nums)-1)
    //Reverse from 0 - k
    reverse(nums, 0, k-1)
    //Reverse from k - end
    reverse(nums, k, len(nums)-1)
}

func reverse(arr []int, start int, end int) []int {
    for start < end{

        arr[start] += arr[end]
        arr[end] = arr[start] - arr[end]
        arr[start] -= arr[end]

        start++
        end--
    }

    return arr
}


