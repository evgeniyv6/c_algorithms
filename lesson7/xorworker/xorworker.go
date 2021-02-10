package xorworker

import (
	b64 "encoding/base64"
	"fmt"
	"log"
)

// Pairs descr
type Pairs struct {
	A,B rune
}

// ZipStringsLikePy description
func ZipStringsLikePy(sl1, sl2 string) ([]Pairs, error) {
	if len(sl1) != len(sl2) {
		return nil, fmt.Errorf("the lengths must be the same")
	}
	zipped := make([]Pairs, len(sl1))
	for i, el := range sl1 {
		zipped[i] = Pairs{el, rune(sl2[i])}
	}
	return zipped, nil
}

//     xored = ''.join((chr(ord(x) ^ ord(y)) for (x,y) in zip(password,itertools.cycle(key))))


// XorFunc descr
func XorFunc(s string, key...string) string {
	keySec := "________" // default value
	if len(key) > 0 {
		keySec = key[0]
	}
	sDec, _ := b64.StdEncoding.DecodeString(s)
	fmt.Println(string(sDec))
	fmt.Println()

	pairs, err := ZipStringsLikePy(string(sDec), keySec); if err != nil {
		log.Fatal("cannot zip two strings")
	}
	fmt.Println(pairs)
	var ordSl []string
	for _, p :=range pairs {
		fmt.Println(p.A^p.B)
		ordSl = append(ordSl,string(p.A^p.B))
	}
	fmt.Println(" ordSl", ordSl)
	return ""
}
