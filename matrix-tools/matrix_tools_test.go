package matrix_tools

import "testing"

func AssertEqualHelper(t *testing.T, expected, actual interface{}) {
    if actual != expected {
        t.Errorf("Expected %v, got %v", expected, actual)
    }
}

func AssertPanic(t *testing.T) {
    defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
}

func TestTranspose(t *testing.T) {
    m := NewMatrix([][]float64{{1, 2, 3}})
    n := NewMatrix([][]float64{{1}, {2}, {3}})
    result := m.EqualTo(n.Transpose())
    AssertEqualHelper(t, true, result)

    result = n.EqualTo(m.Transpose())
    AssertEqualHelper(t, true, result)

    result = n.EqualTo(n.Transpose().Transpose())
    AssertEqualHelper(t, true, result)

    result = m.EqualTo(m.Transpose().Transpose())
    AssertEqualHelper(t, true, result)

    m = NewMatrix([][]float64{{0, 1}, {2, 3}, {4, 5}})
    n = NewMatrix([][]float64{{0, 2, 4}, {1, 3, 5}})
    result = m.EqualTo(n.Transpose())
    AssertEqualHelper(t, true, result)

    result = n.EqualTo(m.Transpose())
    AssertEqualHelper(t, true, result)

    result = n.EqualTo(n.Transpose().Transpose())
    AssertEqualHelper(t, true, result)

    result = m.EqualTo(m.Transpose().Transpose())
    AssertEqualHelper(t, true, result)
}

func TestEqualTo(t *testing.T) {
    m := NewMatrix([][]float64{{1, 2, 3}})
    n := NewMatrix([][]float64{{4, 5, 6}})
    expected := false
    result := m.EqualTo(n)
    AssertEqualHelper(t, expected, result)

    result = n.EqualTo(m)
    AssertEqualHelper(t, expected, result)

    n = NewMatrix([][]float64{{1, 2, 3}})
    expected = true
    result = m.EqualTo(n)
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1}, {2}, {3}})
    n = NewMatrix([][]float64{{4, 5, 6}})
    expected = false
    result = m.EqualTo(n)
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1, 2}, {3, 4}, {5, 6}})
    expected = false
    result = m.EqualTo(n)
    AssertEqualHelper(t, expected, result)

    n = NewMatrix([][]float64{{1, 2}, {3, 4}, {5, 6}})
    expected = true
    result = m.EqualTo(n)
    AssertEqualHelper(t, expected, result)
}

func TestAdd(t *testing.T) {
    m := NewMatrix([][]float64{{1, 2, 3}})
    n := NewMatrix([][]float64{{4, 5, 6}})
    expected := NewMatrix([][]float64{{5, 7, 9}})
    result := m.Add(n)
    AssertEqualHelper(t, true, expected.EqualTo(result))

    m = NewMatrix([][]float64{{1}, {2}, {3}})
    n = NewMatrix([][]float64{{4}, {5}, {6}})
    expected = NewMatrix([][]float64{{5}, {7}, {9}})
    result = m.Add(n)
    AssertEqualHelper(t, true, expected.EqualTo(result))

    m = NewMatrix([][]float64{{1, 2}, {2, 3}, {3, 4}})
    n = NewMatrix([][]float64{{4, 5}, {5, 6}, {6, 7}})
    expected = NewMatrix([][]float64{{5, 7}, {7, 9}, {9, 11}})
    result = m.Add(n)
    AssertEqualHelper(t, true, expected.EqualTo(result))
}

func TestMultiply(t *testing.T) {
    m := NewMatrix([][]float64{{1, 2}, {3, 4}})
    n := NewMatrix([][]float64{{5, 6}, {7, 8}})
    expected := NewMatrix([][]float64{{19, 22}, {43, 50}})
    result := m.Multiply(n)
    AssertEqualHelper(t, true, expected.EqualTo(result))
}

func TestDot(t *testing.T) {
    m := NewMatrix([][]float64{{1}, {2}, {3}})
    n := NewMatrix([][]float64{{4, 5, 6}})
    expected := 32.0
    result := m.Dot(n)
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{4}, {5}, {6}})
    n = NewMatrix([][]float64{{1, 2, 3}})
    expected = 32.0
    result = m.Dot(n)
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1}, {2}, {3}})
    n = NewMatrix([][]float64{{1, 2, 3}})
    expected = 14.0
    result = m.Dot(n)
    AssertEqualHelper(t, expected, result)
}

func TestIsVector(t *testing.T) {
    m := NewMatrix([][]float64{{1}})
    expected := true
    result := m.IsVector()
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1, 2, 3, 4}})
    result = m.IsVector()
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1}, {2}, {3}, {4}})
    result = m.IsVector()
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1, 2}, {2, 3}, {3, 4}, {4, 5}})
    expected = false
    result = m.IsVector()
    AssertEqualHelper(t, expected, result)
}

func TestIsSquare(t *testing.T) {
    m := NewMatrix([][]float64{{1}})
    expected := true
    result := m.IsSquare()
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1, 2}, {3, 4}})
    result = m.IsSquare()
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1, 2, 3}, {3, 4, 5}, {5, 6, 7}})
    result = m.IsSquare()
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1}, {3}})
    expected = false
    result = m.IsSquare()
    AssertEqualHelper(t, expected, result)

    m = NewMatrix([][]float64{{1, 2}, {3, 4}, {5, 6}})
    result = m.IsSquare()
    AssertEqualHelper(t, expected, result)
}
