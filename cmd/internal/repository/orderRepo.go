package repository

import (
	"time"

	"github.com/EltIsma/YandexLavka/cmd/internal/dto"
	"github.com/EltIsma/YandexLavka/cmd/internal/models"
	"github.com/go-pg/pg/v10"
	_ "github.com/go-pg/pg/v10/orm"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type OrdersPostgres struct {
	db *pg.DB
}

func NewOrdersListPostgres(db *pg.DB) *OrdersPostgres {
	return &OrdersPostgres{db: db}
}

func (r *OrdersPostgres) Create(list []*models.Order) error {
	_, err := r.db.Model(&list).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (r *OrdersPostgres) GetAll(limit int, offset int) ([]*models.Order, error) {
	var ordersList []*models.Order
	err := r.db.Model(&ordersList).Limit(limit).Offset(offset).Select()
	if err != nil {
		return nil, err
	}
	return ordersList, nil
}

func (r *OrdersPostgres) GetById(orderId int) (models.Order, error) {
	var orderInfo models.Order
	err := r.db.Model(&orderInfo).Where("id = ?", orderId).Select()
	if err != nil {
		return orderInfo, err
	}
	return orderInfo, nil
}

func (r *OrdersPostgres) Update(ordCom dto.OrderCompleteDto) (int, error) {
	var orderInfo models.Order
	completeTime, err := time.Parse("2006-01-02", ordCom.OrderCompleteTime)
	if err != nil {
		return 0, err
	}
	_, err = r.db.Model(&orderInfo).Set("courier_id = ?", ordCom.CourierId).Set("complete_time = ?", completeTime).Where("id = ?", ordCom.OrderId).Update()
	return ordCom.OrderId, err

}
