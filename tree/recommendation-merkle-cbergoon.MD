## Review of cbergoon/merkletree

### Summary

This repository contains an implementation of the Merkle tree stucture in Go with leaves of arbitrary content and type. It provides a user with the possibility to:
- Create a new tree with a list of items
- Update the whole tree with a list of items
- Get the path (left / right) to a leaf in the tree from the root
- Verify the whole tree
- Verify an item in the tree

It is also possible to use a custom hash strategy which by default is SHA256.

### Requirements Assessment

We need to determine if it makes sense to use this to store a continuous stream of pricing data, initially only with the depth of one level down the tree. Furthermore, it must make sense in the longer run, should we want to expand its usage and capacity for our market pricing mechanism.

A significant part of the logic is based on a way to pack content into a set of nodes, and insert these in a tree. Since we do not need to create a large tree and update it often, some weaknesses of this library does not matter. The overall implementation supports the creation of an arbitrary sized tree and verification of both the tree and individual content.

Tests of this library could at first glance be considered alright with 85.9\% code coverage. However, by a closer look [see attached] one could see that it misses many tests for seemingly important paths. Two fundamental test cases which I would like to see is when you build a tree but the assemble fails (e.g., no leaves to add), or when we look for a path to an item in the tree but cannot find it. Perhaps even more critical, there are no test cases which checks that we calculated a node's child hashes correctly when we verify an item's membership of a tree. This could be considered one of the main reason for as to why we use this data structure in the first place so that is a concern.

Finally, based on the activity of the repository, it seems somewhat inactive where the last commit was done 7 months ago and with no other branches. An open pull request, from Duarte four months ago, has neither been reviewed nor commented in any way. The maintainer has not committed anything to Github since almost 2 months and seems to have not been actively coding (on Github at least) in the last year. We need to take this into consideration since the repository is almost 3 years old with multiple opportunities for improvement, should we want to use it in production. 

### Recommendation

The reasons for not using this library can be summarised as follows
* Insufficient tests of critical paths
* Non-maintained and inactive project
* Not especially fitted to our needs
* The complexity of the solution is low 


My suggestion is to build a similar solution bottoms-up based on similar verification process as this and other libraries. We can keep it even more light-weight and catered to our purpose with a fixed initial capacity of four leaves (i.e. product categories). Furthermore, we can extend tests and adapt it to fit any blockchain specific scenarios needed (e.g. sorting the hashes in a certain way as Duarte requested). 
