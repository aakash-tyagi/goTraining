package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PersonInfo struct {
	Name       string
	Age        int
	Gender     string
	Occupation string
}

func main() {

	piM := PersonInfo{
		Name:       "Aakash Tyagi",
		Age:        27,
		Gender:     "male",
		Occupation: "software engineer",
	}

	bs, err := json.Marshal(piM)
	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(bs)

	var piUm PersonInfo

	err = json.Unmarshal(bs, &piUm)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(piUm)

}
