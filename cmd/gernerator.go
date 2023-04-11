package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Define sets of characters that will be used to generate passwords digits 0-9, lowercase and uppercase letters and special characters
var numbers = "0123456789"
var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var specialChars = "~!@#$%^&*()_+{}|:?-=[];',./"

func generatePasswords(w http.ResponseWriter, r *http.Request) {
	// Parse input parameters from request body
	var params map[string]int
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate passwords by calling generatePassword Function.
	var passwords []string
	for i := 0; i < params["num_passwords"]; i++ {
		password := generatePassword(params["min_length"], params["special_chars"], params["numbers"])
		passwords = append(passwords, password)
	}

	// Return passwords in in an array format
	response := map[string]interface{}{"passwords": passwords}
	json.NewEncoder(w).Encode(response)
}

// Generate random password with specified parameters
func generatePassword(minLength, numSpecialChars, numNumbers int) string {
	rand.Seed(time.Now().UnixNano())

	var password strings.Builder

	for i := 0; i < minLength; i++ {
		if i < numSpecialChars {
			password.WriteByte(specialChars[rand.Intn(len(specialChars))])
		} else if i < numSpecialChars+numNumbers {
			password.WriteByte(numbers[rand.Intn(len(numbers))])
		} else {
			password.WriteByte(letters[rand.Intn(len(letters))])
		}
	}

	// shuffle password to make it more randomized
	passwordBytes := []byte(password.String())
	rand.Shuffle(len(passwordBytes), func(i, j int) {
		passwordBytes[i], passwordBytes[j] = passwordBytes[j], passwordBytes[i]
	})

	return string(passwordBytes)
}
