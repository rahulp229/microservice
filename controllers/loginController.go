package controllers

import (
	"microservice/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	FetchDataController(c *gin.Context)
}

type loginController struct {
	TestSvc service.TestService
}

type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func NewLoginAuthController(testsvc service.TestService) LoginController {
	return &loginController{
		TestSvc: testsvc,
	}
}

type Requestbody struct {
	Fsyms string `form:"fsyms"`
	Tsyms string `form:"tsyms"`
}

func (lc *loginController) FetchDataController(c *gin.Context) {
	var req Requestbody
	err := c.ShouldBindQuery(&req)
	//fmt.Println("fsyms : ", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if req.Fsyms == "" || req.Tsyms == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fsyms should not empty"})
		return
	}

	resp, err := lc.TestSvc.FetchData(req.Fsyms, req.Tsyms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, resp)
}
