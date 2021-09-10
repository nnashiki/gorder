package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"goder/models"
	"goder/mysql"
	"net/http"
	"strconv"
	"context"
)

func CreateOrder(c echo.Context) error {
	creatorIDTmp,_ := strconv.Atoi(c.FormValue("creator_id"))
	creatorID := uint64(creatorIDTmp)
	contractorIDTmp,_  := strconv.Atoi(c.FormValue("contractor_id"))
	contractorID := uint64(contractorIDTmp)

	ctx := context.Background()
	order := models.Order{
		CreatorID: null.NewUint64(creatorID,true),
		ContractorID: null.NewUint64(contractorID,true),
	}
	err := order.Insert(ctx, mysql.DB, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, order)
}
