package handler

import (
	"log"
	"project/internal/auth"
	"project/internal/middleware"
	service "project/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupApi(a *auth.Auth, svc service.UserService) *gin.Engine {
	r := gin.New()

	m, err := middleware.NewMiddleware(a)
	if err != nil {
		log.Panic("middlewares not setup")
		return nil
	}
	h, err := Newhandler(svc)
	if err != nil {
		log.Panic("middlewares not setup")
		return nil
	}

	r.Use(m.Log(), gin.Recovery())

	r.GET("/check")
	r.POST("/signup", h.SignUp)
	r.POST("/signin", h.Signin)
	r.POST("/add", m.Authenticate(h.AddCompany))
	r.GET("/view/all", m.Authenticate(h.ViewAllCompanies))
	r.GET("/view/:id", m.Authenticate(h.ViewCompany))
	r.GET("/job/view/:id", m.Authenticate(h.ViewJob))
	r.POST("/add/:cid", m.Authenticate(h.AddJobs))
	r.GET("/view/all", m.Authenticate(h.ViewAllJobs))
	r.GET("/view/:id", m.Authenticate(h.ViewJobByID))

	return r

}

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "ok",
	})
}
