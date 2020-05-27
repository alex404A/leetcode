import java.util.Arrays;
import java.util.PriorityQueue;
import java.util.Queue;
import java.util.Set;
import java.util.stream.Collectors;

public class Solution {
   
  public static void main(String[] args) {
    Solution solution = new Solution();
    int[][] points = new int[][] { new int[] { 0, 1 }, new int[] { 1, 0 }, };
    int[][] results = solution.kClosest(points, 2);
    // System.out.println(results);
  }


  public int[][] kClosest(int[][] points, int k) {
      if (k == 0) {
          return new int[0][];
      }
              if (points.length == k) {
          return points;
      }
      if (k <= points.length / 2) {
          return findSmall(points, k);
      } else {
          return findLarge(points, k);
      }
  }

  private int[][] findLarge(int[][] points, int K) {
      Queue<Point> heap = new PriorityQueue<>();
      int size = points.length - K;
      for (int i = 0; i < points.length; i++) {
          if (heap.size() < size) {
              heap.add(new Point(points[i], i, calcDistance(points[i])));
          } else {
              int distance = calcDistance(points[i]);
              if (heap.peek().distance < distance) {
                  heap.poll();
                  heap.add(new Point(points[i], i, distance));
              }
          }
      }
      Point[] ps = heap.toArray(new Point[0]);
      Set<Integer> indexes = Arrays.stream(ps).map(p -> p.index).collect(Collectors.toSet());
      int[][] results = new int[K][];
      int cnt = 0;
      for (int i = 0; i < points.length; i++) {
          if (!indexes.contains(i)) {
              results[cnt] = points[i];
              cnt++;
          }
      }
      return results;
  }

  private int[][] findSmall(int[][] points, int K) {
      Queue<Point> heap = new PriorityQueue<>(Collections.reverseOrder());
      for (int i = 0; i < points.length; i++) {
          if (heap.size() < K) {
              heap.add(new Point(points[i], i, calcDistance(points[i])));
          } else {
              int distance = calcDistance(points[i]);
              if (heap.peek().distance > distance) {
                  heap.poll();
                  heap.add(new Point(points[i], i, distance));
              }
          }
      }
      Point[] ps = heap.toArray(new Point[0]);
      int[][] results = new int[ps.length][];
      for (int i = 0; i < ps.length; i++) {
          results[i] = ps[i].p;
      }
      return results;
  }

  private int calcDistance(int[] p) {
      return p[0] * p[0] + p[1] * p[1];
  }

  static class Point implements Comparable<Point> {
      private int[] p;
      private int index;
      private int distance;

      Point(int[] p, int index, int distance) {
          this.p = p;
          this.index = index;
          this.distance = distance;
      }

      @Override
      public int compareTo(Point point) {
          return (this.distance - point.distance);
      }

      @Override
      public boolean equals(Object o) {
          if (!(o instanceof Point)) {
              return false;
          }
          Point other = (Point) o;
          if (other == this) {
              return true;
          } else {
              return this.distance == other.distance;
          }
      }
  }
}