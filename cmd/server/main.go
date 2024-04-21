package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	handlerAppointment "github.com/Eliezer134/final_EB3.git/cmd/server/handler/appointment"
	handlerDentist "github.com/Eliezer134/final_EB3.git/cmd/server/handler/dentist"
	handlerPatient "github.com/Eliezer134/final_EB3.git/cmd/server/handler/patient"
	handlerPing "github.com/Eliezer134/final_EB3.git/cmd/server/handler/ping"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Eliezer134/final_EB3.git/docs"
	"github.com/Eliezer134/final_EB3.git/internal/appointment"
	"github.com/Eliezer134/final_EB3.git/internal/dentist"
	"github.com/Eliezer134/final_EB3.git/internal/patient"
	"github.com/Eliezer134/final_EB3.git/pkg/middleware"
	"github.com/joho/godotenv"
)

// @title           Gestios de citas Clinica Odontológica
// @version         2.3.1
// @description     Reserva de citas para pacientes de una clinica odontologica.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	// Cargar las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//conectado a la base de datos
	db := connectDB()

	// Ping.
	controllerPing := handlerPing.NewControllerPing()

	// Products.
	//repository := dentist.NewMemoryRepository(db)
	repository := dentist.NewMySqlRepository(db)
	service := dentist.NewServiceDentist(repository)
	controllerDentist := handlerDentist.NewControllerDentists(service)

	// Patient.
	repositoryPatient := patient.NewMySqlRepository(db)
	servicePatient := patient.NewServicePatient(repositoryPatient)
	controllerPatient := handlerPatient.NewControllerPatients(servicePatient)

	// Appointment.
	repositoryAppointment := appointment.NewMySqlRepository(db)
	serviceAppointment := appointment.NewServiceAppointment(repositoryAppointment)
	controllerAppointment := handlerAppointment.NewControllerAppointments(serviceAppointment)

	engine := gin.Default()
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())
	//docs.SwaggerInfo.Host = os.Getenv("HOST")

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := engine.Group("/api/v1")
	{
		group.GET("/ping", controllerPing.HandlerPing())

		grupoDentist := group.Group("/dentist")
		{
			grupoDentist.POST("", middleware.Authenticate(), controllerDentist.HandlerCreate())
			grupoDentist.GET("", controllerDentist.HandlerGetAll())
			grupoDentist.GET("/:id", controllerDentist.HandlerGetByID())
			grupoDentist.PUT("/:id", middleware.Authenticate(), controllerDentist.HandlerUpdate())
			grupoDentist.PATCH("/:id", middleware.Authenticate(), controllerDentist.HandlerPatch())
			grupoDentist.DELETE("/:id", middleware.Authenticate(), controllerDentist.HandlerDelete())

		}

		grupoPatient := group.Group("/patient")
		{
			grupoPatient.POST("", middleware.Authenticate(), controllerPatient.HandlerCreate()) //llama a las fciones del controller
			grupoPatient.GET("", controllerPatient.HandlerGetAll())
			grupoPatient.GET("/:id", controllerPatient.HandlerGetByID())
			grupoPatient.PUT("/:id", middleware.Authenticate(), controllerPatient.HandlerUpdate())
			grupoPatient.PATCH("/:id", middleware.Authenticate(), controllerPatient.HandlerPatch())
			grupoPatient.DELETE("/:id", middleware.Authenticate(), controllerPatient.HandlerDelete())

		}

		grupoAppointment := group.Group("/appointment")
		{
			grupoAppointment.POST("", middleware.Authenticate(), controllerAppointment.HandlerCreate()) //llama a las fciones del controller
			grupoAppointment.GET("", controllerAppointment.HandlerGetAll())
			grupoAppointment.GET("/:id", controllerAppointment.HandlerGetByID())
			grupoAppointment.PUT("/:id", middleware.Authenticate(), controllerAppointment.HandlerUpdate())
			grupoAppointment.DELETE("/:id", middleware.Authenticate(), controllerAppointment.HandlerDelete())

		}

	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

//base de datos

func connectDB() *sql.DB {
	var dbUsername, dbPassword, dbHost, dbPort, dbName string

	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = "localhost"
	dbPort = "3306"
	dbName = os.Getenv("DB_NAME")

	// string de conexion
	// "username:password@tcp(host:puerto)/base_datos"
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}
