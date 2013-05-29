package union_find

import (
	"fmt"
	"os"
)

const UF_SIZE = 100

type UF struct {
	id []int // access to component id
	count int // number of component
}

func createUF() *UF{
	uf := new(UF)
	uf.count = UF_SIZE
	uf.id = make([]int, UF_SIZE)
	for i := 0; i < UF_SIZE; i++ {
		uf.id[i] = i
	}
	return uf
}

func (uf *UF) find(p int) int {
	return uf.find(p)
}

func (uf *UF) areConnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UF) union(p, q int) {
	// put p and q together
	pID := uf.find(p)
	qID := uf.find(q)

	// If they are together, do nothing
	if pID == qID {
		return
	}

	// Otherwise make p link to q
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
	uf.count--
}

func getStdinInt() int, err {

}

func Run() {
	uf := createUF()
	fmt.Printf("created Union Find data set of size: %v\n", uf.count)

	n, err := getStdinInt()
}
