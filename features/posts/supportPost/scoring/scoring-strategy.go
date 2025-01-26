package scoring

import "sayeed1999/social-connect-golang-api/models"

type ScoringStrategy interface {
	ApplyScore(post *models.Post)
}
