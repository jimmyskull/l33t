void moveZeroes(int* nums, int numsSize)
{
	    int i = 0, j = 0;
	    while (j < numsSize) {
		        nums[i] = nums[j++];
		        i += nums[i] != 0;
	    }
	    memset(nums + i, 0, (numsSize - i) * sizeof(int));
}