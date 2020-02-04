package main

import "fmt"

func main() {
	fmt.Println("hello world");

	var a bool = true;
	fmt.Println(a);

	var b int = 1;
	fmt.Println(b);

	var c uint = 2;
	fmt.Println(c);

	var d int8 = 3;
	fmt.Println(d);

	var e int16 = 4;
	fmt.Println(e);
	
	var f int32 = 5;
	fmt.Println(f);
	
	var g int64 = 6;
	fmt.Println(g);


	var h uint8 = 7;
	fmt.Println(h);

	var i uint16 = 8;
	fmt.Println(i);

	var j uint32 = 9;
	fmt.Println(j);

	var k uint64 = 10;
	fmt.Println(k);

	var hello string = "hello";
	fmt.Println(hello);

	var world = "world";
	fmt.Println(world);

	var l int
    var m float64
    var n bool
    var o string
    fmt.Printf("%v %v %v %q\n", l, m, n, o);

}