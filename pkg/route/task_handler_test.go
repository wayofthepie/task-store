package route_test

import (
	"bytes"
	"errors"
	"github.com/alicebob/miniredis"
	"github.com/execd/task-store/pkg/model"
	"github.com/execd/task-store/pkg/route"
	"github.com/execd/task-store/pkg/store"
	"github.com/execd/task-store/pkg/task"
	"github.com/execd/task-store/pkg/util"
	. "github.com/onsi/ginkgo"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

var context = GinkgoT()

var _ = Describe("task handler", func() {
	Describe("create task", func() {
		var taskStore *task.StoreImpl
		var directRedis *miniredis.Miniredis
		var handler *route.TaskHandlerImpl

		BeforeEach(func() {
			s, err := miniredis.Run()
			if err != nil {
				panic(err)
			}
			directRedis = s
			redis := store.NewClient(s.Addr())
			uuidGen := util.NewUUIDGenImpl()
			taskStore = task.NewStoreImpl(redis, uuidGen)
			config := &model.Config{
				Manager: model.ManagerInfo{
					ExecutionQueueSize: 10,
					TaskQueueSize:      10,
				},
			}
			handler = route.NewTaskHandlerImpl(taskStore, config)
		})

		It("should return an error if reading body fails", func() {
			// Arrange
			r := bytes.NewReader(nil)
			r.Seek(10, 0)
			req, _ := http.NewRequest("POST", "/handle", errReader(0))
			writer := httptest.NewRecorder()

			// Act
			handler.CreateTask(writer, req)

			// Assert
			assert.Equal(context, 500, writer.Code)
			assert.Equal(context, "error\n", writer.Body.String())
		})

		It("should return error if request body does not contain a task", func() {
			// Arrange
			req, _ := http.NewRequest("POST", "/handle", bytes.NewReader([]byte("")))
			writer := httptest.NewRecorder()

			// Act
			handler.CreateTask(writer, req)

			// Assert
			assert.Equal(context, 500, writer.Code)
			assert.NotNil(context, writer.Body.String())
		})

		It("should return error if storing task fails", func() {
			// Arrange
			directRedis.Close()
			taskString := `{"name": "test", "image": "alpine", "init": "init.sh"}`
			req, _ := http.NewRequest("POST", "/handle", bytes.NewReader([]byte(taskString)))
			writer := httptest.NewRecorder()

			// Act
			handler.CreateTask(writer, req)

			// Assert
			assert.Equal(context, 500, writer.Code)
			assert.NotNil(context, writer.Body.String())
		})

		It("should return task id if task creation succeeds", func() {
			// Arrange
			taskString := `{"name": "test", "image": "alpine", "init": "init.sh"}`
			req, _ := http.NewRequest("POST", "/handle", bytes.NewReader([]byte(taskString)))
			writer := httptest.NewRecorder()

			// Act
			handler.CreateTask(writer, req)

			// Assert
			assert.Equal(context, 201, writer.Code)
			id, err := uuid.FromString(writer.Body.String())
			assert.Nil(context, err)
			assert.NotNil(context, id)
		})

		It("should return error of task queue size is greater than max task queue size", func() {
			// Arrange
			config := &model.Config{
				Manager: model.ManagerInfo{
					ExecutionQueueSize: 10,
					TaskQueueSize:      0,
				},
			}
			handler = route.NewTaskHandlerImpl(taskStore, config)
			givenID := uuid.Must(uuid.NewV4())
			taskStore.PushTask(&givenID)
			taskString := `{"name": "test", "image": "alpine", "init": "init.sh"}`
			req, _ := http.NewRequest("POST", "/handle", bytes.NewReader([]byte(taskString)))
			writer := httptest.NewRecorder()

			// Act
			handler.CreateTask(writer, req)

			// Assert
			assert.Equal(context, 500, writer.Code)
			assert.Equal(context, writer.Body.String(), "Failed to create task, task queue has reached its limit!\n")
		})
	})
})

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error")
}
