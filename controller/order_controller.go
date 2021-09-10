package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"goder/models"
	"goder/mysql"
	"net/http"
	"strconv"
	"context"
)

func GetOrderByID(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	ctx := context.Background()
	order, err := models.Orders(
		qm.Where("id=?",i),
		).One(ctx, mysql.DB)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, order)
}

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

func UpdateOrder(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	ctx := context.Background()
	order, err := models.Orders(
		qm.Where("id=?",i),
	).One(ctx, mysql.DB)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(c.FormValue("creator_id") )
	fmt.Println(c.FormValue("contractor_id") )

	if c.FormValue("creator_id") != "" {
		creatorIDTmp,_ := strconv.Atoi(c.FormValue("creator_id"))
		creatorID := uint64(creatorIDTmp)
		order.CreatorID = null.NewUint64(creatorID,true)
		print("update creator_id")
	}

	if c.FormValue("contractor_id") != "" {
		contractorIDTmp,_  := strconv.Atoi(c.FormValue("contractor_id"))
		contractorID := uint64(contractorIDTmp)
		order.ContractorID = null.NewUint64(contractorID, true)
		print("update contractor_id")
	}
	order.Update(ctx, mysql.DB, boil.Infer())
	return c.JSON(http.StatusOK, order)
}
