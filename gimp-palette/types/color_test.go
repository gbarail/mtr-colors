package types

import "testing"

func TestCompareColors(t *testing.T) {
	type TestData struct {
		c1, c2   Color
		expected int
	}

	testData := []TestData{
		{ // name c1 greater than c2
			c1:       Color{Name: "B", RGB: RGB{0, 0, 0}},
			c2:       Color{Name: "A", RGB: RGB{0, 0, 0}},
			expected: 1,
		},
		{ // name c1 less than c2
			c1:       Color{Name: "A", RGB: RGB{0, 0, 0}},
			c2:       Color{Name: "B", RGB: RGB{0, 0, 0}},
			expected: -1,
		},
		{ // same name, RGB c1 greater than c2
			c1:       Color{Name: "Color", RGB: RGB{10, 20, 30}},
			c2:       Color{Name: "Color", RGB: RGB{1, 2, 3}},
			expected: 1,
		},
		{ // same name, RGB c1 less than c2
			c1:       Color{Name: "Color", RGB: RGB{1, 2, 3}},
			c2:       Color{Name: "Color", RGB: RGB{1, 3, 2}},
			expected: -1,
		},
		{ // same name, RGB c1 equal to c2
			c1:       Color{Name: "Color", RGB: RGB{1, 2, 3}},
			c2:       Color{Name: "Color", RGB: RGB{1, 2, 3}},
			expected: 0,
		},
	}

	for _, td := range testData {
		if c := CompareColors(&td.c1, &td.c2); c != td.expected {
			t.Errorf("CompareColors(%v, %v) = %v, expected %v",
				td.c1, td.c2, c, td.expected)
		}
	}
}

func TestCompareColorNames(t *testing.T) {
	type TestData struct {
		c1, c2   Color
		expected int
	}

	colors := []Color{
		{Name: "red"},
		{Name: "blue"},
	}

	testData := []TestData{
		{ // greater than
			c1:       colors[0],
			c2:       colors[1],
			expected: 1,
		},
		{ // less than
			c1:       colors[1],
			c2:       colors[0],
			expected: -1,
		},
		{ // equal
			c1:       colors[0],
			c2:       colors[0],
			expected: 0,
		},
	}

	for _, td := range testData {
		if c := compareColorNames(&td.c1, &td.c2); c != td.expected {
			t.Errorf("compareColorNames(%v, %v) = %v, expected %v",
				td.c1, td.c2, c, td.expected)
		}
	}
}

func TestCompareRGB(t *testing.T) {
	type TestData struct {
		v1, v2   RGB
		expected int
	}

	testData := []TestData{
		{ // greater than
			v1:       RGB{2, 10, 100},
			v2:       RGB{1, 100, 10},
			expected: 1,
		},
		{ // greater than
			v1:       RGB{2, 10, 100},
			v2:       RGB{2, 8, 100},
			expected: 1,
		},
		{ // greater than
			v1:       RGB{2, 10, 100},
			v2:       RGB{2, 10, 80},
			expected: 1,
		},
		{ // less than
			v1:       RGB{1, 100, 10},
			v2:       RGB{2, 10, 100},
			expected: -1,
		},
		{ // less than
			v1:       RGB{2, 8, 100},
			v2:       RGB{2, 10, 100},
			expected: -1,
		},
		{ // less than
			v1:       RGB{2, 10, 80},
			v2:       RGB{2, 10, 100},
			expected: -1,
		},
		{ // equal
			v1:       RGB{2, 10, 80},
			v2:       RGB{2, 10, 80},
			expected: 0,
		},
	}

	for _, td := range testData {
		if c := compareRGB(&td.v1, &td.v2); c != td.expected {
			t.Errorf("compareRGB(%v, %v) = %v, expected %v",
				td.v1, td.v2, c, td.expected)
		}
	}
}

func TestCompareRGBValues(t *testing.T) {
	type TestData struct {
		v1, v2   RGBValue
		expected int
	}

	testData := []TestData{
		{ // greater than
			v1:       14,
			v2:       2,
			expected: 1,
		},
		{ // less than
			v1:       2,
			v2:       14,
			expected: -1,
		},
		{ // equal
			v1:       10,
			v2:       10,
			expected: 0,
		},
	}

	for _, td := range testData {
		if c := compareRGBValues(td.v1, td.v2); c != td.expected {
			t.Errorf("compareRGBValues(%v, %v) = %v, expected %v",
				td.v1, td.v2, c, td.expected)
		}
	}
}
