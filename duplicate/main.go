package duplicate

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	var tmpmap = make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := tmpmap[nums[i]]; ok {
			return true
		}
		tmpmap[nums[i]] = struct{}{}
	}
	return false
}
