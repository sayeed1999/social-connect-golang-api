package hadithlib

import (
	"math/rand"
	"time"
)

// Requirement: -
// We have signed an MOU with the AS-Sunnah Foundation. We want to show hadith on our
// website using their one day one hadith service. They haven't provided us with their
// client library yet, so we are mocking it.

type hadith struct {
	ID        int
	Text      string
	Reference string
	Date      time.Time
}

func GetHadith() *hadith {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := rand.Intn(10) + 1

	text := "The Prophet (صلى الله عليه وسلم) says, \"The best among you are those who have the best manners and character. (صلى الله عليه وسلم)\""

	// Will return 'nil' on randomness when random number will be 5 out of 1-10.
	if randomInt == 5 {
		return nil
	}

	return &hadith{
		ID:        1,
		Text:      text,
		Reference: "Sahih Bukhari 3559",
		Date:      time.Now(),
	}
}
