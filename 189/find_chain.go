package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    len, err := strconv.Atoi(os.Args[1]) 
    step, err2 := strconv.Atoi(os.Args[2])
    if err == nil && err2 == nil {
        find_chain_length(len, step)
    }
}

func find_chain_length(length int, step int) {
    
    var index int = 0
    var chain_length = 0

    for {

        index = getIndex(index,step,length)
        chain_length++
        
        //fmt.Printf("Index: %d\n",index)

        if index == 0 {
            break
        }
    }


    fmt.Printf("Length: %d\n",length)
    fmt.Printf("Steps: %d\n",step)
    fmt.Printf("Chains: %d\n",length/chain_length)
    fmt.Printf("Chain Length: %d\n",chain_length)
    
}

func getIndex(i int,k int,len int) int {
    if i-k < 0 {
        return len + (i-k)
    }
    return i-k
}
