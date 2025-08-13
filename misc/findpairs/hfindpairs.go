/* Sat Feb  8 02:17:02 PM IST 2025 */
/* By: Victor Nammour */
package main

import (
  "fmt"
)

type Node struct {
  Index int
  Next *Node
}
type List struct {
  Head *Node
  Tail *Node
}
type Pair struct {
  Index1 int 
  Index2 int
}
type Hashtab struct {
  Htab []*List
}

func (l *List) PrintList(index int) {
  fmt.Printf("ht[%d]: ", index)
  for n := l.Head; n!=nil; n = n.Next {
    fmt.Printf("[%d]->", n.Index)
  }
  fmt.Println("/-")
}

func (l *List) Append(index int) {
  n := &Node{index,nil}
  if l.Head == nil {
    l.Head = n
  } else {
    l.Tail.Next = n
  }
  l.Tail = n
}

func (l *List) Prepend(index int) {
  n := &Node{index,nil}
  if l.Head == nil {
    l.Tail = n
  } else {
    n.Next = l.Head
  }
  l.Head = n
}

func Count(l* List) int {
  if l.Head == nil {
    return 0
  } 
  return 1 + Count(&List{l.Head.Next, l.Tail})
}

func (l *List) Delete(index int) {
  if l.Head == nil {
    return
  }
  if l.Head.Index == index {
    x := l.Head.Next
    l.Head.Next = nil
    l.Head = x
  } else {
    for x,y := l.Head, l.Head.Next; y != nil; x,y = x.Next, y.Next {
      if y.Index == index {
        x.Next = y.Next
        y.Next = nil
        break
      }
    }
  }
}

func hash(value int, size int) uint {
  if value < 0 {
    value *= -1
  }
  return uint(value % size)
}

func (ht *Hashtab) install(index, value int) {
  l,key := ht.lookup(value)
  if l == nil {
    l = &List{}
    l.Append(index)
    ht.Htab[key] = l
  } else {
    ht.Htab[key].Append(index)
  }
}

func (ht *Hashtab) PrintTab() {
  for i,l := range (ht.Htab) {
    if l != nil {
      l.PrintList(i)
    } else {
      fmt.Printf("%d: %v\n", i, l)
    }
  }
}

/* returns a pointer to a List at the hashed value, which is the 
index in the slice ht */
func (ht *Hashtab) lookup(value int) (*List,uint) {
  key := hash(value,len(ht.Htab))
  //fmt.Printf("value = %d, len(ht.Htab) = %d, key = %d\n", value, len(ht.Htab), key)
  return ht.Htab[key],key
}

func (ht *Hashtab) getnode(index,value int) *Node {
  var node *Node = nil
  l,key := ht.lookup(value)
  //fmt.Printf("is list null ? %v, key = %d\n", l == nil, key)
  //fmt.Printf("getnode: index = %d, value = %d\n", index, value)
  if l != nil {
    l.PrintList(int(key))
    for n:=l.Head; n != nil; n = n.Next {
      if n.Index == index {
        node = n
        break;
      }
    }
  }
  return node
}
/*
1) Loop over entry array
2) Get the node corresponding to the current index and value in input array from the hashtable
3) Delete the node to mark that it's been processed. This is invasive. It removes it from the hashset.
4) Get the sum's complement
5) Get the hashtable's entry for the complement, which will be a *List
6) Possibly loop over the hashtable's entry linked list, and get the
matching indices. Accumulate the sum in an output integer array slice.
7) Mark each processed index by setting it's value to -1
Note: an alternate way is to have a boolean "isprocessed" per node,
but since I don't have a circular array, I don't need the extra memory,
and I don't need to preserve the original indices, so setting to -1 is suitable.
*/
func (ht *Hashtab) findpairs(arr []int, sum int) []Pair {
  out := make([]Pair,0,10) // an int slice of an initial capacity of 10
  for i,val := range arr {
    l,_ := ht.lookup(val)
    l.Delete(i)
    comp := sum - val
    l = ht.Htab[hash(comp,len(ht.Htab))]
    if l != nil {
      for n := l.Head; n != nil; n = n.Next {
        if arr[n.Index] == comp {
          out = append(out,Pair{i,n.Index})
        }
      }
    }
  }
  return out
}
//func (ht *Hashtab) findpairs(arr []int, sum int) []Pair {
//  out := make([]Pair,0,10) // an int slice of an initial capacity of 10
//  for i,val := range arr {
//    //fmt.Printf("Getting node for index %d, value %d\n", i, val)
//    n := ht.getnode(i,val)
//    if n == nil {
//      fmt.Printf("I got a null node for index %d, value %d\n", i, val)
//    }
//    if n!= nil {
//      n.Index = -1
//    }
//    comp := sum - val
//    l := ht.Htab[hash(comp,len(ht.Htab))]
//    if l != nil {
//      for n := l.Head; n != nil; n = n.Next {
//        if n.Index != -1 && arr[n.Index] == comp {
//          out = append(out,Pair{i,n.Index})
//          //n.Index = -1
//        }
//      }
//    }
//  }
//  return out
//}
//
func main() {
  arr := []int{1,3,9,4,6,6,11,-1,-5,1}
  ht := &Hashtab{ Htab: make([]*List,10) }
  for i,v := range arr {
    //ht.lookup(v)
    ht.install(i,v)
  }
  ht.PrintTab()
  sum := 10
  out := ht.findpairs(arr,sum)
  fmt.Println("Index pairs:", out)
  ht.PrintTab()
}
