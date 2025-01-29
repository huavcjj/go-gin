package main

import (
	"encoding/json"
	"go-gin/infra"
	"go-gin/models"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalln("Error loading .env.test file")
	}

	code := m.Run()

	os.Exit(code)
}

func setupTestData(db *gorm.DB) {

	items := []models.Item{
		{Name: "Test 1", Price: 1000, Description: "Test 1 Description", SoldOut: false, UserID: 1},
		{Name: "Test 2", Price: 2000, Description: "Test 2 Description", SoldOut: true, UserID: 1},
		{Name: "Test 3", Price: 3000, Description: "Test 3 Description", SoldOut: false, UserID: 2},
	}

	users := []models.User{
		{Email: "test1@example.com", Password: "test1password"},
		{Email: "test2@example.com", Password: "test2password"},
	}

	for _, user := range users {
		db.Create(&user)
	}

	for _, item := range items {
		db.Create(&item)
	}
}

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&models.User{}, &models.Item{})

	setupTestData(db)
	router := setupRouter(db)
	return router
}

func TestFindAll(t *testing.T) {
	router := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/items", nil)

	router.ServeHTTP(w, req)

	var res []models.Item
	json.Unmarshal(w.Body.Bytes(), &res)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res))
}
