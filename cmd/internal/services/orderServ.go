package services

import (
	"github.com/EltIsma/YandexLavka/cmd/internal/dto"
	"github.com/EltIsma/YandexLavka/cmd/internal/models"
	"github.com/EltIsma/YandexLavka/cmd/internal/repository"
)

type OrdersService struct {
	repo repository.OrdersList
}


func NewOrdersListService(repo repository.OrdersList) *OrdersService{
	return &OrdersService{repo: repo}
}

func (s *OrdersService) Create(list []*models.Order) (error) {
	return s.repo.Create(list)
}
func (s *OrdersService) GetAll(limit int, offset int) ([]*models.Order, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *OrdersService) GetById(orderId int) (models.Order, error){
	return s.repo.GetById(orderId)
}
func (s *OrdersService)  Update(ordCom dto.OrderCompleteDto) (int, error){
	return s.repo.Update(ordCom)
}