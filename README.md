# layup-sequence
Solution.

Includes solution coded in golang, as well as a graph with the plot points obtained from running main.go

Solution was a straight forward recursive function with memoization to avoid re-calculating values of n resulting in O(n) runtime. Minor optimization recurses n-1 and n-2 steps separately to avoid visiting each n value twice.