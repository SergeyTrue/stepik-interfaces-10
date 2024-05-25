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
	value1, value2, operation := readTask()
	var v1, v2 float64
	var ok bool
	var oper string
	if v1, ok = value1.(float64); !ok {
		fmt.Printf("value=%v: %T\n", value1, value1)
		return
	}

	if v2, ok = value2.(float64); !ok {
		fmt.Printf("value=%v: %T\n", value2, value2)
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

