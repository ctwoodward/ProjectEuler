package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Problem1(t *testing.T) {
	assert.Equal(t, 233168, Problem1())
}
func Test_Problem2(t *testing.T) {
	assert.Equal(t, 4613732, Problem2())
}
func Test_fibProblem2(t *testing.T) {
	assert.Equal(t, 5, fibProblem2(4))
}
func Test_Problem3(t *testing.T) {
	assert.Equal(t, 6857, Problem3())
}
func Test_Problem4(t *testing.T) {
	assert.Equal(t, 906609, Problem4())
}
func Test_Problem5(t *testing.T) {
	assert.Equal(t, 232792560, Problem5())
}
func Test_Problem6(t *testing.T) {
	assert.Equal(t, 25164150, Problem6())
}
