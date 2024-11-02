package presenter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iotassss/puzzdra-monster-rating/internal/apperrors"
	"github.com/iotassss/puzzdra-monster-rating/internal/usecase"
)

type GetMonsterRatingPresenter struct {
	ginctx *gin.Context
}

func NewGetMonsterRatingPresenter() *GetMonsterRatingPresenter {
	return &GetMonsterRatingPresenter{}
}

func (p *GetMonsterRatingPresenter) SetGinContext(ginctx *gin.Context) {
	p.ginctx = ginctx
}

func (p *GetMonsterRatingPresenter) Present(monsterRatingDTO usecase.MonsterRating) error {
	monsterRating := gin.H{
		"No":           monsterRatingDTO.No,
		"Name":         monsterRatingDTO.Name,
		"Game8Monster": monsterRatingDTO.Game8Monster,
	}

	p.ginctx.HTML(http.StatusOK, "monster_rating.html", monsterRating)

	return nil
}

func (p *GetMonsterRatingPresenter) PresentError(err error) error {
	var ErrValidation *apperrors.ErrValidation
	if errors.As(err, &ErrValidation) {
		p.ginctx.HTML(http.StatusBadRequest, "error.html", gin.H{"message": ErrValidation.Message})
		return nil
	}

	var ErrNotFound *apperrors.ErrNotFound
	if errors.As(err, &ErrNotFound) {
		p.ginctx.HTML(http.StatusNotFound, "error.html", gin.H{"message": ErrNotFound.Message})
		return nil
	}

	// その他のエラーの場合は、Internal Server Errorとして扱う
	p.ginctx.HTML(http.StatusInternalServerError, "error.html", gin.H{"message": "Internal Server Error"})
	return err
}
