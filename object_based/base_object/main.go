package main


import (

	"fmt"
	"log"
)



type Object struct {

	Id int

}

func New() (*Object, error) {

	obj := &Object{
		Id: 1,

	}

    return obj, nil
}

func (this *Object) GetId() int {


	return this.Id
}

func (this *Object) SetId(id int) error {

	this.Id = id

	return nil
}






func main() {

	obj, err := New()
	if  err != nil {
		log.Fatalf("new error")
	}

	id := obj.GetId()
	fmt.Println("id: ", id)

	obj.SetId(2)
	id1 := obj.GetId()
	fmt.Println("id: ", id1)


}