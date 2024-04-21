package dentist

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Eliezer134/final_EB3.git/internal/domain"
	"github.com/Eliezer134/final_EB3.git/internal/dentist"
	"github.com/Eliezer134/final_EB3.git/pkg/web"
)

type Controller struct {
	service dentist.Service
}

func NewControllerDentists(service dentist.Service) *Controller {
	return &Controller{service: service}
}

// Dentist godoc
// @Summary dentist example
// @Description Create a new dentist
// @Tags dentist
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var dentistRequest domain.Dentist

		err := ctx.Bind(&dentistRequest)

		if err != nil {

			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		dentist, err := c.service.Create(ctx, dentistRequest)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})*/
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, dentist)
		/*ctx.JSON(http.StatusOK, gin.H{
			"data": dentist,
		})*/

	}
}

// Dentist godoc
// @Summary dentist example
// @Description Get all dentists
// @Tags dentist
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /dentist [get]
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		newContext := addValueToContext(ctx)
		listDentists, err := c.service.GetAll(newContext)
		if err != nil {
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			/*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})*/
			return
		}
		web.Success(ctx, http.StatusOK, listDentists)
		/*ctx.JSON(http.StatusOK, gin.H{
			"data": listDentists,
		})*/
	}
}

// Dentist godoc
// @Summary dentist example
// @Description Get dentist by id
// @Tags dentist
// @Param id path int true "id del dentist"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist/:id [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		web.Success(ctx, http.StatusOK, id)

		/*dentist, err := c.service.GetByID(ctx, id)

		ctx.JSON(http.StatusOK, gin.H{
			"data": dentist,
		})*/
	}
}

// Dentist godoc
// @Summary dentist example
// @Description Update dentist by id
// @Tags dentist
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist/:id [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idParam := ctx.Param("id")
		
		id, err := strconv.Atoi(idParam)

		if err != nil {
			
			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
		}

		var dentistRequest domain.Dentist

		err = ctx.Bind(&dentistRequest)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		dentist, err := c.service.Update(ctx, dentistRequest, id)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})*/
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			
			return
		}

		/*ctx.JSON(http.StatusOK, gin.H{
			"data": dentist,
		})*/
		web.Success(ctx, http.StatusOK, dentist)
	}
}

// Dentist godoc
// @Summary dentist example
// @Description Patch dentist
// @Tags dentist
// @Param id path int true "id del dentist"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist/:id [patch]
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		idParam := ctx.Param("id")

		id, err := strconv.Atoi(idParam)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		var dentistRequest domain.Dentist

		err = ctx.Bind(&dentistRequest)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
			return
		}

		
		patchedDentist, err := c.service.Patch(ctx, dentistRequest, id)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})*/
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		/*ctx.JSON(http.StatusOK, gin.H{
			"data": patchedDentist,
		})*/
		web.Success(ctx, http.StatusOK, patchedDentist)
	}
}

// Dentist godoc
// @Summary dentist example
// @Description Delete dentist by id
// @Tags dentist
// @Param id path int true "id del dentist"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /dentist/:id [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// Recuperamos el id de la request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "id invalido")
			return
		}

		err = c.service.Delete(ctx, id)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})*/
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		/*ctx.JSON(http.StatusOK, gin.H{
			"message": "Dentista eliminado",
		})*/
		web.Success(ctx, http.StatusOK, "Dentist deleted")
	}
}

// addValueToContext ...a la espera de clase de middleware
func addValueToContext(ctx context.Context) context.Context {
	newContext := context.WithValue(ctx, "rol", "admin")
	return newContext
}
