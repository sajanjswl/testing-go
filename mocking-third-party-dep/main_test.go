package main

import (
	"testing"
	"github.com/sajanjswl/testing-go/mocking-third-party-dep/mocks"
	"github.com/stretchr/testify/assert"
	
)

func TestBar(t *testing.T) {
	mockThing := Thing{}

	mockDAL := &mocks.DataAccessLayer{}
	mockDAL.On("Insert", mockThing).Return(nil)

	actual := Bar(mockDAL)

	mockDAL.AssertExpectations(t)

	assert.Equal(t, mockThing, actual, "should return a Thing")
}