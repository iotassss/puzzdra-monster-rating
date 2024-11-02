package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/puzzdra-monster-rating/internal/usecase"
	"gorm.io/gorm"
)

type GetMonsterRatingHandler struct {
	db      *gorm.DB
	usecase usecase.GetMonsterRatingUsecase
}

func NewGetMonsterRatingHandler(
	db *gorm.DB,
	usecase usecase.GetMonsterRatingUsecase,
) *GetMonsterRatingHandler {
	return &GetMonsterRatingHandler{
		db:      db,
		usecase: usecase,
	}
}

func (h *GetMonsterRatingHandler) Execute(c *gin.Context) {
	no := c.Param("no")
	noInt, err := strconv.Atoi(no)
	if err != nil {
		slog.Error("strconv.Atoi failed", slog.Any("error", err))
		c.HTML(http.StatusNotFound, "error.html", gin.H{"message": "Monster Not Found"})
		return
	}

	err = h.usecase.Execute(c, noInt)
	if err != nil {
		slog.Error("usecase.Execute failed", slog.Any("error", err))
		return
	}
}
