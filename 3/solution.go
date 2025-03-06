/*
   This solution approach uses a sliding window. We begin with 2 pointers(left and right) at the
   start of the string and an empty [byte]int map. Wher the key is a byte and the value is an integer.
   We then loop through the string moving the right pointer. When we come across an char not in the map
   we add it along with it's index as the value. If we come across a byte that we have already found we
   update the index for that byte in the map to the new byte we have found. The left pointer is also updated
   to the index after the duplicate byte we found.

   Each iteration we also calcuate the current count and update the longest count if the current count is
   greater.

*/

func lengthOfLongestSubstring(s string) int {

	// create the byte map, left pointer, and longest count
	charSet := make(map[byte]int)
	leftPointer, longestCount := 0, 0

	// for loop through the string
	for rightPointer := 0; rightPointer <= len(s)-1; rightPointer++ {
		currByte := s[rightPointer]
		index, found := charSet[currByte]
		// If we find a duplicate and the duplicate is still in our window
		if found && index >= leftPointer {
			//Move left pointer to the byte after the duplicate byte
			leftPointer = index + 1
		}

		// Add the right byte to the charSet or overwrite an exisiting byte with a duplicate index
		charSet[currByte] = rightPointer

		// calculate the current max length
		newMax := rightPointer - leftPointer + 1

		// update the longest count if the current max is greater
		if newMax > longestCount {
			longestCount = newMax
		}

	}

	return longestCount

}
