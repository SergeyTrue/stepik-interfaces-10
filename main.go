package main

import (
"fmt"
//"encoding/json"
//"os"
)

func readTask()(interface{}, interface{},interface{}){
	return 12.5, 3.4, "+"
}

func main() {
	var v1, v2 float64
	var ok bool
	var oper string
	value1, value2, operation := readTask()
	if v1, ok = value1.(float64); !ok {
		return
	}

	if v2, ok = value2.(float64); !ok {
		return
	}

	if oper, ok = operation.(string); ok {
		switch {
			case oper == "+":
				fmt.Printf("%.4f\n",v1+v2)
			case oper == "-":
				fmt.Printf("%.4f\n",v1-v2)
			case oper == "*":
				fmt.Printf("%.4f\n",v1*v2)
			case oper == "/":
				fmt.Printf("%.4f\n",v1/v2)
			default: 
				fmt.Println("неизвестная операция")
		}
	} else {
	return
	}
}

