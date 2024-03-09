func isPalindrome(s string) bool {
    var left_index int = 0
    var right_index int = len(s)-1

    for left_index < right_index {
        var left_ascii int = ascii_validate(int(s[left_index]))
        var right_ascii int = ascii_validate(int(s[right_index]))

        //Move pointer if char is not valid
        for left_ascii == 0 && left_index < right_index { 
            left_index++
            left_ascii = ascii_validate(int(s[left_index]))
        }
        for right_ascii == 0 && right_index > left_index{
            right_index--
            right_ascii = ascii_validate(int(s[right_index]))
        }
        
        //Check for valid palindrone
        if left_ascii != right_ascii{
            return false
        }

        left_index++
        right_index--
    }

    return true
}

func ascii_validate(x int) int {
    //capital
    if x >= 65 && x <= 90 {
        return x+32
    }

    //lowercase
    if x >= 97 && x <= 122 {
        return x
    }

    if x >= 48 && x <= 57{
        return x
    }

    return 0
}
