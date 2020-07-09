import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.TreeSet;
import java.util.HashSet;

class Solution {

  public static void main(String[] args) {
    List<List<String>> accounts = List.of(
        List.of("John", "johnsmith@mail.com", "john00@mail.com"),
        List.of("John", "johnnybravo@mail.com"),
        List.of("John", "johnsmith@mail.com", "john_newyork@mail.com"),
        List.of("John", "johnnybravo@mail.com", "john_newyork@mail.com"),
        List.of("Mary", "mary@mail.com"));
    Solution solution = new Solution();
    List<List<String>> result = solution.accountsMerge(accounts);
    System.out.println(result);
  }

  public List<List<String>> accountsMerge(List<List<String>> accounts) {
    List<int[]> unions = makeUnions(accounts);
    int[] set = makeSet(accounts.size());
    buildRelation(set, unions);
    return makeResult(set, accounts);
  }

  private List<List<String>> makeResult(int[] set, List<List<String>> accounts) {
    Map<Integer, List<Integer>> partitions = new HashMap<>();
    for (int i = 0; i < set.length; i++) {
      int root = find(set, i); 
      List<Integer> partition = partitions.getOrDefault(root, new LinkedList<>());
      partition.add(i);
      partitions.put(root, partition);
    }
    Map<Integer, Set<String>> container = new HashMap<>();
    for (Map.Entry<Integer, List<Integer>> entry: partitions.entrySet()) {
      Set<String> emails = new HashSet<>();
      container.put(entry.getKey(), emails);
      for (Integer index: entry.getValue()) {
        List<String> account = accounts.get(index);
        for (int i = 1; i < account.size(); i++) {
          emails.add(account.get(i)); 
        }
      }
    }
    List<List<String>> results = new LinkedList<>();
    for (Map.Entry<Integer, Set<String>> entry: container.entrySet()) {
      List<String> account = new LinkedList<>();
      account.add(accounts.get(entry.getKey()).get(0));
      List<String> emails = new LinkedList<>(entry.getValue());
      Collections.sort(emails);
      account.addAll(emails);
      results.add(account);
    }
    return results;
  }

  private void buildRelation(int[] set, List<int[]> unions) {
    for (int[] union: unions) {
      int i = find(set, union[0]);
      int j = find(set, union[1]);
      set[j] = i;
    }
  }

  private Integer find(int[] set, int index) {
    while (set[index] != index) {
      index = set[index]; 
    }
    return index;
  }

  private int[] makeSet(int length) {
    int[] set = new int[length];
    for (int i = 0; i < length; i++) {
      set[i] = i; 
    }
    return set;
  }

  private List<int[]> makeUnions(List<List<String>> accounts) {
    List<int[]> unions = new LinkedList<>();
    Map<String, Integer> container = new HashMap<>();
    for (int i = 0; i < accounts.size(); i++) {
      List<String> account = accounts.get(i);
      for (int j = 1; j < account.size(); j++) {
        String email = account.get(j);           
        Integer anotherAccount = container.get(email);
        if (anotherAccount != null) {
          unions.add(new int[]{anotherAccount, i});
        } else {
          container.put(email, i);
        }
      } 
    }
    return unions;
  }
}

