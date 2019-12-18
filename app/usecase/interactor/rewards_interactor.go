package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
)

type RewardInteractor struct {
	RewardRepository repository.RewardRepository
}

func (interactor *RewardInteractor) FindAllRewards() (rewards domain.Rewards, err error) {
	rewards, err = interactor.RewardRepository.FindAll()
	if err != nil {
		return
	}
	return
}

func (interactor *RewardInteractor) PatchById(reward domain.Reward) (success bool, err error) {
	success, err = interactor.RewardRepository.PatchById(reward)
	if err != nil {
		return
	}
	return
}
