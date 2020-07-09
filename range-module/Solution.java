import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;

class RangeModule {

  public static void main(String[] args) {
    RangeModule rangeModule = new RangeModule(); 
    boolean result; 
    rangeModule.addRange(1, 4);
    System.out.println(rangeModule.container);
    rangeModule.addRange(6, 10);
    System.out.println(rangeModule.container);
    // rangeModule.removeRange(31, 32);
    // System.out.println(rangeModule.container);
    // result = rangeModule.queryRange(10, 14);
    // System.out.println(result);
    result = rangeModule.queryRange(4, 6);
    System.out.println(result);
    // result = rangeModule.queryRange(16, 17);
    // System.out.println(result);
  }
  
  private TreeMap<Integer, Integer> container = new TreeMap<>();

  public RangeModule() {

  }

  public void addRange(int left, int right) {
    if (right <= left) {
      return;
    }
    this.container.put(left, this.container.getOrDefault(left, 0) + 1);
    this.container.put(right, this.container.getOrDefault(right, 0) - 1);
    Integer count = 0;
    boolean flag = false;
    List<Integer> merged = new LinkedList<>();
    for (Integer key: this.container.keySet()) {
      if (this.container.get(key) == 2 ||
          count == 1 && (this.container.get(key) == 0 || this.container.get(key) == 1)) {
        flag = true;
      }
      count += container.get(key);
      if (flag && count == 0) {
        if (this.container.get(key) == -2) {
          this.container.put(key, -1);
        }
        break;
      }
      if (flag) {
        merged.add(key);
      }
    }
    for (Integer key: merged) {
      if (this.container.get(key) == 2) {
        this.container.put(key, 1);
      } else {
        this.container.remove(key);
      }
    }
  }

  public boolean queryRange(Integer left, Integer right) {
    if (right <= left) {
      return true;
    }
    Integer start = this.container.floorKey(left);
    Integer end = this.container.floorKey(right);
    System.out.println(start);
    System.out.println(end);
    if (start == null) {
      return false;
    }
    if (start == end && this.container.get(start) == 1) {
      return true;
    }
    return this.container.get(end) == -1 && right == end && right.equals(this.container.higherKey(left));
  }

  public void removeRange(int left, int right) {
    if (right <= left) {
      return;
    }
    Integer start = this.container.floorKey(left);
    Integer end = this.container.floorKey(right);
    if (start == null && end == null) {
      return;
    }
    if (start == end && this.container.get(start) == -1) {
      return;
    }
    List<Integer> merged = new LinkedList<>();
    Map<Integer, Integer> sub = new HashMap<>();
    boolean flag = start == null;
    for (Integer key: this.container.keySet()) {
      if (key == start) {
        if (start == left && this.container.get(key) == 1) {
          merged.add(key);
        }
        if (start != left && this.container.get(key) == 1) {
          sub.put(left, -1);
        }
        flag = true;
      }
      if (flag && key != start) {
        merged.add(key);
      }
      if (key == end) {
        if (this.container.get(key) == -1) {
          merged.add(key);
        }
        if (end == right && this.container.get(key) == 1) {
          merged.remove(merged.size() - 1);
        }
        if (end != right && this.container.get(key) == 1) {
          sub.put(right, 1);
        }
        break;
      }
    }
    for (Integer key: merged) {
      this.container.remove(key);
    }
    for (Integer key: sub.keySet()) {
      this.container.put(key, sub.get(key));
    }
  }
}
