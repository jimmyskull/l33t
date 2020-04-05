public class Solution {
    public void MoveZeroes(int[] nums) {
        int i = 0;
        int j = 0;
        int n = nums.Length;
        while (j < n) {
            nums[i] = nums[j];
            j++;
            if (nums[i] != 0)
                i++;
        }
        while (i < n) {
            nums[i] = 0;
            i++;
        }
    }
}