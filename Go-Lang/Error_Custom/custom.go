package main

import "fmt"

type validError struct {
	Message string
}

func (v *validError) Error() string {
	return v.Message
}

type notFoundErr struct {
	Message string
}

func (n *notFoundErr) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validError{"Validation Error"}
	}

	if id != "dodo" {
		return &notFoundErr{"Data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("dodo", nil)
	if err != nil {
		// // terjadi error
		// if validationEr, ok := err.(*validError); ok {
		// 	fmt.Println("validation error:", validationEr.Error())
		// } else if notFoundEr, ok := err.(*notFoundErr); ok {
		// 	fmt.Println("not found error:", notFoundEr.Error())
		// } else {
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *validError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundErr:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}
	} else {
		// sukses
		fmt.Println("Sukses")
	}
}
