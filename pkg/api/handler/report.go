package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gokul.go/pkg/usecase/usecaseInterfaces"
)

type SalesHandler struct {
	salesUseCase usecaseInterfaces.SalesUseCase
}

func NewSalesHandler(salesUsecase usecaseInterfaces.SalesUseCase) *SalesHandler {
	return &SalesHandler{salesUseCase: salesUsecase}
}
func (h *SalesHandler) GetMonthlySalesReport(c *gin.Context) {
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid year"})
		return
	}
	month, err := strconv.Atoi(c.Query("month"))
	if err != nil || month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month"})
		return
	}

	sales, err := h.salesUseCase.GetMonthlySalesReport(year, time.Month(month))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sales)
}
