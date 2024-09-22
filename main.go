package main

import (
	"clean-architecture-go-example/application_business_rules/usecase"
	"clean-architecture-go-example/interface_adaptors/handler"
	"clean-architecture-go-example/interface_adaptors/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	studentRepo := repository.NewInMemoryStudentRepository()
	studentUseCase := &usecase.StudentUseCase{StudentRepo: studentRepo}
	studentHandler := &handler.StudentHandler{StudentUseCase: studentUseCase}

	r := gin.Default()
	r.GET("/students/:id", studentHandler.GetStudent)
	r.POST("/students", studentHandler.CreateStudent)

	r.Run(":8080")
}
