package appointment

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Eliezer134/final_EB3.git/internal/appointment"
	"github.com/Eliezer134/final_EB3.git/internal/domain"
	"github.com/Eliezer134/final_EB3.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	service appointment.Service
}

// recibe un servicio y devuelve un puntero de esa estructura
func NewControllerAppointments(service appointment.Service) *Controller {
	return &Controller{service: service} //dev ese controlador donde le inyectamos el servicio que vamos a crear en el main
}

// Appointment godoc
// @Summary appointment example
// @Description Create a new appointment
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var appointmentRequest domain.Appointment

		err := ctx.Bind(&appointmentRequest)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}

		appointment, err := c.service.Create(ctx, appointmentRequest)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": appointment,
		})

	}
}

// Appointment godoc
// @Summary appointment example
// @Description Get all appointments
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /appointment [get]
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		newContext := addValueToContext(ctx)
		listAppointments, err := c.service.GetAll(newContext) 
		if err != nil {                                       
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ //si hay error le decimos que aborte con un json
				"message": "Internal server error", 
			})
			return
		}
		web.Success(ctx, http.StatusOK, listAppointments)
		/*ctx.JSON(http.StatusOK, gin.H{ 
			"data": listAppointments,
		})*/
	}
}

// Appointment godoc
// @Summary appointment example
// @Description Get appointment by id
// @Tags appointment
// @Param id path int true "id del appointment"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		
		idParam := ctx.Param("id")

		

		id, err := strconv.Atoi(idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
			return
		}
		
		web.Success(ctx, http.StatusOK, id)
		/*appointment, err := c.service.GetByID(ctx, id)

		ctx.JSON(http.StatusOK, gin.H{ 
			"data": appointment,
		})*/
	}
}

// Appointment godoc
// @Summary appointment example
// @Description Update appointment by id
// @Tags appointment
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [put]
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

		var dentistAppointment domain.Appointment 

		err = ctx.Bind(&dentistAppointment) 
		
		if err != nil { 
			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		// Llamamos al servicio

		appointment, err := c.service.Update(ctx, dentistAppointment, id)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})*/
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		web.Success(ctx, http.StatusOK, appointment)
		/*ctx.JSON(http.StatusOK, gin.H{
			"data": appointment,
		})*/
	}
}

// Appointment godoc
// @Summary appointment example
// @Description Patch appointment
// @Tags appointment
// @Param id path int true "id del appointment"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [patch]
func (c *Controller) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Recuperamos el turno según la fecha
		dateOnly := ctx.Param("date")

		// Convertir la fecha a entero
		date, err := strconv.Atoi(dateOnly)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "fecha invalida")
			return
		}

		// Binding de la estructura desde el cuerpo de la solicitud
		var requestAppointment domain.RequestAppointmentDateOnly
		if err := ctx.BindJSON(&requestAppointment); err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Bad request",
				"error":   err.Error(),
			})*/
			web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
			return
		}

		// Llamamos al servicio con la fecha y la solicitud
		appointment, err := c.service.Patch(ctx, requestAppointment, date)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})*/
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		// Respuesta exitosa
		/*ctx.JSON(http.StatusOK, gin.H{
			"data": appointment,
		})*/
		web.Success(ctx, http.StatusOK, appointment)
	}
}

// Appointment godoc
// @Summary appointment example
// @Description Delete appointment by id
// @Tags appointment
// @Param id path int true "id del appointment"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /appointment/:id [delete]
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

		// Llamamos al servicio
		err = c.service.Delete(ctx, id)
		if err != nil {
			/*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})*/
			web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
			return
		}

		/*ctx.JSON(http.StatusOK, gin.H{
			"message": "Turno eliminado",
		})*/
		web.Success(ctx, http.StatusOK, "Deleted Appointment")
	}
}

// addValueToContext ...a la espera de clase de middleware
func addValueToContext(ctx context.Context) context.Context {
	newContext := context.WithValue(ctx, "rol", "admin")
	return newContext
}