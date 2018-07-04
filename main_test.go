package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGame(t *testing.T) {
	assert := assert.New(t)
	g := new(Game)
	assert.Equal(g.answer, 0, "should be equal")

	g.initialize()
	assert.NotEqual(g.answer, 0, "should not be equal")

	g.guess = g.answer
	assert.Equal(g.checkAnswer(), GUESS_CORRECT, "should be equal")

	g.guess += 1
	assert.Equal(g.checkAnswer(), GUESS_TOO_HIGH, "should be equal")

	g.guess -= 2
	assert.Equal(g.checkAnswer(), GUESS_TOO_LOW, "should be equal")
}
