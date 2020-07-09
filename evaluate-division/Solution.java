import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Set;

class Solution {

  public static void main(String[] args) {
    Solution solution = new Solution(); 
    List<List<String>> equations = new ArrayList<>();
    equations.add(List.of("a", "b"));
    equations.add(List.of("b", "c"));
    equations.add(List.of("a", "d"));
    equations.add(List.of("b", "e"));
    double[] values = new double[]{2.0, 3.0, 4.0, 5.0};
    List<List<String>> queries = new ArrayList<>();
    // queries.add(List.of("a", "c"));
    // queries.add(List.of("b", "d"));
    queries.add(List.of("e", "c"));
    // queries.add(List.of("a", "a"));
    // queries.add(List.of("a", "f"));
    double[] results = solution.calcEquation(equations, values, queries);
    System.out.println(Arrays.toString(results));
  }

  public double[] calcEquation(List<List<String>> equations, double[] values, List<List<String>> queries) {
    Graph graph = buildGraph(equations, values); 
    double[] results = new double[queries.size()];
    for (int i = 0; i < queries.size(); i++) {
      List<String> query = queries.get(i);
      double result = graph.search(query.get(0), query.get(1));
      results[i] = result;
    }
    return results;
  }
  
  private Graph buildGraph(List<List<String>> equations, double[] values) {
    Map<String, Node> nodes = new HashMap<>();
    for (int i = 0; i < equations.size(); i++) {
      List<String> equation = equations.get(i);
      String first = equation.get(0);
      String second = equation.get(1);
      double ratio = values[i];
      Node firstNode = nodes.getOrDefault(first, new Node(first));
      Node secondNode = nodes.getOrDefault(second, new Node(second));
      firstNode.putNeighbour(second, 1/ ratio);
      secondNode.putNeighbour(first, ratio);
      nodes.put(first, firstNode);
      nodes.put(second, secondNode);
    }
    return new Graph(nodes);
  }
}

class Graph {
  private Map<String, Node> nodes = new HashMap<>();

  Graph(Map<String, Node> nodes) {
    this.nodes = nodes;
  }

  double search(String start, String end) {
    Set<String> visited = new HashSet<>();
    visited.add(end);
    return search(start, end, visited);
  }

  private Double search(String start, String end, Set<String> visited) {
    if (!this.nodes.containsKey(start) || !this.nodes.containsKey(end)) {
      return -1.0;
    }
    if (start.equals(end)) {
      return 1.0;
    }
    Node endNode = this.nodes.get(end);
    for (Map.Entry<String, Double> entry: endNode.neighbours.entrySet()) {
      String next = entry.getKey();
      Double ratio = entry.getValue();
      if (!visited.contains(next)) {
        visited.add(next);
        Double result = ratio * search(start, next, visited);
        if (result >= 0.0) {
          return result;
        }
      }
    }
    return -1.0;
  }
}

class Node {
  String name;
  Map<String, Double> neighbours = new HashMap<>();

  Node(String name) {
    this.name = name;
  }

  void putNeighbour(String name, double ratio) {
    if (name.equals(this.name)) {
      return;
    }
    this.neighbours.put(name, ratio);
  }

}

