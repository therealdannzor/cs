## 3. Sort: Insertion and Merge

There are multiple reasons for as why we would like sorted list. Some obvious ones are a phone book with names and numbers corresponding to a directory, spreadsheets, and mp3 files. An example of a problem which becomes easier once items are sorted is finding a median of an array. Also, with a sorted list, we can use binary search to find an item <img src="/ocw/tex/63bb9849783d01d91403bc9a5fea12a2.svg?invert_in_darkmode&sanitize=true" align=middle width=9.075367949999992pt height=22.831056599999986pt/> in a list much quicker than in an unordered list. We can turn a linear search into a logarithmic search. 

### Insertion Sort

 For <img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/> = 1, 2, ..., <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/> we want to insert  A[<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>] into a sorted array A[0, .., <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>-1] by pairwise swaps down to the correct position. We start with pointing our key variable to index 1 - the second element - because the very first element is sorted by definition.  Comparison is then done between the iterator at A[0] and key variable at A[1]. 

Analysis: we can consider it <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/> in terms of key positions, even if we only move it <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>-1 steps. Each step is <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/> comparison and swaps.

It is assumed that a compare and swap are equal in cost and this is the case for integers. For swapping abstract record types, they are the same as for integers in the sense that we need to change pointer locations. However, for the same type of comparison the function could be much more complex and thus much more expensive than swaps.

_Question_: compares are more expensive than swaps so I care more about compare operations than swaps. How can I modify the insertion algorithm to be more efficient? Assume comparison is <img src="/ocw/tex/ab861451a1dc2c93988ff45f8f66d0ca.svg?invert_in_darkmode&sanitize=true" align=middle width=38.200296749999985pt height=26.76175259999998pt/>

_Answer_: instead of the pairwise comparison, we do a binary search on A already sorted in <img src="/ocw/tex/6a71e13dd6cfe62542618222d924919d.svg?invert_in_darkmode&sanitize=true" align=middle width=50.59475684999998pt height=24.65753399999998pt/>. For each of those steps, we get a total cost of <img src="/ocw/tex/d8c161ead983102da0d32a547419bba6.svg?invert_in_darkmode&sanitize=true" align=middle width=67.40494199999999pt height=24.65753399999998pt/> for compares.

_Question_: does this new method help swaps for an array data structure?

_Answer_: no, because binary search requires insertion in whole A. We need to insert A[i] into the right position and worst-case scenario we need to shift a lot of things to the right which takes us back to <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/>.

**To summarise**: binary search in insertion sort gives you <img src="/ocw/tex/d8c161ead983102da0d32a547419bba6.svg?invert_in_darkmode&sanitize=true" align=middle width=67.40494199999999pt height=24.65753399999998pt/> for compares but still <img src="/ocw/tex/ab861451a1dc2c93988ff45f8f66d0ca.svg?invert_in_darkmode&sanitize=true" align=middle width=38.200296749999985pt height=26.76175259999998pt/> for swaps.

### Merge Sort

Based on divide and conquer, and a standard recursive algorithm. Given an array A, we split the array until the end where we do all the work to assemble them back to the original size.

For the function, the two inputs are assumed to be sorted arrays. For an array of length <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>, we start by splitting it up in two sub-arrays <img src="/ocw/tex/ddcb483302ed36a59286424aa5e0be17.svg?invert_in_darkmode&sanitize=true" align=middle width=11.18724254999999pt height=22.465723500000017pt/> and <img src="/ocw/tex/1e438235ef9ec72fc51ac5025516017c.svg?invert_in_darkmode&sanitize=true" align=middle width=12.60847334999999pt height=22.465723500000017pt/> denoting the left half and right half, respectively.

_Question_: the cost for merge sort is <img src="/ocw/tex/d8c161ead983102da0d32a547419bba6.svg?invert_in_darkmode&sanitize=true" align=middle width=67.40494199999999pt height=24.65753399999998pt/> but how can we convince ourselves this is the case?

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

_Answer_: as seen, the levels are 1 + lg <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/> and the leaves are <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>. 

By adding up the work in each of the levels of the tree, we see that they all have `c*n` amount of work. So it is doing roughly the same amount of work in each level. Were we to write an equation to T(n) we would get: 

(1 + lg n) <img src="/ocw/tex/bdbf342b57819773421273d508dba586.svg?invert_in_darkmode&sanitize=true" align=middle width=12.785434199999989pt height=19.1781018pt/> cn = <img src="/ocw/tex/d8c161ead983102da0d32a547419bba6.svg?invert_in_darkmode&sanitize=true" align=middle width=67.40494199999999pt height=24.65753399999998pt/>

#### Advantages: Insertion vs Merge Sort

Insertion uses in-place sorting and does not require additional space as merge needs to allocate. 

Merge sort <img src="/ocw/tex/777d001ea1ec5971b67bb546ed760f97.svg?invert_in_darkmode&sanitize=true" align=middle width=16.43840384999999pt height=14.15524440000002pt/> <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/> auxiliary space

Insert sort <img src="/ocw/tex/777d001ea1ec5971b67bb546ed760f97.svg?invert_in_darkmode&sanitize=true" align=middle width=16.43840384999999pt height=14.15524440000002pt/> <img src="/ocw/tex/0636c9bb43c30e08e2ae96add442d27c.svg?invert_in_darkmode&sanitize=true" align=middle width=29.17816934999999pt height=24.65753399999998pt/> auxiliary space 
