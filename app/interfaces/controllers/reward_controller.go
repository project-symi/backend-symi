package controllers

import (
	"project-symi-backend/app/domain"
	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase/interactor"
)

type RewardController struct {
	Interactor interactor.RewardInteractor
}

func NewRewardController(sqlHandler database.SqlHandler) *RewardController {
	return &RewardController{
		Interactor: interactor.RewardInteractor{
			RewardRepository: &database.RewardRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *RewardController) AllRewards(c Context) {
	rewards, err := controller.Interactor.FindAllRewards()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, rewards)
}

func (controller *RewardController) PatchRewardById(c Context) {
	reward := domain.Reward{}
	c.BindJSON(&reward)
	rewards, err := controller.Interactor.PatchById(reward)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, rewards)
}
