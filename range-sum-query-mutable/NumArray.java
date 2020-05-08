class NumArray {

  private int[] nums;
  private SegmentTreeNode root;

  static class SegmentTreeNode {
    int start;
    int end;
    SegmentTreeNode left;
    SegmentTreeNode right;
    int sum;

    SegmentTreeNode(int start, int end, SegmentTreeNode left, SegmentTreeNode right) {
      this.sum = 0;
      this.start = start;
      this.end = end;
      this.left = left;
      this.right = right;
      if (this.left != null) {
        this.sum += left.sum;
      }
      if (this.right != null) {
        this.sum += right.sum;
      }
    }

    int sumRange(int start, int end) {
      if (this.start == start && this.end == end) {
        return this.sum;
      }
      int mid = this.start + (this.end - this.start) / 2;
      if (end <= mid) {
        return this.left.sumRange(start, end);
      } else if (start >= mid + 1) {
        return this.right.sumRange(start, end);
      } else {
        return this.left.sumRange(start, mid) + this.right.sumRange(mid + 1, end);
      }
    }

    public void update(int index, int diff) {
      this.sum += diff;
      if (this.left != null && this.left.start <= index && this.left.end >= index) {
        this.left.update(index, diff);
      }
      if (this.right != null && this.right.start <= index && this.right.end >= index) {
        this.right.update(index, diff);
      }
    }
  }

  public static void main(String[] args) {
    int[] nums = new int[]{7,2,7,2,0};
    NumArray obj = new NumArray(nums);
    obj.update(4,6);
    obj.update(0,2);
    obj.update(0,9);
    System.out.println(obj.sumRange(4,4));
    obj.update(3,8);
    System.out.println(obj.sumRange(0,4));
    obj.update(4,1);
    System.out.println(obj.sumRange(0,3));
    System.out.println(obj.sumRange(0,4));
    obj.update(0,4);
  }

  public NumArray(int[] nums) {
    this.nums = nums;
    this.root = buildTree(nums, 0, nums.length - 1);
  }

  public void update(int i, int val) {
    int original = this.nums[i];
    int diff = val - original;
    this.root.update(i, diff);
    this.nums[i] = val;
  }

  public int sumRange(int i, int j) {
    return root.sumRange(i, j);
  }

  private SegmentTreeNode buildTree(int[] nums, int start, int end) {
    if (start > end) {
      return null;
    } else if (start == end) {
      SegmentTreeNode node = new SegmentTreeNode(start, end, null, null);
      node.sum = nums[start];
      return node;
    }
    int mid = start + (end - start) / 2;
    SegmentTreeNode left = buildTree(nums, start, mid);
    SegmentTreeNode right = buildTree(nums, mid + 1, end);
    return new SegmentTreeNode(start, end, left, right);
  }

}
