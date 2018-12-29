package gochuck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandom(t *testing.T) {
	fact, err := GetRandom()

	assert.Nil(t, err)
	assert.NotEmpty(t, fact.Value)
}

func TestGetCategories(t *testing.T) {
	categories, err := GetCategories()

	assert.Nil(t, err)
	assert.NotEmpty(t, categories)
}

func TestGetByQuery(t *testing.T) {
	collection, err := GetByQuery("Circle of Life")

	assert.Nil(t, err)
	assert.NotEmpty(t, collection.Results)

	for _, result := range collection.Results {
		assert.Contains(t, result.Value, "Circle of Life")
	}
}

func TestGetRandomByCategory(t *testing.T) {
	fact, err := GetRandomByCategory("food")

	assert.Nil(t, err)
	assert.NotEmpty(t, fact.Value)
	assert.Contains(t, fact.Category, "food")
}
