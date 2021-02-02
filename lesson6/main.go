package main

import (
	"fmt"
	"github.com/evgeniyv6/c_algorithms/lesson6/trees"
	"log"
	"strings"
)

// Pairs description
type Pairs struct {
	A,B rune
}

// ZipStringsLikePy description
func ZipStringsLikePy(sl1, sl2 string) ([]Pairs, error) {
	if len(sl1) != len(sl2) {
		return nil, fmt.Errorf("the lengths must be the sameß")
	}
	zipped := make([]Pairs, len(sl1))
	for i, el := range sl1 {
		zipped[i] = Pairs{el, rune(sl2[i])}
	}
	return zipped, nil
}

// exercise 1
func simpleHashFunc(s string) int32 {
	var excl = "!"
	alinedLenWithS := strings.Repeat(excl, len([]rune(s)))

	pairs, err := ZipStringsLikePy(s, alinedLenWithS); if err != nil {
		log.Fatal("Cannot zip two strings!")
	}
	var xorSl []rune
	for _, p :=range pairs {
		xorSl = append(xorSl,p.A^p.B)
	}

	var res int32
	for _, ch := range xorSl {
		res +=ch
	}
	return res
}

func main() {
	var newBst = new(trees.BinarySearchTree)
	newBst.Insert(8); newBst.Insert(4); newBst.Insert(10);newBst.Insert(2)
	newBst.Insert(6); newBst.Insert(1);newBst.Insert(3);newBst.Insert(5)
	newBst.Insert(7); newBst.Insert(9); newBst.Insert(11)

	fmt.Print("Inorder: ")
	newBst.Root.InOrderPrint()
	fmt.Println()
	fmt.Print("Preorder: ")
	newBst.Root.PreOrderPrint()
	fmt.Println()
	fmt.Print("Postorder: ")
	newBst.Root.PostOrderPrint()
	fmt.Println()
	var N = 8
	fmt.Printf("Check %d in tree: ", N)
	fmt.Println(newBst.Search(N))

	fmt.Println()
	fmt.Println("Print the tree:")
	newBst.String()

	fmt.Println()
	hw := "hello world!"
	fmt.Printf("simpleHashFunc \"%s\" => %d", hw,simpleHashFunc(hw))
}
