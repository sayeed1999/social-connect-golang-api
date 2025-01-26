package scoring

import "sayeed1999/social-connect-golang-api/models"

type ScoringStrategyFactory struct{}

func (f *ScoringStrategyFactory) GetScoringStrategy(user *models.User) ScoringStrategy {

	if user.IsAdmin != nil && *user.IsAdmin {
		return &adminScoringStrategy{}
	}

	return &regularScoringStrategy{}
}
