package hadithlib

import (
	"errors"
	"fmt"
	"sayeed1999/social-connect-golang-api/infrastructure/external/proxy"
)

func GetHadithWithRetries() (*Hadith, error) {

	// Define an operation function that returns a value and an error.
	operation := func() (interface{}, error) {
		hadith, err := GetHadith()

		if err != nil {
			return hadith, err
		}

		return hadith, nil
	}

	// Use ExecuteWithRetry to handle retries
	result, err := proxy.ExecuteWithRetry(operation)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Cast the result to Hadith struct
	hadithResult, ok := result.(*Hadith)
	if !ok {
		return nil, errors.New("failed to cast result to *Hadith")
	}

	return hadithResult, nil
}
