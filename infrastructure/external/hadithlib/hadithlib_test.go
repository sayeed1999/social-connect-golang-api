package hadithlib_test

import (
	"sayeed1999/social-connect-golang-api/infrastructure/external/hadithlib"
	"testing"
)

// TestGetHadithRandomness tests that calling GetHadith 100 times returns at least one empty string for Text
func TestGetHadithRandomness(t *testing.T) {
	// Set a counter to track how many empty string got from GetHadith()
	nilHadithCount := 0

	for i := 0; i < 100; i++ {
		// hadithlib.GetHadith() returns a pointer to the result, pointer can be 'nil' if randomInd = 5
		// &hadithlib.GetHadith() would return the value, should not be used as GetHadith() doesn't always ensure a value!
		hadith, err := hadithlib.GetHadith()

		if hadith == nil || err != nil {
			nilHadithCount++
		}
	}

	if nilHadithCount == 0 {
		t.Fatalf("Cannot have zero nil hadiths...")
	}

	if nilHadithCount > 20 {
		t.Fatalf("Cannot have more than 20%% nil hadiths...")
	}
}
