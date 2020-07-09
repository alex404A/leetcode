class Solution {

  public static void main(String[] args) {
    ListNode fifth = new ListNode(5);
    ListNode fourth = new ListNode(4, fifth);
    ListNode third = new ListNode(3, fourth);
    ListNode second = new ListNode(2, third);
    ListNode first = new ListNode(1, second);
    Solution solution = new Solution();
    solution.printLinkedList(first);
    ListNode head = solution.reverseList(first);
    solution.printLinkedList(head);
  }

  public void printLinkedList(ListNode head) {
    StringBuilder sb = new StringBuilder();
    while (head != null) {
      sb.append(head.val);
      sb.append(" ");
      head = head.next;
    }
    System.out.println(sb.toString().trim());
  }

  public ListNode reverseList(ListNode head) {
    if (head == null) {
      return null;
    }
    ListNode next = head.next;
    head.next = null;
    while (next != null) {
      ListNode tmp = next.next;
      next.next = head;
      head = next;
      next = tmp;
    }
    return head;
  }
}

class ListNode {
  int val;
  ListNode next;
  ListNode() {}
  ListNode(int val) { this.val = val; }
  ListNode(int val, ListNode next) { this.val = val; this.next = next; }
}

