package greet

import "fmt"

/* "SayHello... public because it start with upper case
print hello to user"
*/
func SayHello(name string) {
	res := fmt.Sprintf("hello %s", name)
	fmt.Println(res)

}

// a private method
func iamPrivate() {
	fmt.Println("I'm private")
}
