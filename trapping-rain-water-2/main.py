from typing import List, Tuple
import heapq


class Node:
  def __init__(self, priority: int, item: Tuple[int, int]):
    self.priority = priority
    self.item = item

  def __lt__(self, other) -> bool:
      return self.priority < other.priority


class PriorityQueue:
  def __init__(self, iterable: List[Node]):
    heapq.heapify(iterable)
    self.iterable = iterable

  def __len__(self):
    return len(self.iterable)

  def pop(self) -> Node:
    if len(self.iterable) > 0:
      return heapq.heappop(self.iterable)
    else:
      raise IndexError("index out of range in heap")

  def push(self, node: Node):
    heapq.heappush(self.iterable, node)


class Solution:
  nx: int = 0
  ny: int = 0
  visited: List[List[bool]]
  boundaries: PriorityQueue

  def trapRainWater(self, heightMap: List[List[int]]) -> int:
    if len(heightMap) == 0 or len(heightMap[0]) == 0:
      return 0
    self.init(heightMap)
    mx = -1
    neighbors = ((0, 1), (1, 0), (0, -1), (-1, 0))
    water = 0
    while len(self.boundaries) > 0:
      node = self.boundaries.pop()
      mx = max(mx, node.priority)
      for neighbor in neighbors:
        x = node.item[0] + neighbor[0]
        y = node.item[1] + neighbor[1]
        if x < 0 or x >= self.nx or y < 0 or y >= self.ny:
          continue
        if self.visited[x][y]:
          continue
        if mx > heightMap[x][y]:
          water += mx - heightMap[x][y]
        next = Node(heightMap[x][y], (x, y))
        self.boundaries.push(next)
        self.visited[x][y] = True
    return water

  def init(self, heightMap: List[List[int]]):
    self.nx = len(heightMap)
    self.ny = len(heightMap[0])
    self.visited = [[False for _ in range(self.ny)] for _ in range(self.nx)]
    boundaries = []
    for i in range(self.nx):
      if not self.visited[i][0]:
        self.visited[i][0] = True
        boundaries.append(Node(heightMap[i][0], (i, 0)))
      if not self.visited[i][self.ny-1]:
        self.visited[i][self.ny-1] = True
        boundaries.append(Node(heightMap[i][self.ny-1], (i, self.ny-1)))
    for j in range(self.ny):
      if not self.visited[0][j]:
        self.visited[0][j] = True
        boundaries.append(Node(heightMap[0][j], (0, j)))
      if not self.visited[self.nx-1][j]:
        self.visited[self.nx-1][j] = True
        boundaries.append(Node(heightMap[self.nx-1][j], (self.nx-1, j)))
    self.boundaries = PriorityQueue(boundaries)
    self.heightMap = heightMap

if __name__ == "__main__":
  heightMap = [
      [1, 4, 3, 1, 3, 2],
      [3, 2, 1, 3, 2, 4],
      [2, 3, 3, 2, 3, 1]
  ]
  solution = Solution()
  result = solution.trapRainWater(heightMap)
  print(result)
