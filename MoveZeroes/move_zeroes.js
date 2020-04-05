var moveZeroes = function(nums) {
    var i = 0;
    var j = 0;
    var n = nums.length;
    while (j < n) {
        nums[i] = nums[j];
        j++;
        if (nums[i] != 0) i++;
    }
    while (i < n) {
        nums[i] = 0;
        i++;
    }
};