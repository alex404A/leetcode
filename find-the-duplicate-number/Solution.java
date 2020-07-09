import java.util.BitSet;

class Solution {

  public static void main(String[] args) {
    Solution solution = new Solution(); 
    int[] nums = new int[]{1,3,4,2,2,};
    int result = solution.findDuplicate(nums);
    System.out.println(result);
  }

  public int findDuplicate(int[] nums) {
    int n = nums.length - 1;    
    BitSet set = new BitSet();
    for (int i = 0; i < n + 1; i++) {
      boolean isExist = set.get(nums[i]);
      if (isExist) {
        return nums[i];
      } else {
        set.set(nums[i]);
      }
    }
    return -1;
  }
}
