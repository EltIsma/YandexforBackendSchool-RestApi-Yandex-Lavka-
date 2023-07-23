package repository

import (
	"github.com/EltIsma/YandexLavka/cmd/internal/dto"
	"github.com/EltIsma/YandexLavka/cmd/internal/models"
	"github.com/go-pg/pg/v10"
	_ "github.com/jmoiron/sqlx"
)

type CouriersList interface {
	Create(list []*models.Courier) (error)
	GetAll(limit int, offset int)([]*models.Courier, error)
	GetById(courierId int) (models.Courier, error)
	GetCouriersSalaryRating(courier_id int, start_date string, end_date string)(models.Courier, int, error)
}

type OrdersList interface {
	Create(list []*models.Order) error
	GetAll(limit int, offset int)([]*models.Order, error)
	GetById(orderId int) (models.Order, error)
	Update(ordCom dto.OrderCompleteDto) (int, error)
}

type Repository struct {
	CouriersList
	OrdersList
}
func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		CouriersList: NewCouriersListPostgres(db),
		OrdersList: NewOrdersListPostgres(db),
	}
}
