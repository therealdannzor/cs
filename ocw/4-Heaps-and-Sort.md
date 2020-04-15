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
* The root of the tree is the first element in the array, <img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/> = 1
* The parent of a node <img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/> is <img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>/2
* The left child of <img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/> is 2<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>
* The right child of <img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/> is 2<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/> + 1

### Max Heap Property
The key of a node is <img src="/ocw/tex/22626c24dc8a958fe7cb1512303a8b77.svg?invert_in_darkmode&sanitize=true" align=middle width=12.785434199999989pt height=20.908638300000003pt/> to the keys of its children. This must be true for every node of the tree. The height of the visualised tree is bounded by log <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/> and it is a nearly perfect binary tree. The set of operations needed:
1. build_max_heap: produce a max heap from an unordered array (changing its order)
2. max_heapify: correct a single violation of the heap property in a subtree's root '

A key assumption for max_heapify is that the trees rooted at left(<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>) and right(<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>) are max heaps. The complexity is <img src="/ocw/tex/9afe6a256a9817c76b579e6f5db9a578.svg?invert_in_darkmode&sanitize=true" align=middle width=12.99542474999999pt height=22.465723500000017pt/>(log<img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>). 

The pseudocode to build a max heap from an unordered array <img src="/ocw/tex/53d147e7f3fe6e47ee05b88b166bd3f6.svg?invert_in_darkmode&sanitize=true" align=middle width=12.32879834999999pt height=22.465723500000017pt/> is to:
1. for <img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/> = <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>/2 down to 1
2. do max_heapify(A,i)

Why can we start with <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>/2? An observation: elements <img src="/ocw/tex/53d147e7f3fe6e47ee05b88b166bd3f6.svg?invert_in_darkmode&sanitize=true" align=middle width=12.32879834999999pt height=22.465723500000017pt/>[<img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>/2+1,...,<img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>] are all leaves and they satisfy the max heap property. In the first iteration, we start with max heap of two leaves which leaves us a trivial example. In the next one, we will proceed with the sibling, and so on. At every step we have satisfied the key assumption. 

A simple analysis of this is <img src="/ocw/tex/9afe6a256a9817c76b579e6f5db9a578.svg?invert_in_darkmode&sanitize=true" align=middle width=12.99542474999999pt height=22.465723500000017pt/>(<img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>log<img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>). Expanding the work of each step results in an arithmetic series which reveals that to **build a max heap is in fact** <img src="/ocw/tex/9afe6a256a9817c76b579e6f5db9a578.svg?invert_in_darkmode&sanitize=true" align=middle width=12.99542474999999pt height=22.465723500000017pt/>(<img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>). 
