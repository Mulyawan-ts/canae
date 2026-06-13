package http

import (
	"net/http"
	"strconv"

	"backend/internal/domain"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	usecase domain.EmployeeUsecase
}

func NewEmployeeHandler(usecase domain.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{usecase}
}

func (h *EmployeeHandler) GetAll(c *gin.Context) {
	employees, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employees})
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	employee, err := h.usecase.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var employee domain.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.usecase.Create(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": employee})
}

func (h *EmployeeHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var employee domain.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee.ID = uint(id)
	if err := h.usecase.Update(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": employee})
}

func (h *EmployeeHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "employee deleted"})
}
