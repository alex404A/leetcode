import java.util.HashSet;
import java.util.Set;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

class Solution {

  public static void main(String[] args) {
    Solution solution = new Solution(); 
    int[][] M = new int[][]{
      new int[]{1,1,0,0,0,0,0,1,0,0,0,0,0,0,0},
      new int[]{1,1,0,0,0,0,0,0,0,0,0,0,0,0,0},
      new int[]{0,0,1,0,0,0,0,0,0,0,0,0,0,0,0},
      new int[]{0,0,0,1,0,1,1,0,0,0,0,0,0,0,0},
      new int[]{0,0,0,0,1,0,0,0,0,1,1,0,0,0,0},
      new int[]{0,0,0,1,0,1,0,0,0,0,1,0,0,0,0},
      new int[]{0,0,0,1,0,0,1,0,1,0,0,0,0,1,0},
      new int[]{1,0,0,0,0,0,0,1,1,0,0,0,0,0,0},
      new int[]{0,0,0,0,0,0,1,1,1,0,0,0,0,1,0},
      new int[]{0,0,0,0,1,0,0,0,0,1,0,1,0,0,1},
      new int[]{0,0,0,0,1,1,0,0,0,0,1,1,0,0,0},
      new int[]{0,0,0,0,0,0,0,0,0,1,1,1,0,0,0},
      new int[]{0,0,0,0,0,0,0,0,0,0,0,0,1,0,0},
      new int[]{0,0,0,0,0,0,1,0,1,0,0,0,0,1,0},
      new int[]{0,0,0,0,0,0,0,0,0,1,0,0,0,0,1},
    };
    int result = solution.findCircleNum(M);
    System.out.println(result);
  }

  public int findCircleNum(int[][] M) {
    int rows = M.length;
    int cols = M[0].length;
    int[] chains = new int[rows];
    for (int i = 0; i < rows; i++) {
      chains[i] = i;
    }
    for (int i = 0; i < rows; i++) {
      for (int j = i + 1; j < cols; j++) {
        if (M[i][j] == 1) {
          int rootI = findRoot(i, chains);
          int rootJ = findRoot(j, chains);
          chains[rootJ] = rootI;
        }
      }
      System.out.println(Arrays.toString(chains));
    }
    int result = 0;
    for (int i = 0; i < rows; i++) {
      if (chains[i] == i) {
        result++;
      }
    }
    return result;
  }

  private int findRoot(int x, int[] chains) {
    while (x != chains[x]) {
      x = chains[x];
    }
    return x;
  }

  public int findCircleNum2(int[][] M) {
    if (M.length == 0 || M[0].length == 0) {
      return 0;
    }
    int rows = M.length;
    int cols = M[0].length;
    List<Set<Integer>> circles = initCircles(M.length);
    for (int i = 0; i < rows; i++) {
      Set<Integer> cur = circles.get(i);
      for (int j = i; j < cols; j++) {
        if (M[i][j] == 1) {
          cur.add(j);
          replace(circles, cur, i, j);
        }
      }
      if (i == 5 || i == 6 || i == 7) {
        System.out.println(i);
        System.out.println(circles);
      }
    }
    System.out.println(circles);
    return count(circles);
  }

  private int count(List<Set<Integer>> circles) {
    int result = 0;
    Set<Integer> visited = new HashSet<>();
    for (Set<Integer> circle: circles) {
      boolean isVisited = circle.stream().anyMatch(i -> visited.contains(i));
      if (!isVisited) {
        result++;
        System.out.println(circle);
      }
      visited.addAll(circle);
      if (visited.size() == circles.size()) {
        break;
      }
    }
    return result;
  }

  private void replace(List<Set<Integer>> circles, Set<Integer> cur, int i, int j) {
    if (circles.get(j).size() == 0) {
      circles.set(j, cur);
    } else {
      circles.get(j).addAll(cur);
      circles.get(i).addAll(circles.get(j));
    }
  }

  private List<Set<Integer>> initCircles(int numbers) {
    List<Set<Integer>> circles = new ArrayList<Set<Integer>>(numbers);
    for (int i = 0; i < numbers; i++) {
      circles.add(new HashSet<>());
    }
    return circles;
  }
}
