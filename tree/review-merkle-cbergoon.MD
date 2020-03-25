### Structures

```golang
// Represents data stored and verified by the tree. Any type which 
// implements this interface can be used as an item in the tree.
type Content interface {
    CalculateHash() ([]byte, error)
	Equals(other Content) (bool, error)
}

// Represents the container of the tree. Holds a pointer to the root of 
// the tree, a list of pointer to the leaf nodes, and the merkle root.
type MerkleTree struct {
    Root         *Node
	merkleRoot   []byte
	Leaves       []*Node
	hashStrategy func() hash.Hash
}

// Represents a node, root, or leaf in the tree. Stores pointers to its immediate
// relationships, a hash, leaf content, and other metadata.
type Node struct {
    Tree   *MerkleTree
	Parent *Node
	Left   *Node
	Right  *Node
	leaf   bool
	dup    bool
	Hash   []byte
	C      content
}
```

## Functions

#### buildIntermediate(nl []*Node, mt *MerkleTree) (*Node, error)
- Iterate over the length of the slice of pointer nodes (leaf nodes) passed.
  Index the first two, they are the left and right child to a parent.
- Create each child hash, concatenate them, and add it to a node.
- Each assembled node is appended to a larger slice of pointer nodes, N.
- For the current iteration of nodes, assign the left and right child a 
  parent which is the last assembled node.
- When we only have two nodes in the slice, we return.
- Call the function again with N and mt.

#### buildWithContent(cs []Content, mt *MerkleTree) (*Node, []*Node, error)
- Calculate the hash of every content item and append it to the node slice.
- If there is an odd amount of leaves, create a duplicate one of the last to append.
- Create the root node with `buildIntermediate` and return it along with the leaves.

#### NewTree(cs []Content) (*MerkleTree, error)
- Create a Merkle tree with the default hash strategy.
- Use `buildWithContent` to get the root and leaves and assigns them to 
the fields of a merke tree struct.

#### NewTreeWithHashStragegy(cs []Content, hashStrategy func() hash.Hash) (*MerkleTree, error)
Same as `NewTree` but with a custom hash strategy


### Methods

#### (Node) verifyNode() ([]byte, error)
- If it is a leaf, return the hash
- Check both right and left nodes recursively, get both byte slices
- Get the hash strategy, default is SHA256
- Concatenate the individual hash of the left and right node, return it

#### (Node) calculateNodeHash() ([]byte, error)
- If it is a leaf, return the hash
- Get the hash strategy, default is SHA256
- Concatenate the individual hash of the left and right node, return it

#### (MerkleTree) GetMerklePath(content Content) ([][]byte, []int64, error)
- Iterate over the leaves and for each leaf check if it is equal to content.
- If that is the case, get the current parent, and check if it is equal 
- [Error?] If the left child hash matches the current leaf hash, assign append 
  1 to the index slice, where 1 should represent a right (!) leaf, and inversely
  for the right child hash. 
  Comment: the tests match this erroneous implementation but the thought way, 
  according to the comment, makes more sense from a readability perspective.
- Go up one generation and assign current parrent to the next one
- Return the path and indices 

#### (MerkleTree) RebuildTree() error
Same as to `NewTree` but uses the existing Merkle tree structure

#### (MerkleTree) RebuildTreeWith(cs []Content) error
Replaces the content in the tree with new content

#### (MerkleTree) VerifyTree() (bool, error)
- Calculate the hash of the root
- Compare the calculated hash with the existing root and compare if they are the same

#### (MerkleTree) VerifyContent(content Content) (bool, error)
- Iterate over the leaves and find the one which matches the content to verify
- Find the parent and calculate its hash based on its children
- Compare it with the parent's stored hash
- Return true if all parents' hashes check out
