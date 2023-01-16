package matrix_tools

type Matrix struct {
    data [][]float64
    rows int
    cols int
}

func NewMatrix(data [][]float64) Matrix {
    rows := len(data)
    cols := len(data[0])
    return Matrix{data: data, rows: rows, cols: cols}
}

func (m Matrix) Transpose() Matrix {
    data := make([][]float64, m.cols)
    for i := 0; i < m.cols; i++ {
        data[i] = make([]float64, m.rows)
        for j := 0; j < m.rows; j++ {
            data[i][j] = m.data[j][i]
        }
    }
    return Matrix{data: data, rows: m.cols, cols: m.rows}
}

func (m Matrix) EqualTo(n Matrix) bool {
    if m.cols != n.cols || m.rows != n.rows {
        return false
    }
    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            if m.data[i][j] != n.data[i][j] {
                return false
            }
        }
    }
    return true
}

func (m Matrix) Add(n Matrix) Matrix {
    if m.cols != n.cols || m.rows != n.rows {
        panic("matrices have different dimensions")
    }

    data := make([][]float64, m.rows)
    for i := 0; i < m.rows; i++ {
        data [i] = make([]float64, m.cols)
        for j := 0; j < m.cols; j++ {
            data[i][j] = m.data[i][j] + n.data[i][j]
        }
    }
    return Matrix{data: data, rows: m.rows, cols: m.cols}
}

func (m Matrix) Multiply(n Matrix) Matrix {
    if m.cols != n.rows {
        panic("matrices have incompatible dimensions")
    }

    data := make([][]float64, m.rows)
    for i := 0; i < m.rows; i++ {
        data[i] = make([]float64, n.cols)
        for j := 0; j < n.cols; j++ {
            sum := 0.0
            for k := 0; k < m.cols; k++ {
                sum += m.data[i][k] * n.data[k][j]
            }
            data[i][j] = sum
        }
    }
    return Matrix{data: data, rows: m.rows, cols: n.cols}
}

func (m Matrix) Dot(n Matrix) float64 {
    if m.cols != 1 || n.rows != 1 {
        panic("matrices should be vectors")
    }

    if m.rows != n.cols {
        panic("vectors have different lengths")
    }

    sum := 0.0
    for i := 0; i < m.rows; i++ {
        sum += m.data[i][0] * n.data[0][i]
    }
    return sum
}

func (m Matrix) IsVector() bool {
    return m.cols == 1 || m.rows == 1
}

func (m Matrix) IsSquare() bool {
    return m.cols == m.rows
}
