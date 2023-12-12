package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileString(t *testing.T) {
	assert.Equal(t, "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet", getFileString("test.txt"))
}

func TestSolve(t *testing.T) {
	assert.Equal(t, 142, solve("test.txt"))
}
