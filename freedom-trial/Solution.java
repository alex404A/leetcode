import java.util.List;
import java.util.Map;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;

public class Solution {

  public static void main(String[] args) {
    Solution solution = new Solution();
    int steps = solution.findRotateSteps("edcba", "abcde"); 
    System.out.println(steps);
  }

  public int findRotateSteps(String ring, String key) {
    if (key.length() == 0 || ring.length() == 0) {
      return 0;
    }
    int ringLength = ring.length(); 
    Map<String, List<Integer>> keyLocations = transformRing(ring);
    String prevKey = key.substring(0, 1);
    int[][] prevSteps = initFirstSteps(keyLocations, prevKey, ringLength);
    for (int i = 1; i < key.length(); i++) {
      String currentKey = key.substring(i, i + 1);
      prevSteps = updateLatestSteps(keyLocations, currentKey, prevKey, prevSteps, ringLength);
    }
    return key.length() + Arrays.stream(prevSteps)
      .map(elm -> elm[1])
      .reduce(Integer.MAX_VALUE, (min, val) -> Math.min(min, val));
  }

  private int[][] updateLatestSteps(Map<String, List<Integer>>keyLocations, String currentKey,
      String lastKey, int[][] prevSteps, int ringLength) {
    List<Integer> currentKeyLocations = keyLocations.get(currentKey);
    int[][] currentSteps = new int[currentKeyLocations.size()][2];
    for (int i = 0; i < currentKeyLocations.size(); i++) {
      int min = Integer.MAX_VALUE;
      for (int j = 0; j < prevSteps.length; j++) {
        min = Math.min(min, prevSteps[j][1] + getMinPath(currentKeyLocations.get(i), prevSteps[j][0], ringLength)); 
      }
      currentSteps[i][0] = currentKeyLocations.get(i);
      currentSteps[i][1] = min;
    }
    return currentSteps;
  }

  private int[][] initFirstSteps(Map<String, List<Integer>> keyLocations, String firstKey, int ringLength) {
    List<Integer> firstKeyLocations = keyLocations.get(firstKey);
    int[][] firstSteps = new int[firstKeyLocations.size()][2];
    for (int i = 0; i < firstKeyLocations.size(); i++) {
      firstSteps[i][0] = firstKeyLocations.get(i);
      firstSteps[i][1] = getMinPath(0, firstKeyLocations.get(i), ringLength);
    }
    return firstSteps;
  }

  private int getMinPath(int i, int j, int length) {
    if (i == j) {
      return 0;
    }
    int first = Math.abs(i - j);
    int small = i < j ? i : j;
    int big = i >= j ? i : j;
    int second = Math.abs(small + length - big);
    return first < second ? first : second;
  }

  private Map<String, List<Integer>> transformRing(String ring) {
    Map<String, List<Integer>> keyLocations = new HashMap<>();
    for (int i = 0; i < ring.length(); i++) {
      String key = ring.substring(i, i + 1);
      List<Integer> locations = keyLocations.getOrDefault(key, new ArrayList<>());
      locations.add(i);
      keyLocations.put(key, locations);
    }
    return keyLocations;
  }
}
