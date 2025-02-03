package controllers

import (
	"HNG_PROJECT_2/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"
)

// function to check if number enter is a prime number
func IsPrimeNumber(number int) bool {

	for num := 2; num < number; num++ {
		if number%num == 0 {
			return false
		}
	}

	return true
}

// function to check if number is perfect
func IsPerfectNumber(number int) bool {
	if number < 2 {
		return false
	}

	sum := 0

	for num := 1; num < number; num++ {
		if number%num == 0 {
			sum += num
		}
	}

	return sum == number
}

// funnction to check and add the properties of a numeber
func NumberProperties(number int) []string {
	Number := strconv.Itoa(number)
	numberLength := len(Number)

	sum := 0.0
	var properties []string

	for _, digit := range strings.Split(Number, "") {
		d, _ := strconv.Atoi(digit)

		sum += math.Pow(float64(d), float64(numberLength))
	}

	if int(sum) == number {
		properties = append(properties, "armstrong")
	}

	if number%2 != 0 {
		properties = append(properties, "odd")
	} else {
		properties = append(properties, "even")
	}

	return properties
}

// function to sum up the digits pf the number
func SumOfNumber(number int) int {
	Number := strconv.Itoa(number)
	sum := 0

	for _, digit := range Number {
		digitInt, _ := strconv.Atoi(string(digit))

		sum += digitInt
	}

	return sum
}

// function to retrive fun fact from numberapi
func FunFactRequest(number int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", number)

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		return "Could not get fun fact"
	}
	defer resp.Body.Close() // Close response body after function exits

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "can not read response"
	}

	// Convert response to a string and return
	return string(body)
}

// the main controller function
func NumberClassificationController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	number := r.URL.Query().Get("number")

	if number == "" {
		http.Error(w, "please enter a value to proceed", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(number)

	if err != nil {
		response := models.Mathfailed{
			Number: "alphabet",
			Error:  true,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := models.Mathsucess{
		Number:     num,
		IsPrime:    IsPrimeNumber(num),
		IsPerfect:  IsPerfectNumber(num),
		Properties: NumberProperties(num),
		DigitSum:   SumOfNumber(num),    //sum of its digits
		FunFact:    FunFactRequest(num), //gotton from the numbers Api
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
