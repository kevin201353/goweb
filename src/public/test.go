package main
import (
		"fmt"
		"unsafe"
	)

const (
	a="abc"
	b=len(a)
	c=unsafe.Sizeof(a)
)

//interface
type Books  struct {
	title string
	author string
}

type Phone interface {
	call()
}

type NokiaPhone struct {
	
}

func (nokiaPhone NokiaPhone)call(){
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {

}

func (iphone IPhone)call(){
	fmt.Println("I am iphone, i can call you!")
}
 
/*
type error interface {
	Error() string	
}
*/

type Error2 interface {
	Call2() string
}

type MyExError struct {
	code int
	ermsg string
}

func (er MyExError)Error() string {
	return fmt.Sprintf("%d%s%d", er.code, er.ermsg, 5)
}

func (er MyExError)Call2() string {
	return fmt.Sprintf("%s", "I am iphone, i can call you!!!!!!!")
}

func main(){
	kvs := map[string]string{"a":"apple", "b":"blue"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k,v)
	}
	var  book  Books
	book.title = "bbb"
	book.author = "zhao"
	fmt.Println("title:", book.title, "author:", book.author)

	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	var myer Error2
	myer = &MyExError{1001, "this is error 1001"}
	fmt.Println(myer.Call2())
}








