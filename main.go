package main

import "fmt"

type DBRow map[string]string

type DataBase []DBRow

type DataBaser interface {
	Append(r DBRow) DataBase
	Delete(r DBRow)
	Get(i int64)
}

func (r DBRow) String() string {
	temp := ""
	for k, v := range r {
		temp = fmt.Sprintf("%s\n%s\t:\t%s\n", temp, k, v)
	}

	return temp
}

func (db DataBase) Append(r DBRow) DataBase {
	db = append(db, r)
	return db
}

func main() {

	var Jack DataBase

	Jack = Jack.Append(DBRow{"Age": "25"})
	fmt.Println(Jack)
}
