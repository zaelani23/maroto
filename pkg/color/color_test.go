package color_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zaelani23/maroto/pkg/color"
	"testing"
)

func TestNewWhite(t *testing.T) {
	// Act
	white := color.NewWhite()

	// Assert
	assert.Equal(t, 255, white.Red)
	assert.Equal(t, 255, white.Green)
	assert.Equal(t, 255, white.Blue)
}

func TestColor_IsWhite(t *testing.T) {
	// Act
	white := color.NewWhite()

	// Assert
	assert.True(t, white.IsWhite())
}
