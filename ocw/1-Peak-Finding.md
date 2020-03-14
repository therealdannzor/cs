## 1. Algorithmic Thinking, Peak Finding

### One-dimensional case

Given: a one-dimensional vector: [<img src="/ocw/tex/44bc9d542a92714cac84e01cbbb7fd61.svg?invert_in_darkmode&sanitize=true" align=middle width=8.68915409999999pt height=14.15524440000002pt/>, <img src="/ocw/tex/4bdc8d9bcfb35e1c9bfb51fc69687dfc.svg?invert_in_darkmode&sanitize=true" align=middle width=7.054796099999991pt height=22.831056599999986pt/>, <img src="/ocw/tex/3e18a4a28fdee1744e5e3f79d13b9ff6.svg?invert_in_darkmode&sanitize=true" align=middle width=7.11380504999999pt height=14.15524440000002pt/>, <img src="/ocw/tex/2103f85b8b1477f430fc407cad462224.svg?invert_in_darkmode&sanitize=true" align=middle width=8.55596444999999pt height=22.831056599999986pt/>, <img src="/ocw/tex/8cd34385ed61aca950a6b06d09fb50ac.svg?invert_in_darkmode&sanitize=true" align=middle width=7.654137149999991pt height=14.15524440000002pt/>, <img src="/ocw/tex/190083ef7a1625fbc75f243cffb9c96d.svg?invert_in_darkmode&sanitize=true" align=middle width=9.81741584999999pt height=22.831056599999986pt/>, <img src="/ocw/tex/3cf4fbd05970446973fc3d9fa3fe3c41.svg?invert_in_darkmode&sanitize=true" align=middle width=8.430376349999989pt height=14.15524440000002pt/>, <img src="/ocw/tex/2ad9d098b937e46f9f58968551adac57.svg?invert_in_darkmode&sanitize=true" align=middle width=9.47111549999999pt height=22.831056599999986pt/>, <img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>], where <img src="/ocw/tex/44bc9d542a92714cac84e01cbbb7fd61.svg?invert_in_darkmode&sanitize=true" align=middle width=8.68915409999999pt height=14.15524440000002pt/> to <img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/> are numbers

<ins>Examples:</ins>
Index 0 is a peak iff a <img src="/ocw/tex/193b8c60b5d3de2f2660077fb7807497.svg?invert_in_darkmode&sanitize=true" align=middle width=12.785434199999989pt height=20.931464400000007pt/> b
Index 1 is a peak iff b <img src="/ocw/tex/193b8c60b5d3de2f2660077fb7807497.svg?invert_in_darkmode&sanitize=true" align=middle width=12.785434199999989pt height=20.931464400000007pt/> a and b <img src="/ocw/tex/193b8c60b5d3de2f2660077fb7807497.svg?invert_in_darkmode&sanitize=true" align=middle width=12.785434199999989pt height=20.931464400000007pt/> c
Index 8 is a peak iff i <img src="/ocw/tex/193b8c60b5d3de2f2660077fb7807497.svg?invert_in_darkmode&sanitize=true" align=middle width=12.785434199999989pt height=20.931464400000007pt/> h

Problem: find a peak, if it exists

#### Straightforward algorithm

Start from the left in 
```
[1, 2, .., .., n/2, .., n-1, n]
```
Look at n/2 elements, there might be a peak in the middle.

Worst case complexity: <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/>

#### Recursive algorithm (Divide & Conquer)

Start from the middle (n/2) position
```
[1, 2, .., n/2-1, n/2, n/2+1, .., .., n-1, n]
```
If a[n/2] < a[n/2-1] then only look at left half, 1 to n/2 - 1, to look for a peak. 

Else if a[n/2] < a[n/2+1] then do the same for the right half.

Let T(n) be the work of the algorithm on an input of size <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>. We can then see that

T(n) = T(n/2) + <img src="/ocw/tex/0636c9bb43c30e08e2ae96add442d27c.svg?invert_in_darkmode&sanitize=true" align=middle width=29.17816934999999pt height=24.65753399999998pt/>

Whereas the base case, an array with only one element, is

T(1) = <img src="/ocw/tex/0636c9bb43c30e08e2ae96add442d27c.svg?invert_in_darkmode&sanitize=true" align=middle width=29.17816934999999pt height=24.65753399999998pt/>

This can be generalised as 

