package main

import "fmt"




func new_map1() map[string]int {

    m := make(map[string]int)

    m["k1"] = 1

    fmt.Printf("m: %p\n", &m)
    return m
}


func new_map2() map[string]int {

    n := map[string]int{"foo": 1, "bar": 2}

    return n
}


func display_map1(m map[string]int) {

    fmt.Printf("m: %p\n", &m)

    for key, val := range m {

        fmt.Printf("%s: %d\n", key, val)

    }

}


func display_map2(m *map[string]int) {

    fmt.Printf("m: %p\n", m)

    for key, val := range *m {

        fmt.Printf("%s: %d\n", key, val)

    }


}


func main() {


    m1 := new_map1()
    fmt.Printf("m1: %p\n", &m1)

    display_map1(m1)
    display_map2(&m1)



}
