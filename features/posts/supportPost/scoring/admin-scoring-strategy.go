package scoring

import "sayeed1999/social-connect-golang-api/models"

type adminScoringStrategy struct{}

func (s *adminScoringStrategy) ApplyScore(post *models.Post) {
	post.Score += 5
}
