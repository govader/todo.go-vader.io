package handler

import (
	"net/http"

	todo "github.com/govader/todo.go-vader.io/api"
	echo "github.com/labstack/echo/v4"
)

const pageSize = 20

type (
	TaskHandler struct {
		Store todo.Store
	}
	addParam struct {
		Content string `json:"content" form:"content"`
	}
	listParam struct {
		Page int `param:"page" validate:"required"`
	}
	result struct {
		TotalPages int         `json:"total_pages"`
		Tasks      []todo.Task `json:"tasks"`
	}
)

func (h *TaskHandler) Add(c echo.Context) error {
	var param addParam
	if err := c.Bind(&param); err != nil {
		return err
	}

	task, err := h.Store.Create(param.Content)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) Get(c echo.Context) error {
	task, err := h.Store.Get(c.Param("id"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) ListAll(c echo.Context) error {
	return listTasks(h.Store.All())(c)
}

func (h *TaskHandler) ListDone(c echo.Context) error {
	return listTasks(h.Store.Done())(c)
}

func (h *TaskHandler) ListNotDone(c echo.Context) error {
	return listTasks(h.Store.NotDone())(c)
}

func (h *TaskHandler) Update(c echo.Context) error {
	task, err := h.Store.MarkAsDone(c.Param("id"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func listTasks(querier todo.TaskQuerier) func(c echo.Context) error {
	return func(c echo.Context) error {
		var param listParam
		if err := c.Bind(param); err != nil {
			return err
		}

		if err := c.Validate(param); err != nil {
			return err
		}

		count, err := querier.Count()
		if err != nil {
			return err
		}
		list, err := querier.List(param.Page*pageSize, pageSize)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, &result{
			TotalPages: count/pageSize + 1,
			Tasks:      list,
		})
	}
}
