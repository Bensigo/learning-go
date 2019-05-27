package main

import "fmt"

// struct
// a strut is a typed,collection of diffrent field. is a grouped of similar data relating
// like  person - has a  name, age, job

type person struct {
	name    string
	age     uint16
	job     string
	married bool
}

// creating a instance of person from the struct
var sigo = person{name: "Sigo Egwey", age: 25, job: "software engineer", married: false}

// accessing with a pointer
var sigo2 = &person{name: "Sigo Egwey", age: 25, job: "software engineer", married: false}

// method are special method also called reciver which a reciver can be pass a value or pointer
func (user *person) describe() string {
	discribe := fmt.Sprintf("My name is %s, i'm  %d %s", user.name, user.age, user.job)
	return discribe
}

// passing a pointer to the reciver
func (user person) setAge(age uint16) uint16 {
	user.age = age
	return user.age
}

// passing a value to the reciver
func (user person) setCareer(job string) string {
	user.job = job
	return job
}

// interface in go is a collection of method
type user interface {
	setAge(age uint16) uint16
	setCareer(job string) string
}

func main() {
	// accessing value from a struct
	var ben user
	ben = person{name: "Sigo Egwey", age: 25, job: "software engineer", married: false}
	fmt.Println(ben.setCareer("full stack engineer"))
	fmt.Println(ben.setAge(27))

}
