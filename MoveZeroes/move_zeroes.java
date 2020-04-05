class Solution {
    public void moveZeroes(int[] nums) {
        int i = 0;
        int j = 0;
        int n = nums.length;
        while (j < n) {
            nums[i] = nums[j];
            j++;
            if (nums[i] != 0) {
                i++;
            }
        }
        while (i < n) {
            nums[i] = 0;
            i++;
        }
    }
}