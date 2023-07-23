package controllers

import (
	"net/http"
	"strconv"

	"github.com/EltIsma/YandexLavka/cmd/internal/dto"
	"github.com/EltIsma/YandexLavka/cmd/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCouriers(c *gin.Context) {
	var input []*dto.CourierDto
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	couriers := make([]*models.Courier, 0, len(input))
	for _, cour := range input {
		couriers = append(couriers, &models.Courier{
			Id:           cour.Id,
			Type:         string(cour.Type),
			Regions:      cour.Regions,
			WorkingHours: cour.WorkingHours,
		})
	}

	err := h.services.CouriersList.Create(couriers)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}

func (h *Handler) getAllCouriers(c *gin.Context) {
	var limit, offset int
	l := c.Query("limit")
	q := c.Query("offset")
	if len(l) > 0 {
		limit, _ = strconv.Atoi(l)
	} else {
		limit = 1
	}

	if len(q) > 0 {
		offset, _ = strconv.Atoi(q)
	} else {
		offset = 0
	}

	couriers_lists, err := h.services.CouriersList.GetAll(limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, couriers_lists)

}

func (h *Handler) getCourierById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	couriers_list, err := h.services.CouriersList.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, couriers_list)

}
func (h *Handler) getCourierRatingSalary(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	start := c.Query("start_date")
	end := c.Query("end_date")
	//var courierRating models.Courier

	earning, rating, err := h.services.CouriersList.GetCouriersSalaryRating(id, start, end)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
   var ratingAndEarning dto.CourierDtoEarningRating
   ratingAndEarning.Earning = earning
   ratingAndEarning.Ratings = rating
	c.JSON(http.StatusOK, ratingAndEarning)
}
