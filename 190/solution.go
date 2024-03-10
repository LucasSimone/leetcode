
func reverseBits(num uint32) uint32 {
    var res uint32
    for i:=0; i < 32; i++ {
        //shift result int 1 to the left 
        //if you shift after you shift the last bit out of place
        res = res << 1
        // Copy bit to result int
        var bit uint32 = (num >> i) & 1
        res = res | bit

    }
    return res
}
