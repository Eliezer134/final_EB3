package patient

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Eliezer134/final_EB3.git/internal/patient"
	"github.com/Eliezer134/final_EB3.git/internal/domain"
	"github.com/Eliezer134/final_EB3.git/pkg/web"
)

// crea una estruct de controlador que va a tener inyectado el servicio de dentist
type Controller struct {
    service patient.Service
}

// recibe un servicio y devuelve un puntero de esa estructura
func NewControllerPatients(service patient.Service) *Controller {
    return &Controller{service: service} //dev ese controlador donde le inyectamos el servicio que vamos a crear en el main
}

// Patient godoc
// @Summary patient example
// @Description Create a new patient
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
    return func(ctx *gin.Context) {

        var patientRequest domain.Patient

        err := ctx.Bind(&patientRequest)

        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
                "message": "bad request",
                "error":   err,
            })
            return
        }

        patient, err := c.service.Create(ctx, patientRequest)
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                "message": "Internal server error",
            })
            return
        }

        ctx.JSON(http.StatusOK, gin.H{
            "data": patient,
        })

    }
}

// Patient godoc
// @Summary patient example
// @Description Get all patients
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 500 {object} web.errorResponse
// @Router /patient [get]
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
    return func(ctx *gin.Context) {

        newContext := addValueToContext(ctx)
        listPatients, err := c.service.GetAll(newContext) 
        if err != nil {                                   
            ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ 
                "message": "Internal server error", 
            })
            return
        }

        web.Success(ctx, http.StatusOK, listPatients)
        /*ctx.JSON(http.StatusOK, gin.H{ 
            "data": listPatients,
        })*/
    }
}

// Patient godoc
// @Summary patient example
// @Description Get patient by id
// @Tags patient
// @Param id path int true "id del patient"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient/:id [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
    return func(ctx *gin.Context) {

        
        idParam := ctx.Param("id")

        
        id, err := strconv.Atoi(idParam)
        
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                "message": "bad request",
                "error": err,
            })
            return
        }
        
        patient, err := c.service.GetByID(ctx, id)
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                "message": "Internal server error",
            })
            return
        }

        web.Success(ctx, http.StatusOK, patient)
        /*ctx.JSON(http.StatusOK, gin.H{ 
            "data": patient,
        })*/
    }
}

// Patient godoc
// @Summary patient example
// @Description Update patient by id
// @Tags patient
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient/:id [put]
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
            return
        }

        
        var patientRequest domain.Patient 

        err = ctx.Bind(&patientRequest) 
        
        if err != nil { 
            /*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
                "message": "bad request",
                "error":   err,
            })*/
            web.Error(ctx, http.StatusBadRequest, "%s", "bad request")
            return
        }

        
        patient, err := c.service.Update(ctx, patientRequest, id)
        if err != nil {
            /*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                "message": "Internal server error",
            })*/
            web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
            return
        }
        web.Success(ctx, http.StatusOK, patient)
        /*ctx.JSON(http.StatusOK, gin.H{
            "data": patient,
        })*/
    }
}

// Patient godoc
// @Summary patient example
// @Description Patch patient
// @Tags producto
// @Param id path int true "id del patient"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient/:id [patch]
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

        var patientRequest domain.Patient

        err = ctx.Bind(&patientRequest)
        if err != nil {
            /*ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
                "message": "bad request",
                "error":   err,
            })*/
            web.Error(ctx, http.StatusBadRequest, "%s", "bad request binding")
            return
        }

        
        patchedPatient, err := c.service.Patch(ctx, patientRequest, id)
        if err != nil {
            /*ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
                "message": "Internal server error",
            })*/
            web.Error(ctx, http.StatusInternalServerError, "%s", "internal server error")
            return
        }

        /*ctx.JSON(http.StatusOK, gin.H{
            "data": patchedPatient,
        })*/
        web.Success(ctx, http.StatusOK, patchedPatient)
    }
}

// Patient godoc
// @Summary patient example
// @Description Delete patient by id
// @Tags patient
// @Param id path int true "id del patient"
// @Accept json
// @Produce json
// @Success 200 {object} web.response
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Router /patient/:id [delete]
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
            "message": "Paciente eliminado",
        })*/
        web.Success(ctx, http.StatusOK, "Patient deleted")
    }
}


// addValueToContext ...
func addValueToContext(ctx context.Context) context.Context {
    newContext := context.WithValue(ctx, "rol", "admin")
    return newContext
}