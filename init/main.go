package main

import "fmt"

func setQueryArgs(args ...string) [][]byte {
	var params []string
	for _, arg := range args {
		params = append(params, arg)
	}

	queryArgs := make([][]byte, len(params))
	for i, p := range params {
		queryArgs[i] = []byte(p)
	}

	return queryArgs
}

func main() {

	a := setQueryArgs("a", "ba", "c")
	fmt.Println(a)

}
