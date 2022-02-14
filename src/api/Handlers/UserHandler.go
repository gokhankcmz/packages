package Handlers

import (
	"Packages/src/api/Helpers/Mapper"
	"Packages/src/api/Helpers/Validator"
	"Packages/src/api/Type/DtoTypes"
	"Packages/src/api/Type/EntityTypes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) GetMany(ctx echo.Context) error {
	response := h.Repository.GetMany(h.Filter.GetFindOptions(ctx.QueryParams(), ctx.Request().Header), h.Filter.GetFilter(ctx.QueryParams(), ctx.Request().Header))
	return ctx.JSON(http.StatusOK, response)
}

func (h Handler) Create(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	var createDTO DtoTypes.CreateUserDto
	_ = json.NewDecoder(ctx.Request().Body).Decode(&createDTO)
	Validator.ValidateModelOrPanic(createDTO)
	UserEntity := Mapper.GetUserFromCreateUserDto(&createDTO)
	response := h.Repository.Create(UserEntity)
	return ctx.JSON(http.StatusCreated, response)
}

func (h Handler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	response := h.Repository.Delete(id)
	return ctx.JSON(http.StatusOK, response)

}

func (h Handler) GetSingle(ctx echo.Context) error {
	/*rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(2000) * time.Millisecond)*/
	id := ctx.Param("id")
	UserEntity := EntityTypes.User{}
	h.Repository.GetSingle(id, &UserEntity)
	return ctx.JSON(http.StatusOK, UserEntity)
}

func (h Handler) Update(ctx echo.Context) error {
	defer ctx.Request().Body.Close()
	updateDto := &DtoTypes.UpdateUserDto{}
	_ = json.NewDecoder(ctx.Request().Body).Decode(updateDto)
	Validator.ValidateModelOrPanic(updateDto)

	id := ctx.Param("id")
	UserEntity := &EntityTypes.User{}
	h.Repository.GetSingle(id, &UserEntity)
	Mapper.GetUserFromUpdateUserDto(updateDto, UserEntity)
	response := h.Repository.Update(UserEntity)
	return ctx.JSON(http.StatusOK, response)

}
