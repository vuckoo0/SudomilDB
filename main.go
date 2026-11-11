package main

import (
	"fmt"
)

type DBRow []any

type DB map[string]DBRow

type DataBaser interface {
	AppendElement(elem any, key string)
	DeleteElement(elem any, key string) error
	GetElement(elem any, key string) int64
}

func (db DB) String() string {
	temp := ""
	for k, v := range db {
		temp = fmt.Sprintf("%s\n%s\t:\t[", temp, k)
		for i := range v {
			temp = fmt.Sprintf("%s %v", temp, v[i])
		}
		temp = fmt.Sprintf("%s ]", temp)
	}

	return temp
}

func (db *DB) AppendElement(elem any, key string) {
	if (*db) == nil {
		*db = make(DB)
	}

	(*db)[key] = append((*db)[key], elem)
}

func (db *DB) DeleteElement(elem any, key string) error {
	_, ok := (*db)[key]
	if !ok {
		return fmt.Errorf("No row with name \"%s\" found!", key)
	}

	return nil
}

func main() {

	var Jack DB

	Jack.AppendElement([]int{25, 65, 29, 54, 38}, "Age")
	Jack.AppendElement("dj", "Age")
	fmt.Println(Jack)
}
