## 2. Models of Computation

### What Is An Algorithm

A computational procedure for solving a problem: input -> algorithm -> output

A model of computation specifies:
- allowed operations of an algorithm
- cost (time, memory) of each operation


### Comparison

|      Program       |        Algorithm       |
|--------------------|:----------------------:|
| Prog language      | Pseudo code            |
| Computer           | Model of computation   |
| col 3 is           | right-aligned          | 


### Random Access Machine
The memory side of a computer, based on a Random Access Machine (RAM), can be modelled as an array. In constant time $\theta(1)$, we
- load $\theta(1)$ words
- do $\theta(1)$ computations
- store $\theta(1)$ words
- all based on $\theta(1)$ registers

A word in this class is assumed to be a constant amount of bits. 

### Pointer Machine

In contrast to RAM, which has no dynamic memory allocation by default, this is taken care of in this model. In a nutshell, this model corresponds to object-oriented programming. We have
- dynamically located objects
- object has $\theta(1)$ fields

A field is equivalent to a word (e.g. int) or a pointer to an object or null (nil, None).

#### Linked list

Each node has a value and a reference to the next node. It may or may not have to the previous one, the head or the tail etc. This is a structure in the pointer machine.

### Python Model

1) list = array
    `L[i] = L[j] + 5` is constant time
2) object with constant time attribute
    `x = x.next`  is constant time
3) `L.append(x)` is constant time and uses table-doubling. 
4) `L = L1 + L2` 

### Document Distance Problem

Operation: $d$($D_1$, $D_2$)

- document: sequence of words
- word: string of alphanumeric characters
- idea: shared words 
- think of document as a vector
- D[w] = number of occurrences of w in D

The algorithm to find occurrences of words in common between two documents:

1. split document into words
2. compute word frequency
3. dot product
