package task_test

import (
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/execd/task-store/mocks"
	"github.com/execd/task-store/pkg/model"
	"github.com/execd/task-store/pkg/store"
	"github.com/execd/task-store/pkg/task"
	. "github.com/onsi/ginkgo"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"log"
)

var context = GinkgoT()

var _ = Describe("store", func() {
	var taskManager *task.StoreImpl
	var directRedis *miniredis.Miniredis
	var givenTaskSpec model.Spec
	var uuidGenMock mocks.UUIDGen

	BeforeEach(func() {
		s, err := miniredis.Run()
		if err != nil {
			panic(err)
		}

		redis := store.NewClient(s.Addr())
		uuidGenMock = mocks.UUIDGen{}
		taskManager = task.NewStoreImpl(redis, &uuidGenMock)
		directRedis = s
		givenTaskSpec = model.Spec{
			Name:     "test",
			Image:    "alpine",
			Init:     "init.sh",
			InitArgs: []string{"10"},
		}
	})

	AfterEach(func() {
		defer directRedis.Close()
	})

	Describe("pushing an item on the queue", func() {

		BeforeEach(func() {
			uuidGenMock.On("GenV4").Return(uuid.Must(uuid.NewV4()), nil)
		})

		It("should return pushed task id", func() {
			// Arrange
			givenId := uuid.Must(uuid.NewV4())

			// Act
			id, err := taskManager.Schedule(&givenId)
			failOnError(err)

			// Assert
			assert.Nil(context, err)
			assert.Equal(context, givenId, *id)
		})

		It("should return error if pushing onto queue fails", func() {
			//Arrange
			directRedis.Close()
			givenId := uuid.Must(uuid.NewV4())

			// Act
			_, err := taskManager.Schedule(&givenId)

			// Assert
			assert.NotNil(context, err)
		})
	})

	Describe("popping the next task", func() {

		BeforeEach(func() {
			uuidGenMock.On("GenV4").Return(uuid.Must(uuid.NewV4()), nil)
		})

		It("should return the task", func() {
			// Arrange
			givenId := uuid.Must(uuid.NewV4())
			_, err := taskManager.Schedule(&givenId)
			failOnError(err)

			// Act
			id, err := taskManager.PopNextTask()

			// Assert
			assert.Nil(context, err)
			assert.Equal(context, &givenId, id)
		})

		It("should return an error if popping task fails", func() {
			// Arrange
			directRedis.Close()

			// Act
			_, err := taskManager.PopNextTask()

			// Assert
			assert.NotNil(context, err)
			assert.Contains(context, err.Error(), "failed to retrieve next task to execute :")
		})
	})

	Describe("adding a task to executing set", func() {

		BeforeEach(func() {
			uuidGenMock.On("GenV4").Return(uuid.Must(uuid.NewV4()), nil)
		})

		It("should return error if adding fails", func() {
			// Arrange
			directRedis.Close()
			givenId := uuid.Must(uuid.NewV4())

			// Act

			err := taskManager.MoveTaskToExecutingSet(&givenId)

			// Assert
			assert.NotNil(context, err)
			assert.Contains(context, err.Error(), "failed to add task to executing set")
		})
	})

	Describe("storing a task", func() {
		It("should return task id", func() {
			// Arrange
			uuidGenMock.On("GenV4").Return(uuid.Must(uuid.NewV4()), nil)

			// Act
			id, err := taskManager.StoreTask(givenTaskSpec)

			// Assert
			assert.Nil(context, err)
			assert.NotNil(context, id)
			assert.NotEqual(context, id, uuid.Nil)
		})

		It("should return error if task with id already exists", func() {
			// Arrange
			givenID := uuid.Must(uuid.NewV4())
			uuidGenMock.On("GenV4").Return(givenID, nil)

			_, err := taskManager.StoreTask(givenTaskSpec)
			assert.Nil(context, err)

			// Act
			_, err = taskManager.StoreTask(givenTaskSpec)

			// Assert
			assert.NotNil(context, err)
			assert.Equal(context, fmt.Sprintf("task with id %s already exists", &givenID), err.Error())
		})

		It("should return error if storing task fails", func() {
			// Arrange
			uuidGenMock.On("GenV4").Return(uuid.Must(uuid.NewV4()), nil)
			directRedis.Close()

			// Act
			_, err := taskManager.StoreTask(givenTaskSpec)

			// Assert
			assert.NotNil(context, err)
			assert.Contains(context, err.Error(), "storing task with id")
		})
	})

	Describe("retrieving  a task", func() {

		BeforeEach(func() {
			uuidGenMock.On("GenV4").Return(uuid.Must(uuid.NewV4()), nil)
		})

		It("should return the task related to the given id", func() {
			// Arrange
			id, err := taskManager.StoreTask(givenTaskSpec)
			failOnError(err)

			// Act
			task, err := taskManager.GetTask(id)

			// Assert
			assert.Nil(context, err)
			assert.Equal(context, givenTaskSpec.Name, task.Name)
			assert.Equal(context, givenTaskSpec.Init, task.Init)
			assert.Equal(context, givenTaskSpec.InitArgs, task.InitArgs)
			assert.Equal(context, givenTaskSpec.Image, task.Image)
		})

		It("should return an error if task with given id does not exist", func() {
			// Arrange
			givenID := uuid.Must(uuid.NewV4())

			// Act
			_, err := taskManager.GetTask(&givenID)

			// Assert
			assert.NotNil(context, err)
			assert.Equal(context, fmt.Sprintf("failed to retrieve task with id %s", givenID.String()), err.Error())
		})

		It("should return an error if retrieved task fails to be re-constructed", func() {
			// Arrange
			badData := "test"
			id, err := taskManager.StoreTask(givenTaskSpec)
			failOnError(err)
			taskID := directRedis.Keys()[0]
			directRedis.Set(taskID, badData)

			// Act
			_, err = taskManager.GetTask(id)

			// Assert
			assert.NotNil(context, err)
			assert.Equal(context,
				fmt.Sprintf("failed to build task with id %s from retrieved data %s", id.String(), badData),
				err.Error())
		})
	})
})

func failOnError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
