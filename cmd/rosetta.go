package cmd

import (
	"encoding/json"
	"fmt"
)

type Vaca struct {
	Foo string
	Bar string
}

type Bacalao struct {
	Foo   string
	Bar   string
	Child Vaca
}

func Execute() {
	g := Bacalao{"a", "a", Vaca{"amarillo", "a"}}
	ret, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(ret))
		/*
			{
			 "name": "satori",
			 "age": 16,
			 "gender": "f",
			 "Where": "Oriental Diling hall",
			 "is_married": false
			}
		*/
	}
}
