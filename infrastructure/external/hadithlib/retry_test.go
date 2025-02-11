package hadithlib_test

import (
	"sayeed1999/social-connect-golang-api/infrastructure/external/hadithlib"
	"testing"
)

func TestGetHadithWithRetries(t *testing.T) {
	hadith, err := hadithlib.GetHadithWithRetries()

	if hadith == nil || err != nil {
		t.Fatalf("Get Hadith even failed with retries!!!")
	}
}
