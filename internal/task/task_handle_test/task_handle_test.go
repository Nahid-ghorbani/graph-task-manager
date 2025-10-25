package task_handle_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nahid-ghorbani/graph-task-manager/internal/task"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock the test repository and it's method for testing
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(task *task.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Delete(task *task.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) Update(task *task.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) FindTask(task *task.Task, id int) error {
	args := m.Called(task, id)
	return args.Error(0)
}

func (m *MockTaskRepository) GetAll(tasks *[]task.Task) error {
	args := m.Called(tasks)
	return args.Error(0)
}

// seperate the repeated setup and define it as setupTest
func setupTest() (*gin.Engine, *MockTaskRepository, *httptest.ResponseRecorder) {
	mockRepo := new(MockTaskRepository)

	taskHandler := task.NewTaskHandler(mockRepo)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	taskHandler.RegisterRoutes(router)
	response := httptest.NewRecorder()

	return router, mockRepo, response
}

// testing the create task func
func TestCreateTask(t *testing.T) {
	router, mockRepo, response := setupTest()

	//define task for testing, also body needed for present json format (the crud required format)
	taskInstance := &task.Task{Title: "test1", Description: "test create task", Status: "Todo", Assignee: "nahid"}
	body, _ := json.Marshal(taskInstance)

	//set expected result
	mockRepo.On("Create", taskInstance).Return(nil)

	//create post request and set the json request
	request, _ := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)

	mockRepo.AssertExpectations(t)
}

// testing the delete task crud
func TestDeleteTask(t *testing.T) {
	router, mockRepo, response := setupTest()
	id := 0

	mockRepo.On("FindTask", mock.AnythingOfType("*task.Task"), id).Return(nil)
	mockRepo.On("Delete", mock.AnythingOfType("*task.Task")).Return(nil)

	request, _ := http.NewRequest(http.MethodDelete, "/tasks/0", nil)

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	mockRepo.AssertExpectations(t)
}

func TestGetAllTasks(t *testing.T) {
	router, mockRepo, response := setupTest()

	mockRepo.On("GetAll", mock.AnythingOfType("*[]task.Task")).Return(nil)

	request, _ := http.NewRequest(http.MethodGet, "/tasks", nil)

	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	mockRepo.AssertExpectations(t)
}

func TestGetTaskDetail(t *testing.T) {
	router, mockRepo, response := setupTest()
	id := 0

	mockRepo.On("FindTask", mock.AnythingOfType("*task.Task"), id).Return(nil)

	request, _ := http.NewRequest(http.MethodGet, "/tasks/0", nil)

	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	mockRepo.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	router, mockRepo, response := setupTest()

	id := 0
	taskInstance := &task.Task{Title: "test1", Description: "test create task", Status: "Todo", Assignee: "nahid"}

	mockRepo.On("FindTask", mock.AnythingOfType("*task.Task"), id).Return(nil)
	mockRepo.On("Update", taskInstance).Return(nil)

	body, _ := json.Marshal(taskInstance)
	request, _ := http.NewRequest(http.MethodPatch, "/tasks/0", bytes.NewBuffer(body))

	router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)

	mockRepo.AssertExpectations(t)
}
