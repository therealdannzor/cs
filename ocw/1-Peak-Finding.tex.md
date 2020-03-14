## 1. Algorithmic Thinking, Peak Finding

### One-dimensional case

Given: a one-dimensional vector: [$a$, $b$, $c$, $d$, $e$, $f$, $g$, $h$, $i$], where $a$ to $i$ are numbers

<ins>Examples:</ins>
Index 0 is a peak iff a $\geqslant$ b
Index 1 is a peak iff b $\geqslant$ a and b $\geqslant$ c
Index 8 is a peak iff i $\geqslant$ h

Problem: find a peak, if it exists

#### Straightforward algorithm

Start from the left in 
```
[1, 2, .., .., n/2, .., n-1, n]
```
Look at n/2 elements, there might be a peak in the middle.

Worst case complexity: $\theta(n)$

#### Recursive algorithm (Divide & Conquer)

Start from the middle (n/2) position
```
[1, 2, .., n/2-1, n/2, n/2+1, .., .., n-1, n]
```
If a[n/2] < a[n/2-1] then only look at left half, 1 to n/2 - 1, to look for a peak. 

Else if a[n/2] < a[n/2+1] then do the same for the right half.

Let T(n) be the work of the algorithm on an input of size $n$. We can then see that

T(n) = T(n/2) + $\theta(1)$

Whereas the base case, an array with only one element, is

T(1) = $\theta(1)$

This can be generalised as 

T(n) = $\theta(1)$ + ... + $\theta(1)$ = {$log_2$ times} = $\theta($$log_2$$n)$

### Two-dimensional case

Given matrix B with $n$ rows and $m$ columns

$$
\begin{bmatrix}
. & c & . & .\\
b & a & d & .\\
. & e & . & .\\
. & . & . & . \\
\end{bmatrix}
$$

$a$ is a 2D-peak iff a $\geqslant$ b, c, d, e

Example with the Greedy Ascent algorithm: pick a starting point, e.g. in the middle, and the default traversal directions are.
$$
\begin{bmatrix}
. & . & . & .\\
14 & 13 & 12 & .\\
15 & 9 & 11 & 17\\  
16 & 17 & 19 & 20 \\ 
\end{bmatrix}
$$

We start at 12. Look to your left and right, and follow the greatest direction. In this case the path is 12, 13, 14, 15, 16, 17, 19, 20.

Worst case complexity: $\theta(mn)$ = {if $m$ = $n$ } = $\theta(n^2)$

Second attempt: 
* Pick the middle column  $j$ = $m/2$
* Find a global max on column $j$  at ($i$, $j$)
* Compare with ($i$, $j-1$), ($i$, $j$), ($i$, $j+1$).
* Pick left column if ($i$, $j-1$) > ($i$, $j$), and similarly for right 
* If ($i$, $j$) $\geqslant$ ($i$,$j-1$), ($i$,$j+1$) &Rightarrow; ($i$, $j$) is a 2D peak
* Solve the new problem with the half number of columns
* You are done when you have a single column (find that maximum)

This new approach matches the 1D case more closely.

The complexity is: $T(n, m)$ = $T(n, m/2)$ + $\theta(n)$ &larr; max

First we break the work down into half the number of columns, $m/2$. In order to find the global maximum we do $\theta(n)$ work.

Finally, we find that 

$T(n, 1)$ = $\theta(n)$
$T(n, m)$ = $\theta(n)$  + ... + $\theta(n)$  = $\theta(n$$log_2$$m)$

## Reflection

_Summary_: This lecture discussed the methodology on how to find a working algorithm and analyse its efficiency. It showed both how to approach it in a brute-force way and how to use a more efficient approach, namely a divide-and-conquer algorithm. 

_Principles_: When developing a solution, make sure to think holistically (especially in a multi dimensional case) to get the solution correct rather than focusing only on optimisation.

_Application_: If given a fixed size array of size $n$, if the problem statement allows it, try to solve sub-parts of the problem while adhering to the constraints. 
