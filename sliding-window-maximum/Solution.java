import java.util.ArrayDeque;
import java.util.Arrays;
import java.util.Deque;

class Solution {

  public static void main(String[] args) {
    int[] a = new int[] {1,3,-1,-3,5,3,6,7};
    int k = 3;
    Solution solution = new Solution();
    int[] results = solution.maxSlidingWindow(a, 9);
    System.out.println(Arrays.toString(results));
  }

  public int[] maxSlidingWindow(int[] a, int k) {		
    int n = a.length;
    int[] results = new int[n - k + 1];
    int index = 0;
    Deque<Integer> queue = new ArrayDeque<>();
    for (int i = 0; i < n; i++) {
      while (!queue.isEmpty() && queue.getFirst() < i - k + 1) {
        queue.removeFirst();
      }
      while (!queue.isEmpty() && a[queue.getLast()] <= a[i]) {
        queue.removeLast();
      }
      queue.offerLast(i);
      if (i + 1 >= k) {
        results[index++] = a[queue.getFirst()];
      }
    }
    return results;
  }

}
