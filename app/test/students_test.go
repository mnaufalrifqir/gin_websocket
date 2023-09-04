package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"gin-framework/app/controllers"
	"gin-framework/app/db"
	"gin-framework/app/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type TestCase struct {
	name                   string
	path                   string
	expectedStatus         int
	expectedBodyStartsWith string
}

var (
	DB *gorm.DB
)

var config = db.Config{
	DB_Username: "root",
	DB_Password: "",
	DB_Host:     "localhost",
	DB_Port:     "3306",
	DB_Name:     "gin-framework",
}

func InitGin() *gin.Engine {
	DB = config.ConnectDB()
	db.InitialMigration()
	r := gin.Default()

	return r
}

func TestCreateStudent_Success(t *testing.T) {
	testcase := TestCase{
		name:                   "success",
		path:                   "/students",
		expectedStatus:         http.StatusCreated,
		expectedBodyStartsWith: `{"message":`,
	}

	InitGin()

	var studentInput models.StudentInput = models.StudentInput{
		Nim:     "12345",
		Name:    "John Doe",
		Jurusan: "Computer Science",
	}

	studentInputJSON, _ := json.Marshal(studentInput)

	req, _ := http.NewRequest(http.MethodPost, testcase.path, bytes.NewBuffer(studentInputJSON))
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req
	ctx.Request.URL.Path = testcase.path

	controllers.CreateStudent(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)

	body := rec.Body.String()
	fmt.Println(body)

	assert.True(t, strings.HasPrefix(body, testcase.expectedBodyStartsWith))

	t.Cleanup(func() {
		db.CleanSeeders(DB)
	})
}

func TestGetAllStudents_Success(t *testing.T) {
	testcase := TestCase{
		name:           "success",
		path:           "/students",
		expectedStatus: http.StatusOK,
	}

	InitGin()

	_, err := db.SeedStudent(DB)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req, _ := http.NewRequest(http.MethodGet, testcase.path, nil)
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Request = req

	ctx.Request.URL.Path = testcase.path

	controllers.ReadAllStudents(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)

	t.Cleanup(func() {
		db.CleanSeeders(DB)
	})
}

func TestGetStudentByID_Success(t *testing.T) {
	testcase := TestCase{
		name:           "success",
		path:           "/students",
		expectedStatus: http.StatusOK,
	}

	InitGin()

	student, err := db.SeedStudent(DB)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req, _ := http.NewRequest(http.MethodGet, testcase.path, nil)
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: strconv.Itoa(int(student.ID)),
		},
	}
	ctx.Request = req

	ctx.Request.URL.Path = testcase.path

	controllers.GetStudentByID(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)

	t.Cleanup(func() {
		db.CleanSeeders(DB)
	})
}

func TestUpdateStudentByID(t *testing.T) {
	testcase := TestCase{
		name:           "success",
		path:           "/students",
		expectedStatus: http.StatusOK,
	}

	InitGin()

	student, err := db.SeedStudent(DB)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	var studentInput models.StudentUpdate = models.StudentUpdate{
		Name: "John Doe Updated",
	}

	studentInputJSON, _ := json.Marshal(studentInput)

	req, _ := http.NewRequest(http.MethodPut, testcase.path, bytes.NewBuffer(studentInputJSON))
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: strconv.Itoa(int(student.ID)),
		},
	}
	ctx.Request = req

	ctx.Request.URL.Path = testcase.path

	controllers.UpdateStudentByID(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)

	t.Cleanup(func() {
		db.CleanSeeders(DB)
	})
}

func TestDeleteStudentByID(t *testing.T) {
	testcase := TestCase{
		name:           "success",
		path:           "/students",
		expectedStatus: http.StatusOK,
	}

	InitGin()

	student, err := db.SeedStudent(DB)
	if err != nil {
		t.Errorf("error: %v\n", err)
	}

	req, _ := http.NewRequest(http.MethodDelete, testcase.path, nil)
	rec := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(rec)
	ctx.Params = []gin.Param{
		{
			Key:   "id",
			Value: strconv.Itoa(int(student.ID)),
		},
	}

	ctx.Request = req

	ctx.Request.URL.Path = testcase.path

	controllers.DeleteStudentByID(ctx)

	assert.Equal(t, testcase.expectedStatus, rec.Code)

	t.Cleanup(func() {
		db.CleanSeeders(DB)
	})
}
