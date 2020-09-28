package handlers

import (
  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  "github.com/mauriciobergallo/school-students-api/pkg/adding"
  "github.com/mauriciobergallo/school-students-api/pkg/deleting"
  "github.com/mauriciobergallo/school-students-api/pkg/listing"
  "github.com/mauriciobergallo/school-students-api/pkg/updating"
)

func NewRestService(as adding.Service, ls listing.Service, ds deleting.Service, us updating.Service) {
  r := gin.Default()
  rest := &restService{addingService: as, listingService: ls, deletingService: ds, updatingService: us}

  r.GET("/api/health", rest.getHealth)
  r.GET("/api/students", rest.getStudents)
  r.GET("/api/students/:id", rest.getStudentById)
  r.POST("/api/students", rest.postStudent)
  r.DELETE("/api/students/:id", rest.deleteStudent)
  r.PUT("/api/students/:id", rest.updateStudent)

  r.Run()
}

type restService struct {
  addingService   adding.Service
  listingService  listing.Service
  deletingService deleting.Service
  updatingService updating.Service
}

func (rs *restService) getHealth(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "I'm Alive",
  })
}

func (rs *restService) getStudents(c *gin.Context) {
  sl, err := rs.listingService.ListStudents()
  if err != nil {
    c.JSON(500, gin.H{
      "status":  "InternalServerError",
      "message": err.Error(),
    })

    return
  }

  c.JSON(200, gin.H{
    "status":  "OK",
    "message": sl,
  })
}

func (rs *restService) getStudentById(c *gin.Context) {
  id, err := uuid.Parse(c.Param("id"))
  if err != nil {
    c.JSON(400, gin.H{
      "status":  "BadRequest",
      "message": err.Error(),
    })

    return
  }

  s, err := rs.listingService.ObtainStudentById(id)
  if err != nil {
    c.JSON(400, gin.H{
      "status":  "BadRequest",
      "message": err.Error(),
    })

    return
  }

  c.JSON(200, gin.H{
    "status":  "OK",
    "message": s,
  })
}

func (rs *restService) postStudent(c *gin.Context) {
  var st adding.Student
  c.BindJSON(&st)

  s, err := rs.addingService.AddStudent(st)

  if err != nil {
    c.JSON(400, gin.H{
      "status":  "BadRequest",
      "message": err.Error(),
    })

    return
  }

  c.JSON(200, gin.H{
    "status": "OK",
    "data":   s,
  })
}

func (rs *restService) deleteStudent(c *gin.Context) {
  id, err := uuid.Parse(c.Param("id"))
  if err != nil {
    c.JSON(400, gin.H{
      "status":  "BadRequest",
      "message": err.Error(),
    })

    return
  }

  st := &deleting.Student{ID: id}
  err = rs.deletingService.RemoveStudent(*st)
  if err != nil {
    c.JSON(400, gin.H{
      "status":  "BadRequest",
      "message": err.Error(),
    })

    return
  }

  c.JSON(200, gin.H{
    "status": "OK",
    "data":   nil,
  })
}

func (rs *restService) updateStudent(c *gin.Context) {
  _, err := uuid.Parse(c.Param("id"))
  if err != nil {
    c.JSON(400, gin.H{
      "status":  "BadRequest",
      "message": err.Error(),
    })

    return
  }

  var st updating.Student
  c.BindJSON(&st)

  s, err := rs.updatingService.ReplaceStudent(st)

  if err != nil {
    c.JSON(400, gin.H{
      "status":  "BadRequest",
      "message": err.Error(),
    })

    return
  }

  c.JSON(200, gin.H{
    "status": "OK",
    "data":   s,
  })
}
