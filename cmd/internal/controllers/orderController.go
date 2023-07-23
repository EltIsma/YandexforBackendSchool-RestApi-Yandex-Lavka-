package controllers

import (
	"net/http"
	"strconv"

	"github.com/EltIsma/YandexLavka/cmd/internal/dto"
	"github.com/EltIsma/YandexLavka/cmd/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createOrders(c *gin.Context) {
	var input []*dto.OrderDto
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	orders := make([]*models.Order, 0, len(input))
	for _, ord := range input {
		orders = append(orders, &models.Order{
			Id:           int64(ord.Id),
			Weight:         ord.Weight,
			Region:         ord.Region,
			DeliveryHours:  ord.DeliveryHours,
			Cost:           ord.Cost,
		})
	}

	err := h.services.OrdersList.Create(orders)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

}

func (h *Handler) getAllOrders(c *gin.Context) {
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

	order_lists, err := h.services.OrdersList.GetAll(limit, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}
	var orderDTO_lists []*dto.OrderDto
	for _, ord := range order_lists {
		orderDTO_lists = append(orderDTO_lists, &dto.OrderDto{
			Id:             int(ord.Id),
			Weight:         ord.Weight,
			Region:         ord.Region,
			DeliveryHours:  ord.DeliveryHours,
			Cost:           ord.Cost,
		})
	}
	c.JSON(http.StatusOK, orderDTO_lists)

}

func (h *Handler) getOrderById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	orders_list, err := h.services.OrdersList.GetById(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}
	var orderDTO_lists dto.OrderDto
	
	orderDTO_lists = dto.OrderDto{
			Id:             int(orders_list.Id),
			Weight:         orders_list.Weight,
			Region:         orders_list.Region,
			DeliveryHours:  orders_list.DeliveryHours,
			Cost:           orders_list.Cost,
	}
	

	c.JSON(http.StatusOK, orderDTO_lists)

}


func (h *Handler) ordersComplete(c *gin.Context) {
    var input dto.OrderCompleteDto
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return 
	}
	id, err := h.services.OrdersList.Update(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	
}