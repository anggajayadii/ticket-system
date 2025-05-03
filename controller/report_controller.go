package controller

import (
	"net/http"
	"strconv"
	"ticket-system/service"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	service service.ReportService
}

func NewReportController(service service.ReportService) *ReportController {
	return &ReportController{service: service}
}

func (c *ReportController) EventReport(ctx *gin.Context) {
	eventID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	report, err := c.service.GenerateEventReport(uint(eventID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, report)
}

func (c *ReportController) SummaryReport(ctx *gin.Context) {
	period := ctx.DefaultQuery("period", "month")
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")

	summary, err := c.service.GenerateRevenueSummary(period, startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, summary)
}
