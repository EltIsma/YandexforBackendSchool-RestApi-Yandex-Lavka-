package services

import (
	"github.com/EltIsma/YandexLavka/cmd/internal/dto"
	"github.com/EltIsma/YandexLavka/cmd/internal/models"
	"github.com/EltIsma/YandexLavka/cmd/internal/repository"
)

type CouriersList interface {
	Create(list []*models.Courier) error
	GetAll(limit int, offset int)([]*models.Courier, error)
	GetById(courierId int) (models.Courier, error)
	GetCouriersSalaryRating(courier_id int, start_date string, end_date string)(int, int,error)
}

type OrdersList interface {
	Create(list []*models.Order) error
	GetAll(limit int, offset int)([]*models.Order, error)
	GetById(orderId int) (models.Order, error)
	Update(ordCom dto.OrderCompleteDto) (int, error)
}

type Service struct {
	CouriersList
	OrdersList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		CouriersList: NewCouriersListService(repos.CouriersList),
		OrdersList: NewOrdersListService(repos.OrdersList), 
	}
}
