func addBinary(a string, b string) string {
    var longer []byte
    var shorter []byte
    var sum string

    //Can store in a and b instead but this is easier to understand
    if len(a) > len(b) {
        longer = []byte(a)
        shorter = []byte(b)
    }else{
        longer = []byte(b)
        shorter = []byte(a)
    }

    var carry_byte byte = '0'
    for i:=len(longer)-1; i >= 0; i-- {
        var sum_byte byte
        var long_byte byte = longer[i]
        var short_byte byte = '0'

        //Check if there are still bits left in the short string
        if i - (len(longer)-len(shorter)) >= 0 {
            short_byte = shorter[i - (len(longer)-len(shorter))]
        }

        //add the two bytes
        sum_byte, carry_byte = sum_bits(long_byte, short_byte, carry_byte)

        //add new bit to sum string
        sum = string(sum_byte) + sum
    }

    //Check for extra carry
    if carry_byte == '1' {sum = string(carry_byte) + sum }

    return sum
}


//takes 3 bytes (2 binary 1 carry) returning the sum and carry 
func sum_bits(x byte, y byte, carry byte) (byte, byte) {
    
    //and 
    if x == '1' && y == '1' {
        return carry, '1' 
    }

    //or
    if x == '1' || y == '1' {
        if carry == '1' {
            return '0', '1'
        }

        return '1', '0'
    }

    //carry
    if carry == '1'{
        return '1', '0'
    }

    return '0', '0'
}
