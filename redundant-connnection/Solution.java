import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

class Solution {

  public static void main(String[] args) {
    int[][] edges = new int[][] {
      new int[]{1,5},
      new int[]{1,2},
      new int[]{2,3},
      new int[]{3,4},
      new int[]{1,4},
    };
    // int[][] edges = new int[][] {
    //   new int[]{1,2},
    //   new int[]{1,3},
    //   new int[]{2,4},
    //   new int[]{2,8},
    //   new int[]{4,8},
    //   new int[]{4,6},
    //   new int[]{3,5},
    //   new int[]{5,7},
    // };
    Solution solution = new Solution();
    int[] edge = solution.findRedundantConnection(edges);
    System.out.println(Arrays.toString(edge));
  }

  public int[] findRedundantConnection(int[][] edges) {
    Graph graph = Graph.build(edges);
    for (int i = edges.length - 1; i >= 0; i--) {
      int start = edges[i][0];
      int end = edges[i][1];
      graph.removeEdge(start, end);
      boolean isLoop = graph.checkLoop();
      if (!isLoop) {
        return edges[i];
      }
      graph.addEdge(start, end);
    }
    return new int[0];
  }

  static class Graph {
    Map<Integer, Set<Integer>> nodes = new HashMap<>();

    static Graph build(int[][] edges) {
      Graph graph = new Graph();
      for (int i = 0; i < edges.length; i++) {
        int start = edges[i][0];
        int end = edges[i][1];
        graph.addEdge(start, end);
      }
      return graph;
    }

    boolean checkLoop() {
      Set<Integer> visited = new HashSet<>();
      while (visited.size() != this.nodes.size()) {
        Integer start = chooseStart(visited);
        visited.add(start);
        boolean isLoop = checkLoop(-1, start, visited);
        if (isLoop) {
          return true;
        }
      }
      return false;
    }

    private boolean checkLoop(Integer from, Integer to, Set<Integer> visited) {
      Set<Integer> neighbors = this.nodes.get(to).stream()
        .filter(n -> !n.equals(from))
        .collect(Collectors.toSet());
      Set<Integer> unvisitedNeighbors = neighbors.stream().filter(n -> !visited.contains(n)).collect(Collectors.toSet());
      if (neighbors.size() > 0 && unvisitedNeighbors.size() == 0) {
        return true;
      }
      for (Integer neighbor: unvisitedNeighbors) {
        visited.add(neighbor);
        boolean isLoop = checkLoop(to, neighbor, visited);
        if (isLoop) {
          return true;
        }
      }
      return false;
    }

    private int chooseStart(Set<Integer> visited) {
      for (Integer key: this.nodes.keySet()) {
        if (!visited.contains(key)) {
          return key;
        }
      }
      return -1;
    }

    void addEdge(int start, int end) {
      Set<Integer> neighbors = this.nodes.getOrDefault(start, new HashSet<>());
      neighbors.add(end);
      this.nodes.put(start, neighbors);
      neighbors = this.nodes.getOrDefault(end, new HashSet<>());
      neighbors.add(start);
      this.nodes.put(end, neighbors);
    }

    void removeEdge(int start, int end) {
      Set<Integer> neighbors = this.nodes.getOrDefault(start, new HashSet<>());
      neighbors.remove(end);
      if (neighbors.size() == 0) {
        this.nodes.remove(start);
      }
      neighbors = this.nodes.getOrDefault(end, new HashSet<>());
      neighbors.remove(start);
      if (neighbors.size() == 0) {
        this.nodes.remove(end);
      }
    }
  }
}

