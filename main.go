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
	if v1, ok := value1.(float64); ok {
		fmt.Printf("%f норм конвертировалось\n",v1)
	}

	if v2, ok := value2.(float64); ok {
		fmt.Printf("%f норм конвертировалось\n",v2)
	}
	
	if oper, ok := operation.(string); ok {
		fmt.Printf("%s норм конвертировалось\n",oper)
	}
}


