import collections

class Solution(object):
    def leastInterval(self, tasks, N):
        task_counts = collections.Counter(tasks).values()
        M = max(task_counts)
        Mct = task_counts.count(M)
        return max(len(tasks), (M - 1) * (N + 1) + Mct)

if __name__ == '__main__':
    solution = Solution()
    tasks = ['A', 'A', 'A', 'B', 'B', 'B']
    print(solution.leastInterval(tasks, 2))
