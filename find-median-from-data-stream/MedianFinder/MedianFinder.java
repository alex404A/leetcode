import java.util.PriorityQueue;
import java.util.Queue;

class MedianFinder {

  private Queue<Integer> large = new PriorityQueue<>();
  private Queue<Integer> small = new PriorityQueue<>(5,(a,b) -> a > b ? -1 : a == b ? 0 : 1);

  public static void main(String[] args) {
    MedianFinder finder = new MedianFinder();
    finder.addNum(1); 
    System.out.println(finder.findMedian());
    finder.addNum(2); 
    System.out.println(finder.findMedian());
    finder.addNum(3); 
    System.out.println(finder.findMedian());
    finder.addNum(1); 
    System.out.println(finder.findMedian());
    finder.addNum(2); 
    System.out.println(finder.findMedian());
    finder.addNum(3); 
    System.out.println(finder.findMedian());
  }

  /** initialize your data structure here. */
  public MedianFinder() {
  }

  public void addNum(int num) {
    if (large.size() == 0) {
      large.add(num);
      return;
    }
    if (large.peek() <= num) {
      large.add(num);
      if (large.size() > small.size() + 1) {
        int min = large.poll();
        small.add(min);
      }
    } else {
      small.add(num);
      if (small.size() > large.size()) {
        int max = small.poll();
        large.add(max);
      }
    }
  }

  public double findMedian() {
    if (large.size() > small.size()) {
      return large.peek();
    } else {
      return (large.peek() + small.peek() + 0.0) / 2;
    }
  }
}