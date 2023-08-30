package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Printf("co = {num: %v, str: %v}\n", co.num, co.str)
	fmt.Printf("also num: %v\n", co.base.num)
	fmt.Println("describe: ", co.describe())
	fmt.Println("base describe: ", co.base.describe())

	// Embedding structs with methods may be used to bestow 
	// interface implementations onto other structs. Here we 
	// see that a container now implements the describer 
	// interface because it embeds base.
	type describer interface {
		describe() string
	}
	// why d type is container? not interface?
	var d describer = co
	fmt.Printf("co type is %T\n", co)
	fmt.Printf("d type is %T\n", d)
	fmt.Println("describer: ", d.describe())
	var f describer = co.base
	fmt.Println("base describer: ", f.describe())
	fmt.Printf("base type %T\n", f)	
}