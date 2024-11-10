package repository

import (
	"fmt"
	"task-management-api/internal/app/model"
	"task-management-api/test"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestListTasks(t *testing.T) {
	db, mock, err := test.SetupMockDB(t)
	assert.NoError(t, err)
	defer db.DB()

	repo := NewTaskRepository(db)

	// Define test data
	taskID := 1
	task := model.Task{
		ID:          taskID,
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      model.STATUS_OF_TASK_IN_PROGRESS,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Test Case 1: Success Case
	t.Run("Success", func(t *testing.T) {
		// Set up expected query and rows for success scenario
		rows := sqlmock.NewRows([]string{"id", "title", "description", "status", "created_at", "updated_at"}).
			AddRow(task.ID, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)

		mock.ExpectQuery(`SELECT \* FROM "tasks"`).
			WithArgs().
			WillReturnRows(rows)

		// Run the test
		result, err := repo.ListTasks()
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, len(result))
	})

	// Test Case 2: Error Case - General DB Error
	t.Run("Error_DBError", func(t *testing.T) {
		// Simulate a general database error
		mock.ExpectQuery(`SELECT \* FROM "tasks"`).
			WithArgs().
			WillReturnError(fmt.Errorf("database connection error")) // Simulate a DB error

		// Run the test
		result, err := repo.ListTasks()
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database connection error", err.Error()) // Ensure the error message is as expected
	})
}

func TestGetTaskByID(t *testing.T) {
	db, mock, err := test.SetupMockDB(t)
	assert.NoError(t, err)
	defer db.DB()

	repo := NewTaskRepository(db)

	// Define test data
	taskID := 1
	task := model.Task{
		ID:          taskID,
		Title:       "Test Task",
		Description: "This is a test task",
		Status:      model.STATUS_OF_TASK_IN_PROGRESS,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Test Case 1: Success Case
	t.Run("Success", func(t *testing.T) {
		// Set up expected query and rows for success scenario
		rows := sqlmock.NewRows([]string{"id", "title", "description", "status", "created_at", "updated_at"}).
			AddRow(task.ID, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)

		mock.ExpectQuery(`SELECT \* FROM "tasks" WHERE id = \$1 ORDER BY "tasks"\."id" LIMIT \$2`).
			WithArgs(taskID, 1).
			WillReturnRows(rows)

		// Run the test
		result, err := repo.GetTaskByID(taskID)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, task.Title, result.Title)
		assert.Equal(t, task.Description, result.Description)
	})

	// Test Case 2: Error Case - Record Not Found
	t.Run("Error_RecordNotFound", func(t *testing.T) {
		// Simulate a scenario where the task is not found
		mock.ExpectQuery(`SELECT \* FROM "tasks" WHERE id = \$1 ORDER BY "tasks"\."id" LIMIT \$2`).
			WithArgs(taskID, 1).
			WillReturnError(gorm.ErrRecordNotFound) // Simulate no record found error

		// Run the test
		result, err := repo.GetTaskByID(taskID)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, gorm.ErrRecordNotFound, err) // Ensure the error is RecordNotFound
	})

	// Test Case 3: Error Case - General DB Error
	t.Run("Error_DBError", func(t *testing.T) {
		// Simulate a general database error
		mock.ExpectQuery(`SELECT \* FROM "tasks" WHERE id = \$1 ORDER BY "tasks"\."id" LIMIT \$2`).
			WithArgs(taskID, 1).
			WillReturnError(fmt.Errorf("database connection error")) // Simulate a DB error

		// Run the test
		result, err := repo.GetTaskByID(taskID)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database connection error", err.Error()) // Ensure the error message is as expected
	})
}

