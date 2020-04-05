impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut i = 0;
        let mut j = 0;
        let n = nums.len();
        while j < n {
            nums[i] = nums[j];
            j += 1;
            if nums[i] != 0 {
                i += 1;
            }
        }
        while i < n {
            nums[i] = 0;
            i += 1;
        }
    }
}