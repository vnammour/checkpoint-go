/* Fri Feb  7 05:47:25 PM IST 2025 */
/* By: Victor Nammour */
/*
The keys in a binary search tree are always stored in such a way as to satisfy the binary-search-tree property:
Let x be a node in a binary search tree. If y is a node in the left subtree of x, then key[y] <= key[x].
If y is a node in the right subtree of x, then key[x] <= key[y].
*/
package main

import (
  "fmt"
)

type Node struct {
  Key int
  P *Node
  L *Node
  R *Node
}

type Bintree struct {
  Root *Node
}

func main() {
  tree := &Bintree{}
  fmt.Println("Inserting keys from static array ...")
  arr := []int{15, 6, 3, 18, 7, 4, 13, 9, 20, 17, 2}
  for _,v := range arr {
    tree.insert(&Node{Key: v})
  }
  fmt.Printf("Root of the tree has key %d\n", tree.Root.Key)
  tree.inorderwalk(tree.Root)
  fmt.Println("\nRecursively searching for node with key 3.")
  x := tree.search_rec(tree.Root,3)
  if x != nil {
    fmt.Printf("x is found. x.Key = %d\n", x.Key)
  }
  fmt.Println("Iteratively searching for node with key 4.")
  x = tree.search(tree.Root,4)
  if x != nil {
    fmt.Println("Deleting node with key 4.")
    x = tree.delete(x)
    fmt.Printf("x with key = %d is deleted\n", x.Key)
  }
  fmt.Println("Printing b-tree in order.")
  tree.inorderwalk(tree.Root)
  fmt.Println()
}

/* O(lg n) 
The pointer x traces the path and the pointer y is maintained as the parent of x. */
func (tree *Bintree) insert(z *Node) {
  var y *Node
  x := tree.Root
  for x != nil {
    y = x
    if z.Key <= x.Key {
      x = x.L
    } else {
      x = x.R
    }
  }
  if y == nil {
    tree.Root = z
  } else if z.Key <= y.Key {
    y.L = z
  } else {
    y.R = z
  }
  z.P = y
}

/* O(n) */
func (tree *Bintree) inorderwalk(x *Node) {
  if x != nil {
    tree.inorderwalk(x.L)
    fmt.Printf("%3d, ",x.Key)
    tree.inorderwalk(x.R)
  }
}

/* O(lg n). Search tree for node with key k starting from root x */
func (tree *Bintree) search_rec(x *Node, k int) *Node {
  if x == nil || x.Key == k {
    return x
  }
  if k < x.Key {
    return tree.search_rec(x.L,k)
  } else {
    return tree.search_rec(x.R,k)
  }
}

/* O(lg n): Iterative approach of search (more efficient on most computers) */
func (tree *Bintree) search(x *Node, k int) *Node {
  for x != nil && k != x.Key {
    if k < x.Key {
      x = x.L
    } else {
      x = x.R
    }
  }
  return x
}

/* O(lg n). Min element in a tree. Can be found by following left child pointers from
   root until a nil is found. */
func (tree *Bintree) minimum(x *Node) *Node {
  if x == nil { return x }
  for x.L != nil {
    x = x.L
  }
  return x
}

func (tree *Bintree) maximum(x *Node) *Node {
  if x == nil {return x }
  for x.R != nil {
    x = x.R
  }
  return x
}

/* O(lg n). predecessor is symmetric */
func (tree *Bintree) successor(x *Node) *Node {
  if x.R != nil {
    return tree.minimum(x.R)
  }
  y := x.P // get the parent of x
  for y != nil && x == y.R {
    x, y = y, y.P
  }
  return y
}

/*
O(lg n). Input: pointer to node to be deleted.
The procedure considers three cases:
1) If z has no children, we modify its parent z.P to replace z with nil as its child.
2) If the node has only a single child, we "splice out" z by making a new link between its child and its parent.
3) If the node has two children, we splice out z's successor y, which has no left child and replace z's key
and satellite data with y's key and satellite data.
*/
func (tree *Bintree) delete(z *Node) *Node {
  if z == nil {
    return z
  }
  var x, y *Node
  if z.L == nil || z.R == nil {
    y = z
  } else {
    y = tree.successor(z)
  }
  if y.L != nil {
    x = y.L
  } else {
    x = y.R
  }
  if x != nil {
    x.P = y.P
  }
  if y.P == nil {
    tree.Root = x
  } else if y == y.P.L {
    y.P.L = x
  } else {
    y.P.R = x
  }
  if y != z {
    z.Key = y.Key
    // code to copy y's satellite data into z
  }
  return y
}