T(n) = <img src="/ocw/tex/0636c9bb43c30e08e2ae96add442d27c.svg?invert_in_darkmode&sanitize=true" align=middle width=29.17816934999999pt height=24.65753399999998pt/> + ... + <img src="/ocw/tex/0636c9bb43c30e08e2ae96add442d27c.svg?invert_in_darkmode&sanitize=true" align=middle width=29.17816934999999pt height=24.65753399999998pt/> = {<img src="/ocw/tex/e5b11324452d2e5806450dc47cbfe76f.svg?invert_in_darkmode&sanitize=true" align=middle width=27.589524599999987pt height=22.831056599999986pt/> times} = <img src="/ocw/tex/b5398a0e54d0b0389508811e3f58246c.svg?invert_in_darkmode&sanitize=true" align=middle width=14.566244549999992pt height=24.65753399999998pt/><img src="/ocw/tex/e5b11324452d2e5806450dc47cbfe76f.svg?invert_in_darkmode&sanitize=true" align=middle width=27.589524599999987pt height=22.831056599999986pt/><img src="/ocw/tex/d5832d8437237f9b3dbe873f044d5de9.svg?invert_in_darkmode&sanitize=true" align=middle width=16.25959334999999pt height=24.65753399999998pt/>

### Two-dimensional case

Given matrix B with <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/> rows and <img src="/ocw/tex/0e51a2dede42189d77627c4d742822c3.svg?invert_in_darkmode&sanitize=true" align=middle width=14.433101099999991pt height=14.15524440000002pt/> columns

<p align="center"><img src="/ocw/tex/8168486d0e11a8fb48d7517f0a122c43.svg?invert_in_darkmode&sanitize=true" align=middle width=100.0990221pt height=78.9048876pt/></p>

<img src="/ocw/tex/44bc9d542a92714cac84e01cbbb7fd61.svg?invert_in_darkmode&sanitize=true" align=middle width=8.68915409999999pt height=14.15524440000002pt/> is a 2D-peak iff a <img src="/ocw/tex/193b8c60b5d3de2f2660077fb7807497.svg?invert_in_darkmode&sanitize=true" align=middle width=12.785434199999989pt height=20.931464400000007pt/> b, c, d, e

Example with the Greedy Ascent algorithm: pick a starting point, e.g. in the middle, and the default traversal directions are.
<p align="center"><img src="/ocw/tex/16c77624a4f2c3d4878545565e156036.svg?invert_in_darkmode&sanitize=true" align=middle width=136.98657225pt height=78.9048876pt/></p>

We start at 12. Look to your left and right, and follow the greatest direction. In this case the path is 12, 13, 14, 15, 16, 17, 19, 20.

Worst case complexity: <img src="/ocw/tex/97baba2c8fdeb1248ae6f3abc77f9257.svg?invert_in_darkmode&sanitize=true" align=middle width=45.25893734999999pt height=24.65753399999998pt/> = {if <img src="/ocw/tex/0e51a2dede42189d77627c4d742822c3.svg?invert_in_darkmode&sanitize=true" align=middle width=14.433101099999991pt height=14.15524440000002pt/> = <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/> } = <img src="/ocw/tex/ab861451a1dc2c93988ff45f8f66d0ca.svg?invert_in_darkmode&sanitize=true" align=middle width=38.200296749999985pt height=26.76175259999998pt/>

