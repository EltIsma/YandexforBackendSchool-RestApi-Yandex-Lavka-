package repository

import (
	"github.com/EltIsma/YandexLavka/cmd/internal/models"
	"github.com/go-pg/pg/v10"
	_ "github.com/go-pg/pg/v10/orm"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CouriersPostgres struct {
	db *pg.DB
}

func NewCouriersListPostgres(db *pg.DB) *CouriersPostgres {
	return &CouriersPostgres{db: db}
}

func (r *CouriersPostgres) Create(list []*models.Courier) error {
	_, err := r.db.Model(&list).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (r *CouriersPostgres) GetAll(limit int, offset int) ([]*models.Courier, error) {
	var couriersList []*models.Courier
	err := r.db.Model(&couriersList).Limit(limit).Offset(offset).Select()
	if err != nil {
		return nil, err
	}
	return couriersList, nil
}

func (r *CouriersPostgres) GetById(courierId int) (models.Courier, error) {
	var courierInfo models.Courier
	err := r.db.Model(&courierInfo).Where("id = ?", courierId).Select()
	if err != nil {
		return courierInfo, err
	}
	return courierInfo, nil
}

func (r *CouriersPostgres) GetCouriersSalaryRating(courierId int, start_date string, end_date string) (models.Courier, int, error) {
	courierRE, err := r.GetById(courierId)
	if err != nil {
		return models.Courier{}, 0, err
	}
	var count int
	var metaOrder models.Order
	err = r.db.Model(&metaOrder).ColumnExpr("count(id)").Where("complete_time is not null").Where("courier_id = ?", courierId).Where("complete_time BETWEEN ? AND ?",start_date, end_date).Select(&count)
    
	if count == 0 {
		return models.Courier{}, 0, nil
	}
	return courierRE, count, err
}