package main

import (
	"fmt"
)



type Expression interface {
	Compile() string
}



type EqualsExpression struct {
	one string
	two string
}

func (equelExp EqualsExpression) Compile() string {
	return "1 == 2"
}



type Query struct {

}
func (query Query) Filter(ex Expression) Query {
	return query
}



type Db int
func Connect() Db {
	return 1
}
func(db Db) Query() Query {
	return Query{}
}

type Field interface  {

}

type Model struct {
	Tablename string
}

func(model *Model) GetTablename() string {
	return model.Tablename
}
/*
type Integer struct {

}
*/

type Integer int
type Email string
func (e Email) Equals(email string) Expression {
	return EqualsExpression{string(e), email}
}

type User struct{
	Tablename string
	id Integer
	email Email
}


func main() {

	user := &User{}
	user.id = 1
	user.Tablename = "user"

	fmt.Println(user)

	db := Connect()

	_ = db.Query().Filter(user.email.Equals("123"))

	// fmt.Println(query.Compile())
}


