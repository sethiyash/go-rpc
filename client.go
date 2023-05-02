package main

import (
	"fmt"
	"net/rpc"
)

// Student struct represents a student.
type Student struct {
	ID                  int
	FirstName, LastName string
}

// FullName returns the fullname of the student.
func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

/*---------------*/

// College struct represents a college.
type College struct {
	database map[int]Student // private
}

// Add methods adds a student to the college (procedure).
func (c *College) Add(payload Student, reply *Student) error {

	// check if student already exists in the database
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("student with id '%d' already exists", payload.ID)
	}

	// add student `p` in the database
	c.database[payload.ID] = payload

	// set reply value
	*reply = payload

	// return `nil` error
	return nil
}

// Get methods returns a student with specific id (procedure).
func (c *College) Get(payload int, reply *Student) error {

	// get student with id `p` from the database
	result, ok := c.database[payload]

	// check if student exists in the database
	if !ok {
		return fmt.Errorf("student with id '%d' does not exist", payload)
	}

	// set reply value
	*reply = result

	// return `nil` error
	return nil
}

// NewCollege function returns a new instance of College (pointer).
func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}


func main() {

	// get RPC client by dialing at `rpc.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTPPath("tcp", "127.0.0.1:9001", "/rpc")

	/*--------------*/

	// create john variable of type `common.Student`
	var john Student

	// get student by id `1`
	if err := client.Call("College.Get", 1, &john); err != nil {
		fmt.Println("Error:1 College.Get()", err)
	} else {
		fmt.Printf("Success:1 Student '%s' found with id=1 \n", john.FullName())
	}

	/*--------------*/

	// add student by id `1`
	if err := client.Call("College.Add", Student{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}, &john); err != nil {
		fmt.Println("Error:2 College.Add()", err)
	} else {
		fmt.Printf("Success:2 Student '%s' created with id=1 \n", john.FullName())
	}

}
