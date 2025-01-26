package scoring

import "sayeed1999/social-connect-golang-api/models"

type regularScoringStrategy struct{}

func (s *regularScoringStrategy) ApplyScore(post *models.Post) {
	post.Score++
}
