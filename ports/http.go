package ports

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/egigiffari/nusa-test/app"
	"github.com/egigiffari/nusa-test/app/schedule"
	"github.com/gin-gonic/gin"
)

type HttpHandlers struct {
	app app.Application
}

func NewHttpHandlers(router *gin.RouterGroup, app app.Application) {
	h := HttpHandlers{app: app}

	router.GET("/schedules", h.Schedules)
	router.GET("/export-schedules", h.ExportCSV)
	router.GET("/check-schedule", h.CheckUserSchedule)
}

func (h HttpHandlers) Schedules(ctx *gin.Context) {
	rangeDate, err := h.getRangeDate(ctx)
	if err != nil {
		h.responseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	userUUID, ok := ctx.GetQuery("user_id")
	if ok {

		userSchedules, err := h.app.SingleUserSchedules.Handle(ctx, userUUID, rangeDate)
		if err != nil {
			h.responseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		h.responseOk(ctx, []schedule.UserSchedule{
			*userSchedules,
		})
		return
	}

	allUserSchedules := h.app.AllUserSchedules.Handle(ctx, rangeDate)
	h.responseOk(ctx, allUserSchedules)
}

func (h HttpHandlers) CheckUserSchedule(ctx *gin.Context) {
	userUUID, ok := ctx.GetQuery("user_id")
	if !ok {
		h.responseError(ctx, "Missing user_id or parameter.", http.StatusBadRequest)
		return
	}

	dateStr, ok := ctx.GetQuery("date")
	if !ok {
		h.responseError(ctx, "Missing date or parameter.", http.StatusBadRequest)
		return
	}

	layout := "2006-01-02"
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		h.responseError(ctx, err.Error(), http.StatusBadRequest)
	}

	scheduleStatus, err := h.app.CheckUserSchedule.Handle(ctx, userUUID, date)
	if err != nil {
		h.responseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	h.responseOk(ctx, gin.H{
		"id":    scheduleStatus.UserUUID,
		"name":  scheduleStatus.UserName,
		"date":  scheduleStatus.Date,
		"shift": scheduleStatus.Cycle,
	})
}

func (h HttpHandlers) ExportCSV(ctx *gin.Context) {
	rangeDate, err := h.getRangeDate(ctx)
	if err != nil {
		h.responseError(ctx, err.Error(), http.StatusBadRequest)
		return
	}

	userUUID, ok := ctx.GetQuery("user_id")
	if ok {
		fileName := fmt.Sprintf("jadwal_shift_%s_%s_to_%s.csv",
			userUUID,
			ctx.Query("start_date"),
			ctx.Query("end_date"))

		ctx.Writer.Header().Set("Content-Type", "text/csv")
		ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
		err := h.app.GenerateCSVSingleUserSchedules.Handle(ctx, userUUID, rangeDate, ctx.Writer)
		if err != nil {
			h.responseError(ctx, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	fileName := fmt.Sprintf("jadwal_shift_all_%s_to_%s.csv",
		ctx.Query("start_date"),
		ctx.Query("end_date"))

	ctx.Writer.Header().Set("Content-Type", "text/csv")
	ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	err = h.app.GenerateCSVAllUserSchedules.Handle(ctx, rangeDate, ctx.Writer)
	if err != nil {
		h.responseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h HttpHandlers) responseOk(ctx *gin.Context, data any) {
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":    "success",
			"data":      data,
			"http_code": http.StatusOK,
		},
	)
}

func (h HttpHandlers) responseError(ctx *gin.Context, message string, code int) {
	ctx.JSON(
		code,
		gin.H{
			"status":    "error",
			"message":   message,
			"http_code": code,
		},
	)
}

func (h HttpHandlers) getRangeDate(ctx *gin.Context) (schedule.RangeDates, error) {
	startDateStr, ok := ctx.GetQuery("start_date")
	if !ok {
		return schedule.RangeDates{}, errors.New("Missing start_date parameter.")
	}

	endDateStr, ok := ctx.GetQuery("end_date")
	if !ok {
		return schedule.RangeDates{}, errors.New("Missing end_date parameter.")
	}

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		return schedule.RangeDates{}, err
	}

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		return schedule.RangeDates{}, err
	}

	return schedule.RangeDates{
		From: startDate,
		To:   endDate,
	}, nil
}
