func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    //Could use m-1 and n-1 in the for loop but this is easier to understand
    var mIndex int = m-1
    var nIndex int = n-1

    //Loop through both m+n-1 times starting at end of nums1 and nums2 check which is greater and store in decreaing order at the end if nums1
    for i := m+n-1; i>=0; i-- {

        //Check If either index is out of range
        if mIndex < 0 {
            nums1[i] = nums2[nIndex]
            nIndex--
        }else
        if nIndex < 0 {
            nums1[i] = nums1[mIndex]
            mIndex--
        }else

        // Check which value is bigger
        if nums1[mIndex] > nums2[nIndex]{
            nums1[i] = nums1[mIndex]
            mIndex--
        }else{
            nums1[i] = nums2[nIndex]
            nIndex--
        }
    }
}