func TestCreateTask(t *testing.T) {
	db, mock, err := test.SetupMockDB(t)
	assert.NoError(t, err)
	defer db.DB()

	repo := NewTaskRepository(db)

	// Define test data for task creation
	taskRequest := model.TaskRequest{
		Title:       "New Task",
		Description: "New task description",
		Status:      model.STATUS_OF_TASK_TO_DO,
	}

	// Test Case 1: Success Case
	t.Run("Success", func(t *testing.T) {
		// Set up expected query for INSERT
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "tasks" \("title","description","status","created_at","updated_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5\) RETURNING "id"`).
			WithArgs(taskRequest.Title, taskRequest.Description, taskRequest.Status, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		// Run the test
		result, err := repo.CreateTask(taskRequest)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, taskRequest.Title, result.Title)
		assert.Equal(t, taskRequest.Description, result.Description)
		assert.Equal(t, model.STATUS_OF_TASK_TO_DO, result.Status)
	})

	// Test Case 2: Error Case - DB Error
	t.Run("Error_DBError", func(t *testing.T) {
		// Simulate a general database error
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "tasks" \("title","description","status","created_at","updated_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5\) RETURNING "id"`).
			WithArgs(taskRequest.Title, taskRequest.Description, taskRequest.Status, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnError(fmt.Errorf("database connection error")) // Simulate a DB error
		mock.ExpectRollback()

		// Run the test
		result, err := repo.CreateTask(taskRequest)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database connection error", err.Error()) // Ensure the error message is as expected
	})
}

func TestUpdateTask(t *testing.T) {
	db, mock, err := test.SetupMockDB(t)
	assert.NoError(t, err)
	defer db.DB()

	repo := NewTaskRepository(db)

	// Define test data for task update
	task := model.Task{
		ID:          1,
		Title:       "Updated Task",
		Description: "Updated task description",
		Status:      model.STATUS_OF_TASK_IN_PROGRESS,
	}

	// Test Case 1: Success Case
	t.Run("Success", func(t *testing.T) {
		// Set up expected query for UPDATE
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "tasks" SET "description"=\$1,"status"=\$2,"title"=\$3,"updated_at"=\$4 WHERE id = \$5`).
			WithArgs(task.Description, task.Status, task.Title, sqlmock.AnyArg(), task.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// Run the test
		err := repo.UpdateTask(task.ID, task)
		assert.NoError(t, err)
	})

	// Test Case 2: Error Case - DB Error (e.g., a general database error)
	t.Run("Error_DBError", func(t *testing.T) {
		// Simulate a general database error
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "tasks" SET "description"=\$1,"status"=\$2,"title"=\$3,"updated_at"=\$4 WHERE id = \$5`).
			WithArgs(task.Description, task.Status, task.Title, sqlmock.AnyArg(), task.ID).
			WillReturnError(fmt.Errorf("database connection error")) // Simulate a DB error
		mock.ExpectRollback()

		// Run the test
		err := repo.UpdateTask(task.ID, task)
		assert.Error(t, err)
		assert.Equal(t, "database connection error", err.Error()) // Ensure the error message is as expected
	})
}

func TestUpdateTaskStatus(t *testing.T) {
	db, mock, err := test.SetupMockDB(t)
	assert.NoError(t, err)
	defer db.DB()

	repo := NewTaskRepository(db)

	// Define test data for task status update
	taskID := 1
	newStatus := model.STATUS_OF_TASK_DONE

	// Test Case 1: Success Case
	t.Run("Success", func(t *testing.T) {
		// Set up expected query for updating the status
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "tasks" SET "status"=\$1,"updated_at"=\$2 WHERE id = \$3`).
			WithArgs(newStatus, time.Now(), taskID).
			WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate that 1 row was updated
		mock.ExpectCommit()

		// Run the test
		err := repo.UpdateTaskStatus(taskID, newStatus)
		assert.NoError(t, err)
	})

	// Test Case 2: Error Case - DB Error
	t.Run("Error_DBError", func(t *testing.T) {
		// Simulate a general database error
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "tasks" SET "status"=\$1,"updated_at"=\$2 WHERE id = \$3`).
			WithArgs(newStatus, time.Now(), taskID).
			WillReturnError(fmt.Errorf("database connection error")) // Simulate a DB error
		mock.ExpectRollback()

		// Run the test
		err := repo.UpdateTaskStatus(taskID, newStatus)
		assert.Error(t, err)
		assert.Equal(t, "database connection error", err.Error()) // Ensure the error message is as expected
	})
}

func TestDeleteTaskByID(t *testing.T) {
	db, mock, err := test.SetupMockDB(t)
	assert.NoError(t, err)
	defer db.DB()

	repo := NewTaskRepository(db)

	// Define test data for task deletion
	taskID := 1

	// Test Case 1: Success Case
	t.Run("Success", func(t *testing.T) {
		// Set up expected query for deleting a task
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "tasks" WHERE id = \$1`).
			WithArgs(taskID).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		// Run the test
		err := repo.DeleteTaskByID(taskID)
		assert.NoError(t, err)
	})

	// Test Case 2: Error Case - DB Error
	t.Run("Error_DBError", func(t *testing.T) {
		// Simulate a general database error
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "tasks" WHERE id = \$1`).
			WithArgs(taskID).
			WillReturnError(fmt.Errorf("database connection error")) // Simulate a DB error
		mock.ExpectRollback()

		// Run the test
		err := repo.DeleteTaskByID(taskID)
		assert.Error(t, err)
		assert.Equal(t, "database connection error", err.Error()) // Ensure the error message is as expected
	})
}
