package main

import "github.com/mjekson/GoTest/Test4/funccc"

func main() {
	results := funccc.ChanOwner()
	funccc.Consumer(results)
}
