import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.PriorityQueue;

class Solution {

    class Node implements Comparable<Node> {
        int x;
        int y;
        boolean isStart;

        public Node(int x, int y, boolean isStart) {
            this.x = x;
            this.y = y;
            this.isStart = isStart;
        }

        @Override
        public int compareTo(Node other) {
            if (this.x < other.x) {
                return -1;
            } else if (this.x > other.x) {
                return 1;
            }
            if (this.y == other.y) {
                return this.isStart ? -1 : 1;
            }
            if (this.isStart && other.isStart) {
                return this.y > other.y ? -1 : 1;
            } else if (!this.isStart && !other.isStart) {
                return this.y > other.y ? 1 : -1;
            } else {
                return this.isStart ? -1 : 1;
            }
        }
    }

    public static void main(String[] args) {
        int[][] buildings = new int[][]{
            // new int[]{0, 2, 3},
            // new int[]{2, 5, 3},
            new int[]{2, 9, 10},
            new int[]{9, 12, 15},
            // new int[]{3, 7, 15},
            // new int[]{5, 12, 12},
            // new int[]{15, 20, 10},
            // new int[]{19, 24, 8},

            //new int[]{0,5,7},
            //new int[]{5,10,7},
            //new int[]{5,10,12},
            //new int[]{10,15,7},
            //new int[]{15,20,7},
            //new int[]{15,20,12},
            //new int[]{20,25,7},

        };
        Solution solution = new Solution();
        List<List<Integer>> pts = solution.getSkyline(buildings);
        System.out.println(pts);
    }

    public List<List<Integer>> getSkyline(int[][] buildings) {
        List<List<Integer>> results = new ArrayList<>();
        if (buildings.length == 0) {
            return results;
        }
        List<Node> nodes = buildNodes(buildings);
        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
        pq.add(0);
        for (Node node: nodes) {
            if (node.isStart) {
                Integer preMax = pq.peek();
                pq.add(node.y);
                Integer curMax = pq.peek();
                if (preMax < curMax) {
                    results.add(List.of(node.x, curMax));
                }
            } else {
                Integer preMax = pq.peek();
                pq.remove(node.y);
                Integer curMax = pq.peek();
                if (curMax < preMax) {
                    results.add(List.of(node.x, curMax));
                }
            }
        }
        return results;
    }

    private List<Node> buildNodes(int[][] buildings) {
        List<Node> nodes = new ArrayList<>();
        for (int i = 0; i < buildings.length; i++) {
            int[] building = buildings[i];
            nodes.add(new Node(building[0], building[2], true));
            nodes.add(new Node(building[1], building[2], false));
        }
        Collections.sort(nodes);
        return nodes;
    }
}