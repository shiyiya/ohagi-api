package usecase

import (
	"fmt"

	"github.com/haton14/ohagi-api/controller/schema"
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/haton14/ohagi-api/repository"
	"github.com/labstack/echo/v4"
)

type CreateFoodIF interface {
	Create(request schema.FoodRequestIF, logger echo.Logger) (entity.Food, error)
}

type Food struct {
	CreateFoodIF
}

type CreateFood struct {
	foodRepo repository.FoodIF
}

func NewFood(foodRepo repository.FoodIF) Food {
	return Food{
		CreateFood{foodRepo: foodRepo},
	}
}

func (u CreateFood) Create(request schema.FoodRequestIF, logger echo.Logger) (entity.Food, error) {
	food, err := entity.NewFood(0, request.GetName(), 0, request.GetUnit())
	if err != nil {
		return entity.Food{}, err
	}
	conflict, err := u.foodRepo.FindByNameUnit(food.Name(), food.Unit())
	if conflict != nil {
		return entity.Food{}, fmt.Errorf("conflict food, name: %s, unit: %s", food.Name(), food.Unit())
	}
	err = u.foodRepo.Save(&food)
	if err != nil {
		return entity.Food{}, fmt.Errorf("food save err: %s", err)
	}
	return food, nil
}