package handler

import (
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/sales"
	"github.com/gin-gonic/gin"
)

type Sales struct {
	s sales.Service
}

func NewHandlerSales(s sales.Service) *Sales {
	return &Sales{s}
}

func (s *Sales) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices, err := s.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, invoices)
	}
}

func (s *Sales) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sale := domain.Sales{}
		err := ctx.ShouldBindJSON(&sale)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = s.s.Create(&sale)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": sale})
	}
}

func (c *Sales) BatchPost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sales := []domain.Sales{}
		err := ctx.ShouldBindJSON(&sales)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		for _, item := range sales {
			err = c.s.Create(&item)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}

		ctx.JSON(201, gin.H{"data": "sales created sucessffully"})
	}
}
