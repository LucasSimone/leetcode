func canConstruct(ransomNote string, magazine string) bool {
    //If ransom note is great there aren't enough letters
    if len(ransomNote) > len(magazine) {
        return false
    }

    //2 maps aren't needed 
    //Alternatively you can create only one map then loop though the other string and subtract
    ransom_count := make(map[byte]int)
    magazine_count := make(map[byte]int)

    //Make Hashmap of characters
    for i:=0; i<len(magazine); i++ {
        if i < len(ransomNote) {
            ransom_count[ransomNote[i]]++
        }
        magazine_count[magazine[i]]++
    }

    //Check ransom count excceds magazine for any character
    for k, v := range ransom_count { 
        if v > magazine_count[k] {
            return false
        }
    }
    
    return true

}  
