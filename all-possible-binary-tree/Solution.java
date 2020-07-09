import java.util.LinkedList;
import java.util.List;

class Solution {

  public static void main(String[] args) {
    Solution solution = new Solution(); 
    List<TreeNode> results = solution.allPossibleFBT(7);
    System.out.println(results.size());
  }

  public List<TreeNode> allPossibleFBT(int N) {
    if (N % 2 == 0) {
      return new LinkedList<>();
    }
    return possibleFBT(N);
  }

  private List<TreeNode> possibleFBT(int N) {
    if (N == 1) {
      return List.of(new TreeNode(0));
    } else if (N == 3) {
      return List.of(new TreeNode(0, new TreeNode(0), new TreeNode(0)));
    } else {
      List<TreeNode> results = new LinkedList<>();
      for (int i = 1; i < N - 1; i += 2) {
        if ((N - 1 - i) % 2 != 0) {
          List<TreeNode> left = possibleFBT(i);
          List<TreeNode> right = possibleFBT(N - 1 - i);
          results.addAll(combine(left, right));
        }
      }
      return results;
    }
  }

  private List<TreeNode> combine(List<TreeNode> left, List<TreeNode> right) {
    List<TreeNode> result = new LinkedList<>();
    for (TreeNode a: left) {
      for (TreeNode b: right) {
        TreeNode parent = new TreeNode(0, a, b);
        result.add(parent);
      }
    }
    return result;
  }
}

public class TreeNode {
  int val;
  TreeNode left;
  TreeNode right;
  TreeNode() {}
  TreeNode(int val) { this.val = val; }
  TreeNode(int val, TreeNode left, TreeNode right) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}

