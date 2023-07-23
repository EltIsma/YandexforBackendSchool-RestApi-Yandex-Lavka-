package services

import (
	"github.com/EltIsma/YandexLavka/cmd/internal/models"
	"github.com/EltIsma/YandexLavka/cmd/internal/repository"
)

type CouriersService struct {
	repo repository.CouriersList
}


func NewCouriersListService(repo repository.CouriersList) *CouriersService{
	return &CouriersService{repo: repo}
}

func (s *CouriersService) Create(list []*models.Courier) (error) {
	return s.repo.Create(list)
}
func (s *CouriersService) GetAll(limit int, offset int) ([]*models.Courier, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *CouriersService) GetById(courierId int) (models.Courier, error){
	return s.repo.GetById(courierId)
}

func(s *CouriersService) GetCouriersSalaryRating(courier_id int, start_date string, end_date string)(int, int, error){
	metacourier, countOrdersCompleted, err := s.repo.GetCouriersSalaryRating(courier_id, start_date, end_date)
   if err != nil{
      return 0,0 ,  err
   }
	earning := metacourier.Earnings(countOrdersCompleted)
	rating := metacourier.Rating(countOrdersCompleted, start_date, end_date)
	earn :=earning
	rat := rating

	return earn, rat, nil
}



