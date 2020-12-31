package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Choheeseok/react-go-practice/back/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CORSMiddlewareWrapper return CORSMiddleware with config
func CORSMiddlewareWrapper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := ctx.Request()
		dynamicCORSConfig := middleware.CORSConfig{
			AllowOrigins: []string{req.Header.Get("Origin"), "http://localhost:3000"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}
		CORSMiddleware := middleware.CORSWithConfig(dynamicCORSConfig)
		CORSHandler := CORSMiddleware(next)
		return CORSHandler(ctx)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Handler (echo, dbHandler) struct
type Handler struct {
	echo *echo.Echo
	db model.DBHandler
}

func (ah *Handler) index(c echo.Context) error {
	fmt.Printf("%v - %v\n", c.Request().Method, c.Request().RequestURI)

	toDos := ah.db.GetToDos()

	return c.JSON(http.StatusOK, toDos)
}

func (ah *Handler) addToDo(c echo.Context) error {
	fmt.Printf("%v - %v\n", c.Request().Method, c.Request().RequestURI)

	toDo := &model.ToDo{}
	err := c.Bind(toDo)
	checkErr(err)

	toDo, err = ah.db.AddToDo(toDo)
	
	return c.JSON(http.StatusCreated, toDo)
}

func (ah *Handler) deleteToDo(c echo.Context) error {
	fmt.Printf("%v - %v\n", c.Request().Method, c.Request().RequestURI)

	id, err := strconv.Atoi(c.Param("id"))
	checkErr(err)

	ah.db.DeleteToDo(id)

	return c.String(http.StatusOK, fmt.Sprintf("Delete Success : %v", id))
}

func (ah *Handler) completeToDo(c echo.Context) error {
	fmt.Printf("%v - %v\n", c.Request().Method, c.Request().RequestURI)

	id, err := strconv.Atoi(c.Param("id"))
	checkErr(err)

	ah.db.CompleteToDo(id)

	return c.String(http.StatusOK, fmt.Sprintf("Set Complete : %v", id))
}

func (ah *Handler) getDetail(c echo.Context) error {
	fmt.Printf("%v - %v\n", c.Request().Method, c.Request().RequestURI)

	id, err := strconv.Atoi(c.Param("id"))
	checkErr(err)

	toDo, err := ah.db.GetDetail(id)
	checkErr(err)

	return c.JSON(http.StatusOK, toDo)
}

// Start server
func Start(port string) {
	appHandler := &Handler{
		echo: echo.New(),
		db: model.NewDBHandler(),
	}
	defer appHandler.db.Close()
	defer appHandler.echo.Close()
	appHandler.echo.Use(CORSMiddlewareWrapper)
	appHandler.echo.GET("/", appHandler.index)
	appHandler.echo.POST("/", appHandler.addToDo)
	appHandler.echo.DELETE("/:id", appHandler.deleteToDo)
	appHandler.echo.PUT("/:id", appHandler.completeToDo)
	appHandler.echo.GET("/detail/:id", appHandler.getDetail)

	appHandler.echo.Start(port)
}