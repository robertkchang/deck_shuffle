package main

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestMain(t *testing.T) {
  // just a dummy test
  assert.Equal(t, "hello", "hello")
}
