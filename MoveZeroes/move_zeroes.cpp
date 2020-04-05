class Solution {
 public:
  void moveZeroes(std::vector<int>& nums) {
    unsigned size = nums.size(), i = 0, j = 0;
 
    while (j < size) {
      nums[i] = nums[j++];
      i += nums[i] != 0;
    }
    memset(nums.data() + i, 0, (size - i) * sizeof(int));
  }
};