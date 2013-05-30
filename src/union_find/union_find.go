package union_find

import (
	"fmt"
)

// three implementations of union find
// quick find, quick union, and weighted quick union

const UF_SIZE = 10

var dataSet = []int{ // 2 component groups
	4, 3,
	3, 8,
	6, 5,
	9, 4,
	2, 1,
	5, 0,
	7, 2,
	6, 1,
}

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


// quick find implementation ---------------------
func quickFind(uf *UF) {
	for i := 0; i < UF_SIZE; i = i + 2 {
		p := dataSet[i]
		q := dataSet[i+1]
		if uf.QF_connected(p, q) {
			continue
		}
		uf.QF_union(p, q)
		fmt.Println("union:", p, q)
	}
	fmt.Println(uf.count, "components in data set")
}

func (uf *UF) QF_find(p int) int {
	return uf.id[p]
}

func (uf *UF) QF_union(p, q int) {
	// put p and q together
	pID := uf.QF_find(p)
	qID := uf.QF_find(q)
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

func (uf *UF) QF_connected(p, q int) bool {
	return uf.QF_find(p) == uf.QF_find(q)
}
// end quick find ---------------------------

func Run() {
	uf := createUF()
	fmt.Printf("created Union Find data set of size: %v\n", uf.count)
	quickFind(uf)
}