Second attempt: 
* Pick the middle column  <img src="/ocw/tex/36b5afebdba34564d884d347484ac0c7.svg?invert_in_darkmode&sanitize=true" align=middle width=7.710416999999989pt height=21.68300969999999pt/> = <img src="/ocw/tex/b6d34a0531e535d85056999ee1ced8a2.svg?invert_in_darkmode&sanitize=true" align=middle width=30.87151979999999pt height=24.65753399999998pt/>
* Find a global max on column <img src="/ocw/tex/36b5afebdba34564d884d347484ac0c7.svg?invert_in_darkmode&sanitize=true" align=middle width=7.710416999999989pt height=21.68300969999999pt/>  at (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>, <img src="/ocw/tex/36b5afebdba34564d884d347484ac0c7.svg?invert_in_darkmode&sanitize=true" align=middle width=7.710416999999989pt height=21.68300969999999pt/>)
* Compare with (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>, <img src="/ocw/tex/fda4e6a332eec2a3985445b0c195f6f6.svg?invert_in_darkmode&sanitize=true" align=middle width=36.02081834999999pt height=21.68300969999999pt/>), (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>, <img src="/ocw/tex/36b5afebdba34564d884d347484ac0c7.svg?invert_in_darkmode&sanitize=true" align=middle width=7.710416999999989pt height=21.68300969999999pt/>), (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>, <img src="/ocw/tex/cff3823534d4d7b1833d65d60121442f.svg?invert_in_darkmode&sanitize=true" align=middle width=36.02081834999999pt height=21.68300969999999pt/>).
* Pick left column if (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>, <img src="/ocw/tex/fda4e6a332eec2a3985445b0c195f6f6.svg?invert_in_darkmode&sanitize=true" align=middle width=36.02081834999999pt height=21.68300969999999pt/>) > (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>, <img src="/ocw/tex/36b5afebdba34564d884d347484ac0c7.svg?invert_in_darkmode&sanitize=true" align=middle width=7.710416999999989pt height=21.68300969999999pt/>), and similarly for right 
* If (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>, <img src="/ocw/tex/36b5afebdba34564d884d347484ac0c7.svg?invert_in_darkmode&sanitize=true" align=middle width=7.710416999999989pt height=21.68300969999999pt/>) <img src="/ocw/tex/193b8c60b5d3de2f2660077fb7807497.svg?invert_in_darkmode&sanitize=true" align=middle width=12.785434199999989pt height=20.931464400000007pt/> (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>,<img src="/ocw/tex/fda4e6a332eec2a3985445b0c195f6f6.svg?invert_in_darkmode&sanitize=true" align=middle width=36.02081834999999pt height=21.68300969999999pt/>), (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>,<img src="/ocw/tex/cff3823534d4d7b1833d65d60121442f.svg?invert_in_darkmode&sanitize=true" align=middle width=36.02081834999999pt height=21.68300969999999pt/>) &Rightarrow; (<img src="/ocw/tex/77a3b857d53fb44e33b53e4c8b68351a.svg?invert_in_darkmode&sanitize=true" align=middle width=5.663225699999989pt height=21.68300969999999pt/>, <img src="/ocw/tex/36b5afebdba34564d884d347484ac0c7.svg?invert_in_darkmode&sanitize=true" align=middle width=7.710416999999989pt height=21.68300969999999pt/>) is a 2D peak
* Solve the new problem with the half number of columns
* You are done when you have a single column (find that maximum)

This new approach matches the 1D case more closely.

The complexity is: <img src="/ocw/tex/b521a1c83eec9d0b28b2a3ab8684809c.svg?invert_in_darkmode&sanitize=true" align=middle width=56.28060404999999pt height=24.65753399999998pt/> = <img src="/ocw/tex/41b4682b15a3eafb303044a3058bff72.svg?invert_in_darkmode&sanitize=true" align=middle width=72.71902275pt height=24.65753399999998pt/> + <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/> &larr; max

First we break the work down into half the number of columns, <img src="/ocw/tex/b6d34a0531e535d85056999ee1ced8a2.svg?invert_in_darkmode&sanitize=true" align=middle width=30.87151979999999pt height=24.65753399999998pt/>. In order to find the global maximum we do <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/> work.

Finally, we find that 

<img src="/ocw/tex/dad813e80244a2ed0f7b7f169329444b.svg?invert_in_darkmode&sanitize=true" align=middle width=50.066713949999986pt height=24.65753399999998pt/> = <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/>
<img src="/ocw/tex/b521a1c83eec9d0b28b2a3ab8684809c.svg?invert_in_darkmode&sanitize=true" align=middle width=56.28060404999999pt height=24.65753399999998pt/> = <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/>  + ... + <img src="/ocw/tex/7da70c489e27c3ebd518e94f26a8185b.svg?invert_in_darkmode&sanitize=true" align=middle width=30.825837899999993pt height=24.65753399999998pt/>  = <img src="/ocw/tex/b29cd6abd4e0f86743056ebbee536918.svg?invert_in_darkmode&sanitize=true" align=middle width=24.433120799999987pt height=24.65753399999998pt/><img src="/ocw/tex/e5b11324452d2e5806450dc47cbfe76f.svg?invert_in_darkmode&sanitize=true" align=middle width=27.589524599999987pt height=22.831056599999986pt/><img src="/ocw/tex/fc4673be68822a98eaf6bc1417f533bf.svg?invert_in_darkmode&sanitize=true" align=middle width=20.82581654999999pt height=24.65753399999998pt/>

## Reflection

_Summary_: This lecture discussed the methodology on how to find a working algorithm and analyse its efficiency. It showed both how to approach it in a brute-force way and how to use a more efficient approach, namely a divide-and-conquer algorithm. 

_Principles_: When developing a solution, make sure to think holistically (especially in a multi dimensional case) to get the solution correct rather than focusing only on optimisation.

_Application_: If given a fixed size array of size <img src="/ocw/tex/55a049b8f161ae7cfeb0197d75aff967.svg?invert_in_darkmode&sanitize=true" align=middle width=9.86687624999999pt height=14.15524440000002pt/>, if the problem statement allows it, try to solve sub-parts of the problem while adhering to the constraints. 
