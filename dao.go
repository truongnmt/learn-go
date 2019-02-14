package main

import (
	"fmt"
)

type S struct {
	name string
}

func (z *S) toD(d *S){
	fmt.Println("-------- Begin toD --------")
	fmt.Println("z before ---- ", &z, z)
	z = d
	fmt.Println("z after ---- ", &z, z)
	fmt.Println("-------- After toD --------")
}

func assignPointer(z *S) {
	fmt.Println("-------- Begin assignPointer --------")
	fmt.Println("z before ---- ", &z, z)
	d := &S{name: "im d"}
	fmt.Println("d ---- ", &d, d)
	z = d
	fmt.Println("z after ---- ", &z, z)
	fmt.Println("-------- After assignPointer --------")
}

//func main() {
//	fmt.Println("-------- Begin main --------")
//	z := &S{name: "im z"}
//	d := &S{name: "im d"}
//
//	fmt.Println("z before ------ ", &z, z)
//	assignPointer(z)
//	fmt.Println("z after assignPointer ------ ", &z, z)
//	z.toD(d)
//	fmt.Println("z after toD ------ ", &z, z)
//	fmt.Println("-------- After main --------")
//}