package interactor

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/usecase/repository"
)

type SlackInteractor struct {
	SlackRepository repository.SlackRepository
}

func (interactor *SlackInteractor) FindSlackInfo(name string) (slack domain.Slack, err error) {
	slack, err = interactor.SlackRepository.FindByName(name)
	return
}
