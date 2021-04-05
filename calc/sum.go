package calc

import "fmt"

func Sum(a int, b int) (int, error) {
	if a == 0 || b == 0 {
		// fmt.Println("Error: args is less than 0")
		return 0, fmt.Errorf("error: %s", "args is less than 0.")
	}
	return a + b, nil
}
