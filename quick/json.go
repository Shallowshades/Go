package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string   `json:"title"`
	Year  int      `json:"year"`
	Price int      `json:"rmb"`
	Actor []string `json:"actors"`
}

func main() {

	movie := Movie{"Final Fansty", 2007, 10, []string{"Cloud", "Tifa", "Alice"}}

	//编码过程 struct->json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码过程 json->struct
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("myMovie = %v\n", myMovie)
}
