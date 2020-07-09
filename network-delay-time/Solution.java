import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.PriorityQueue;
import java.util.Queue;
import java.util.Set;

class Solution {

  public static void main(String[] args) {

    int[][] times = new int[][] {
      new int[]{1,2,1},
      new int[]{2,3,1},
      new int[]{1,3,4},
      new int[]{3,4,4},
      new int[]{5,6,4},
    };
    int N = 6;
    int K = 1;
    Solution solution = new Solution();
    int max = solution.networkDelayTime(times, N, K);
    System.out.println(max);
  }

  public int networkDelayTime(int[][] times, int N, int K) {
    Graph graph = Graph.buildGraph(times);   
    Set<Integer> visited = new HashSet<>();
    Queue<QueueItem> queue = new PriorityQueue<>();
    queue.add(new QueueItem(K, 0));
    Map<Integer, Integer> pathContainer = buildPathContainer(N);
    while (!queue.isEmpty()) {
      QueueItem item = queue.remove();
      if (visited.contains(item.number)) {
        continue;
      }
      visited.add(item.number);
      Map<Integer,Integer> neighbors = graph.getNeighbors(item.number);
      int min = Math.min(item.time, pathContainer.get(item.number));
      pathContainer.put(item.number, min);
      for (Map.Entry<Integer,Integer> neighbor : neighbors.entrySet()) {
        int number = neighbor.getKey();
        int time = neighbor.getValue();
        QueueItem nextItem = new QueueItem(number, item.time + time);
        queue.add(nextItem);
      }
    }
    if (visited.size() == N) {
      return decideFurthest(pathContainer);
    } else {
      return -1;
    }
  }

  private Map<Integer,Integer> buildPathContainer(int N) {
    Map<Integer,Integer> container = new HashMap<>(N);
    for (int i = 1; i <= N; i++) {
      container.put(i, Integer.MAX_VALUE);
    }
    return container;
  }

  private int decideFurthest(Map<Integer,Integer> container) {
    int max = 0;
    for (Integer time: container.values()) {
      max = Math.max(time, max);
    }
    return max;
  }

  static class Graph {
    Map<Integer, Map<Integer, Integer>> nodes = new HashMap<>();

    static Graph buildGraph(int[][] times) {
      Graph graph = new Graph();
      for (int i = 0; i < times.length; i++) {
        int[] item = times[i];
        Map<Integer, Integer> neighbors = graph.nodes.getOrDefault(item[0], new HashMap<>());
        neighbors.put(item[1], item[2]);
        graph.nodes.put(item[0], neighbors);
      }
      return graph;
    }

    Map<Integer,Integer> getNeighbors(int number) {
      return this.nodes.getOrDefault(number, new HashMap<>());
    }
  }

  static class QueueItem implements Comparable<QueueItem> {
    int number;
    int time;

    QueueItem(int number, int time) {
      this.number = number;
      this.time = time;
    }

    @Override
    public int compareTo(QueueItem other) {
      if (this.time > other.time) {
        return 1;
      } else if (this.time == other.time) {
        return 0;
      } else {
        return -1;
      }
    }
  }

}
