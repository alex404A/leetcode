import java.util.Map;
import java.util.TreeMap;

class Solution {
  public boolean isNStraightHand(int[] hand, int W) {
    if (hand.length % W != 0) {
      return false;
    }
    Map<Integer, Integer> container = new TreeMap<>();
    for (int i = 0; i < hand.length; i++) {
      container.put(hand[i], container.getOrDefault(hand[i], 0) + 1); 
    }
    for (Integer key: container.keySet()) {
      if (container.get(key) == 0) {
        continue;
      }
      while (container.get(key) != 0) {
        if (!satisfy(container, key, W)) {
          return false;
        }
        deduct(container, key, W);
      }
    }
    return true;
  }

  private boolean satisfy(Map<Integer, Integer> container, Integer key, int W) {
    for (int i = 1; i < W; i++) {
      if (container.get(key + i) == null || container.get(key + i) == 0) {
        return false;
      }
    }
    return true;
  }

  private void deduct(Map<Integer, Integer> container, Integer key, int W) {
    for (int i = 0; i < W; i++) {
      container.put(key + i, container.get(key + i) - 1);
    }
  }
}
