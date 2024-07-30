package main

import (
	"errors"
	"exception-handling/pkg"
	"fmt"
)

type Data struct {
	ID   int
	Name string
}

func getDataFromMockDB(id int) *Data {
	// Mock data
	mockData := map[int]*Data{
		1: {ID: 1, Name: "Data 1"},
		2: {ID: 2, Name: "Data 2"},
		3: {ID: 3, Name: "Data 3"},
	}

	// Mengembalikan data jika ditemukan, atau nil jika tidak ditemukan
	return mockData[id]
}

func findDataByID(id int) (*Data, error) {
	data := getDataFromMockDB(id) // Asumsi fungsi ini mencari data dari database
	if data == nil {
		return nil, &pkg.ErrorConstant.NotFound
	}
	return data, nil
}

func validateInput(input string) error {
	if input == "" {
		return &pkg.ErrorConstant.Validation
	}
	return nil
}

func processSomething() error {
	return errors.New("terjadi kesalahan tak terduga")
}

func validationError() {
	err := validateInput("")
	if err != nil {
		errorResponse := pkg.ErrorResponse(err)
		fmt.Println(errorResponse.Response.Meta.Message)
		fmt.Println(errorResponse.Code)
		return
	}
	fmt.Println("Input valid")
}

func authenticateUser(username, password string) error {
	if username != "user" || password != "password" {
		return &pkg.ErrorConstant.EmailOrPasswordIncorrect
	}
	return nil
}

func searchData() {
	data, err := findDataByID(5)
	if err != nil {
		errorResponse := pkg.ErrorResponse(err)
		fmt.Println(errorResponse.Response.Meta.Message)
		fmt.Println(errorResponse.Code)
		return
	} else {
		fmt.Println(data)
	}
}

func auth() {
	err := authenticateUser("wrongUser", "wrongPassword")
	if err != nil {
		errorResponse := pkg.ErrorResponse(err)
		fmt.Println(errorResponse.Response.Meta.Message)
		fmt.Println(errorResponse.Code) // Output: 401
		return
	}
	fmt.Println("Autentikasi berhasil")
}

func simulateIntError() {
	err := processSomething()
	if err != nil {
		errorResponse := pkg.ErrorResponse(err)
		fmt.Println(errorResponse.Response.Meta.Message)
		fmt.Println(errorResponse.Code)
		return
	}
	fmt.Println("Process success")
}

// Panic
func divide(a, b int) int {
	if b == 0 {
		panic("division by zero is not allowed")
	}
	return a / b
}

// Recover from panic
func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	if b == 0 {
		panic("division by zero is not allowed")
	}
	return a / b
}

func divideByZero() {
	divide(10, 0)
}

func divideByZeroRecover() {
	result := safeDivide(10, 0)
	fmt.Println(result)
}

func main() {
	// auth()
	// searchData()
	// validationError()
	// simulateIntError()
	// divideByZero()
	divideByZeroRecover()
}
