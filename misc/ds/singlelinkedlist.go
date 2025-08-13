/* Wed Feb  5 10:45:34 AM IST 2025 */
/* By: Victor Nammour */
package ds

import (
	"fmt"
)

type Node[T any] struct {
	Data T
	Next *Node[T]
}
type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) Print(name string) {
	fmt.Printf("%s: ", name)
	for n := l.Head; n != nil; n = n.Next {
		fmt.Printf("[%v]->", n.Data)
	}
	fmt.Println("/-")
}

func (l *List[T]) Append(data T) {
	n := &Node[T]{data, nil}
	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func (l *List[T]) Prepend(data T) {
	n := &Node[T]{data, nil}
	if l.Head == nil {
		l.Tail = n
	} else {
		n.Next = l.Head
	}
	l.Head = n
}

func (l *List[T]) Reverse() {
	if l.Head == nil {
		return
	}
	var x, y, z *Node[T] = l.Head, l.Head.Next, nil
	for y != nil {
		z = y.Next
		y.Next = x
		x, y = y, z
	}
	l.Head.Next = nil
	l.Tail = l.Head
	l.Head = x
}

// func Count(l* List[T]) int {
//   if l.Head == nil {
//     return 0
//   }
//   return 1 + Count(&List[T]{l.Head.Next, l.Tail})
// }

func (l *List[T]) Count() int {
	return countHelper(l.Head)
}

func countHelper[T any](node *Node[T]) int {
	if node == nil {
		return 0
	}
	return 1 + countHelper(node.Next)
}

func (l *List[T]) Delete(data T) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == data {
		x := l.Head.Next
		l.Head.Next = nil
		l.Head = x
	} else {
		for x, y := l.Head, l.Head.Next; y != nil; x, y = x.Next, y.Next {
			if y.Data == data {
				x.Next = y.Next
				y.Next = nil
				break
			}
		}
	}
}
