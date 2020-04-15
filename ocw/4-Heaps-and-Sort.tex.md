## 3. Heaps and Heap Sort

### Priority Queue
Priority queue: implements a set of elements where each element is associated with a key. 
The set of operations needed:
1. Insert(S,x): insert element x into set S
2. max(S): return largest key in S
3. extract_max(S): return and remove the largest key in S
4. update_key(S,x,k): update key in S to a new value k

### Heap
A heap is an implementation of a priority queue. An array visualised as a nearly complete binary tree. The tree representation can be created from an array. The rules are:
* The root of the tree is the first element in the array, $i$ = 1
* The parent of a node $i$ is $i$/2
* The left child of $i$ is 2$i$
* The right child of $i$ is 2$i$ + 1

### Max Heap Property
The key of a node is $\ge$ to the keys of its children. This must be true for every node of the tree. The height of the visualised tree is bounded by log $n$ and it is a nearly perfect binary tree. The set of operations needed:
1. build_max_heap: produce a max heap from an unordered array (changing its order)
2. max_heapify: correct a single violation of the heap property in a subtree's root '

A key assumption for max_heapify is that the trees rooted at left($i$) and right($i$) are max heaps. The complexity is $O$(log$n$). 

The pseudocode to build a max heap from an unordered array $A$ is to:
1. for $i$ = $n$/2 down to 1
2. do max_heapify(A,i)

Why can we start with $n$/2? An observation: elements $A$[$n$/2+1,...,$n$] are all leaves and they satisfy the max heap property. In the first iteration, we start with max heap of two leaves which leaves us a trivial example. In the next one, we will proceed with the sibling, and so on. At every step we have satisfied the key assumption. 

A simple analysis of this is $O$($n$log$n$). Expanding the work of each step results in an arithmetic series which reveals that to **build a max heap is in fact** $O$($n$). 
