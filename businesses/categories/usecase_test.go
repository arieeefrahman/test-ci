package categories_test

import (
	"echo-notes/businesses/categories"
	_categoryMock "echo-notes/businesses/categories/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	categoryRepository _categoryMock.Repository
	categoryService    categories.Usecase

	categoryDomain categories.Domain
)

func TestMain(m *testing.M) {
	categoryService = categories.NewCategoryUsecase(&categoryRepository)
	categoryDomain = categories.Domain{
		Name: "test",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All | Valid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]categories.Domain{categoryDomain}).Once()

		result := categoryService.GetAll()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All | InValid", func(t *testing.T) {
		categoryRepository.On("GetAll").Return([]categories.Domain{}).Once()

		result := categoryService.GetAll()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		categoryRepository.On("GetByID", "1").Return(categoryDomain).Once()

		result := categoryService.GetByID("1")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		categoryRepository.On("GetByID", "-1").Return(categories.Domain{}).Once()

		result := categoryService.GetByID("-1")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		categoryRepository.On("Create", &categoryDomain).Return(categoryDomain).Once()

		result := categoryService.Create(&categoryDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		categoryRepository.On("Create", &categories.Domain{}).Return(categories.Domain{}).Once()

		result := categoryService.Create(&categories.Domain{})

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		categoryRepository.On("Update", "1", &categoryDomain).Return(categoryDomain).Once()

		result := categoryService.Update("1", &categoryDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		categoryRepository.On("Update", "1", &categories.Domain{}).Return(categories.Domain{}).Once()

		result := categoryService.Update("1", &categories.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		categoryRepository.On("Delete", "1").Return(true).Once()

		result := categoryService.Delete("1")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		categoryRepository.On("Delete", "-1").Return(false).Once()

		result := categoryService.Delete("-1")

		assert.False(t, result)
	})
}
