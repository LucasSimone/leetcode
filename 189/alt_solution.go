import "fmt"

//Need to fix. The number of chains is wrong. What is the math to find the chains without adding time complexity
func rotate(nums []int, k int)  {

    k = len(nums) % k

    if k == 0 {
        return 
    }

    var chains int = 0
    
    if len(nums) % 2 == 0{
        if len(nums) % k != 0 {
            chains = len(nums) % k 
        }else {
            chains = k
        }
    }else{
        chains = 1
    }

    fmt.Printf("len: %d\n",len(nums))
    fmt.Printf("k: %d\n",k)


    for i:=0; i<chains; i++ {
        var first_chain_hold int = nums[i]
        var write_index = i
        for j:=i; j<len(nums)/chains; j++ {
            fmt.Printf("chain_index: %d\n",j)
            var read_index = getIndex(write_index,k,len(nums))
            nums[chain_index] = nums[read_index]
            write_index = read_index
        }
        write_index = first_chain_hold
    }

}

func getIndex(i,k,len) int {
    if i+k > len {
        return k % len
    }
    return i
}
// 3
// 1 2 3 4 5 6
// chains - 3
//
// 4
// 1 2 3 4 5 6
// chains 2

