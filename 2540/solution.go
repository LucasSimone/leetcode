func getCommon(nums1 []int, nums2 []int) int {
    var nums1_index int = 0
    var nums2_index int = 0

    for nums1_index < len(nums1) && nums2_index < len(nums2){
        if nums1[nums1_index] == nums2[nums2_index] {
            return nums1[nums1_index]
        }

        if nums1[nums1_index] < nums2[nums2_index] {
            nums1_index++
        }else {
            nums2_index++
        }
    }

    return -1
}
