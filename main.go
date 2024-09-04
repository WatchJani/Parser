package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	var (
		name     = "Janko"
		lastName = "Kondic"
		year     = "21"
		tel      = "+386 66 311 063"
		sex      = "mail"
	)

	user := New(name, lastName, tel, year, sex)
	build := user.Build()

	bufIndex := make([]byte, 4096)

	CreateIndex(build, bufIndex, []int{3, 2, 1})

	fmt.Println(string(bufIndex))
}

type User struct {
	Name     string
	LastName string
	Old      string
	Tel      string
	Sex      string
}

func New(name, lastName, tel, old, sex string) User {
	return User{
		Name:     name,
		LastName: lastName,
		Old:      old,
		Tel:      tel,
		Sex:      sex,
	}
}

func (u *User) Build() []byte {
	var buf bytes.Buffer

	v := reflect.ValueOf(u).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()
		str := fmt.Sprintf("%v", fieldValue)
		value := fmt.Sprintf("%4d%s", len(str), str)
		buf.WriteString(value)
	}

	return buf.Bytes()
}

func CreateIndex(input, keyIndex []byte, index []int) error {
	counter, position := 0, make([]int, 0, 10)

	for counter < len(input) {
		num, err := strconv.Atoi(strings.TrimLeft(string(input[counter:counter+4]), " "))
		if err != nil {
			return err
		}

		counter += num + 4
		position = append(position, counter-num, counter)
	}

	counter = 0
	for i := 0; i < len(index); i++ {
		value := input[position[index[i]<<1]:position[(index[i]<<1)+1]]
		copy(keyIndex[counter:], value)
		counter += len(value)
	}

	return nil
}
