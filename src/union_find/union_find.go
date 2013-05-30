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
	5, 6,
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


// quick find implementation

type quickFind struct {
	*UF
}

func (qf *quickFind) execute() {
	for i := 0; i < len(dataSet); i = i + 2 {
		p := dataSet[i]
		q := dataSet[i+1]
		qf.union(p, q)
	}
	fmt.Println(qf.count, "components in data set")
}

func (qf *quickFind) find(p int) int {
	return qf.id[p]
}

func (qf *quickFind) union(p, q int) {
	// put p and q together
	pID := qf.find(p)
	qID := qf.find(q)
	// If they are together, do nothing
	if pID == qID {
		fmt.Println(p, q, "already part of same group")
		return
	}
	fmt.Println("union:", p, q)
	// Otherwise make p link to q
	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pID {
			qf.id[i] = qID
		}
	}
	qf.count--
}

func (qf *quickFind) connected(p, q int) bool {
	return qf.find(p) == qf.find(q)
}

// quick union implementation

type quickUnion struct {
	*UF
}

func (qu *quickUnion) execute() {
	for i := 0; i < len(dataSet); i = i + 2 {
		p := dataSet[i]
		q := dataSet[i+1]
		qu.union(p, q)
	}
	fmt.Println(qu.count, "components in data set")
}

func (qu *quickUnion) find(p int) int {
	// chase down the root node
	for p != qu.id[p] {
		p = qu.id[p]
	}
	return p
}

func (qu *quickUnion) union(p, q int) {
	pRoot := qu.find(p)
	qRoot := qu.find(q)
	if pRoot == qRoot {
		fmt.Println(pRoot, qRoot, "already part of same group")
		return
	}
	fmt.Println("union:", p, q)
	qu.id[pRoot] = qRoot
	qu.count--
}

func Run() {
	uf := createUF()
	fmt.Printf("created Union Find data set of size: %v\n", uf.count)
	qf := quickFind{UF: uf}
	fmt.Println("running 'Quick Find'")
	qf.execute()
	fmt.Println("running 'Union Find'")
	uf = createUF()
	qu := quickUnion{UF: uf}
	qu.execute()
}
