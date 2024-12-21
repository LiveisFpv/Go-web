package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	client := getTestClient()

	response, err := client.createStudent(1, "ПИ-22-1", "a@mail.ru", "test2", "test1", "test3")
	assert.NoError(t, err)
	assert.Equal(t, response.Data.Id_num_student, int64(1))
	assert.Equal(t, response.Data.Name_group, "ПИ-22-1")
	assert.Equal(t, response.Data.Email_student, "a@mail.ru")
	assert.Equal(t, response.Data.Second_name_student, "test2")
	assert.Equal(t, response.Data.First_name_student, "test1")
	assert.Equal(t, response.Data.Surname_student, "test3")
}
func TestGetStudent(t *testing.T) {
	client := getTestClient()

	student, err := client.getStudent(1)
	assert.NoError(t, err)
	assert.Equal(t, student.Data.Id_num_student, int64(1))
	assert.Equal(t, student.Data.Name_group, "ПИ-22-1")
	assert.Equal(t, student.Data.Email_student, "a@mail.ru")
	assert.Equal(t, student.Data.Second_name_student, "test2")
	assert.Equal(t, student.Data.First_name_student, "test1")
	assert.Equal(t, student.Data.Surname_student, "test3")
}
