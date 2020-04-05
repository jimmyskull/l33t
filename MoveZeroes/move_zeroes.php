class Solution {
 
    function moveZeroes(&$nums) {
        $j = 0;
        for($i = 0; $i < count($nums);) {
            if ($j < count($nums)) {
                $nums[$i] = $nums[$j];
                if ($nums[$i] != 0)
                    $i++;
                $j++;
            } else {
                $nums[$i] = 0;
                $i++;
            }
        }
    }
}