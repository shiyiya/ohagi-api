package schema

import (
	"github.com/haton14/ohagi-api/domain/entity"
	"github.com/labstack/echo/v4"
)

type food struct {
	ID            *int    `json:"id,omitempty"`
	Name          string  `json:"name"`
	Amount        float64 `json:"amount,omitempty"`
	Unit          string  `json:"unit"`
	LastUpdatedAt *int64  `json:"last_updated_at,omitempty"`
}

type FoodResponse struct {
	c echo.Context
	food
}
type FoodRequestIF interface {
	GetName() string
	GetUnit() string
}
type FoodResponseIF interface {
	JSON(code int) error
	FoodRequestIF
}

func NewFoodRequest(c echo.Context) (FoodRequestIF, error) {
	s := food{}
	if err := c.Bind(&s); err != nil {
		return nil, err
	}
	return s, nil
}

func NewFoodResponse(c echo.Context, f entity.Food) FoodResponseIF {
	id := f.ID()
	return FoodResponse{c, food{ID: &id, Name: f.Name(), Unit: f.Unit()}}
}

func (s food) GetName() string {
	return s.Name
}

func (s food) GetUnit() string {
	return s.Unit
}
func (s FoodResponse) JSON(code int) error {
	return s.c.JSON(code, s.food)
}
