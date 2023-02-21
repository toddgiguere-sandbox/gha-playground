package tests

import "testing"

func TestOne(t *testing.T) {
	t.Parallel()

	One()
}

func TestTwo(t *testing.T) {
	t.Parallel()

	Two()
}

func threeCases(t *testing.T, text string) {
	t.Parallel()
	Three(text)
}

func TestThree(t *testing.T) {
	t.Parallel()

	textList := []string{"A", "B", "C"}

	for _, text := range textList {
		t.Run(text, func(t *testing.T) {
			threeCases(t, text)
		})
	}
}
