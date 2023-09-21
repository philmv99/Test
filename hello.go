package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Operation struct {
	Operator string    `json:"operator"`
	Operands []float64 `json:"operands"`
}

func main() {
	url := "https://jlsherrill.fedorapeople.org/test.json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var operations []Operation
	if err := json.Unmarshal(body, &operations); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, op := range operations {
		var result float64
		operand1 := op.Operands[0]
		operand2 := op.Operands[1]

		switch op.Operator {
		case "addition":
			result = operand1 + operand2
		case "subtraction":
			result = operand1 - operand2
		case "multiplication":
			result = operand1 * operand2
		case "division":
			if operand2 == 0 {
				fmt.Println("Cannot divide by zero!")
				continue
			}
			result = operand1 / operand2
		default:
			fmt.Println("Unknown operator:", op.Operator)
			continue
		}

		fmt.Printf("%v %s %v = %v\n", operand1, op.Operator, operand2, result)
	}
}
