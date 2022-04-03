package main

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		//есть место в базовом массиве поэтому увеличиваем срез
		z = x[:zlen]
	} else {
		//увеличиваем срез в два раза
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {

	// i := 1

	// defer func() {
	// 	fmt.Println(i)
	// }()

	// i++
	// fmt.Println(i)
	// var m map[string]int
	// r := map[string]int{}

	// m["first"] = 1
	// r["first"] = 1
	//
	// Printf("Value r:%#v type:%T\n", r, r)
	// slice := []int{5, 6, 7, 8, 9}
	// fmt.Println(cap(slice))
	// fmt.Println(len(slice))
	// fmt.Println(remove(slice, 2))
	// var x, y []int
	// for i := 0; i < 10; i++ {
	// 	y = appendInt(x, i)
	// 	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }
	// x := make([]int, 3)
	// fmt.Printf("len(x)==0 :%v type %T", len(x), x)
}
