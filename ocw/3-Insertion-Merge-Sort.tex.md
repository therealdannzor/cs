## 3. Sort: Insertion and Merge

There are multiple reasons for as why we would like sorted list. Some obvious ones are a phone book with names and numbers corresponding to a directory, spreadsheets, and mp3 files. An example of a problem which becomes easier once items are sorted is finding a median of an array. Also, with a sorted list, we can use binary search to find an item $k$ in a list much quicker than in an unordered list. We can turn a linear search into a logarithmic search. 

### Insertion Sort

 For $i$ = 1, 2, ..., $n$ we want to insert  A[$i$] into a sorted array A[0, .., $n$-1] by pairwise swaps down to the correct position. We start with pointing our key variable to index 1 - the second element - because the very first element is sorted by definition.  Comparison is then done between the iterator at A[0] and key variable at A[1]. 

Analysis: we can consider it $\theta(n)$ in terms of key positions, even if we only move it $n$-1 steps. Each step is $\theta(n)$ comparison and swaps.

It is assumed that a compare and swap are equal in cost and this is the case for integers. For swapping abstract record types, they are the same as for integers in the sense that we need to change pointer locations. However, for the same type of comparison the function could be much more complex and thus much more expensive than swaps.

_Question_: compares are more expensive than swaps so I care more about compare operations than swaps. How can I modify the insertion algorithm to be more efficient? Assume comparison is $\theta(n^2)$

_Answer_: instead of the pairwise comparison, we do a binary search on A already sorted in $\theta(\log i )$. For each of those steps, we get a total cost of $\theta(n \log n)$ for compares.

_Question_: does this new method help swaps for an array data structure?

_Answer_: no, because binary search requires insertion in whole A. We need to insert A[i] into the right position and worst-case scenario we need to shift a lot of things to the right which takes us back to $\theta(n)$.

**To summarise**: binary search in insertion sort gives you $\theta(n \log n)$ for compares but still $\theta(n^2)$ for swaps.

### Merge Sort

Based on divide and conquer, and a standard recursive algorithm. Given an array A, we split the array until the end where we do all the work to assemble them back to the original size.

For the function, the two inputs are assumed to be sorted arrays. For an array of length $n$, we start by splitting it up in two sub-arrays $L$ and $R$ denoting the left half and right half, respectively.

_Question_: the cost for merge sort is $\theta(n \log n)$ but how can we convince ourselves this is the case?

_Answer_: assume we model the complexity the following:
```
T(n) = c_1 + 2T(n/2) + c*n 
```
where `C_1` is the cost to divide the array into sub-arrays, `T(n/2)` is the recursive part, and `c*n` is the merge part. To model the recurrence, we ignore `c_1` since `c*n` dominates it, so it becomes

```
            c*n
           /   \
      c*n/2     c*n/2
     /     \   /     \
 c*n/4 c*n/4 c*n/4 c*n/4
```

_Question_: how can we calculate the levels and leaves?

_Answer_: as seen, the levels are 1 + lg $n$ and the leaves are $n$. 

By adding up the work in each of the levels of the tree, we see that they all have `c*n` amount of work. So it is doing roughly the same amount of work in each level. Were we to write an equation to T(n) we would get: 

(1 + lg n) $\times$ cn = $\theta(n \log n)$

#### Advantages: Insertion vs Merge Sort

Insertion uses in-place sorting and does not require additional space as merge needs to allocate. 

Merge sort $\Rightarrow$ $\theta(n)$ auxiliary space

Insert sort $\Rightarrow$ $\theta(1)$ auxiliary space 
