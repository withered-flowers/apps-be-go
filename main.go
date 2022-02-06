package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// Data Structure
type Batch struct {
	Name      string    `json:"batchName"`
	Code      string    `json:"batchCode"`
	Type      string    `json:"batchType"`
	ImageUrl  string    `json:"imageUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Phase struct {
	Number    int       `json:"phaseNumber"`
	BatchId   int       `json:"batchId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Student struct {
	FirstName string    `json:"studentFirstName"`
	LastName  string    `json:"studentLastName"`
	Email     string    `json:"studentEmail"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Course struct {
	PhaseId   int       `json:"phaseId"`
	StudentId int       `json:"studentId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func main() {
	r := gin.Default()

	// Root Endpoint
	r.GET("/", getRootHandler)

	// Batch Endpoint
	r.GET("/batch", getBatchRootHandler)
	r.POST("/batch", postBatchRootHandler)

	r.Run(":8080")
}

func getRootHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "Hello World",
	})
}

func getBatchRootHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"statusCode": 200,
		"message":    "Not implemented yet",
	})
}

func postBatchRootHandler(ctx *gin.Context) {
	batchName := ctx.PostForm("batch-name")
	batchCode := ctx.PostForm("batch-code")
	batchType := ctx.PostForm("batch-type")
	batchImageUrl := ctx.PostForm("batch-image-url")

	batch := &Batch{
		Name:      batchName,
		Code:      batchCode,
		Type:      batchType,
		ImageUrl:  batchImageUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	ctx.JSON(201, gin.H{
		"statusCode": 201,
		"batch":      batch,
	})
}
