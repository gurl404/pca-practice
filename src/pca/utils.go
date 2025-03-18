package pca

import (
    "encoding/csv"
    "os"
	"fmt" 
    "strconv"
    "gonum.org/v1/gonum/mat"
)

// LoadCSV reads water quality data CSV file and returns the data as a 2D float64 slice
func LoadCSV(filePath string) ([][]float64, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var data [][]float64
    for i, row := range records {
        if i == 0 { // Skip the header row
            continue
        }

        var numericRow []float64
        for _, colIndex := range []int{3, 4, 5, 6, 7, 8, 9} { // Indices for numeric columns
            val := row[colIndex]
            if val == "N/A" || val == "" {
                // Handle missing values (replace with 0, mean, etc.)
                val = "0"
            }
            num, err := strconv.ParseFloat(val, 64)
            if err != nil {
                return nil, fmt.Errorf("Invalid value in CSV: %v", err)
            }
            numericRow = append(numericRow, num)
        }
        data = append(data, numericRow)
    }

    return data, nil
}


// SaveCSV writes a 2D float64 slice to a CSV file
func SaveCSV(data [][]float64, filePath string) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, row := range data {
        strRow := make([]string, len(row))
        for i, val := range row {
            strRow[i] = strconv.FormatFloat(val, 'f', -1, 64)
        }
        err := writer.Write(strRow)
        if err != nil {
            return err
        }
    }

    return nil
}

// FlattenData converts a 2D slice of float64 to a 1D slice for gonum matrices
func FlattenData(data [][]float64) []float64 {
    var flat []float64
    for _, row := range data {
        flat = append(flat, row...)
    }
    return flat
}

// ReshapeData converts a 1D slice back into a 2D slice with given dimensions
func ReshapeData(data []float64, rows, cols int) [][]float64 {
    reshaped := make([][]float64, rows)
    for i := 0; i < rows; i++ {
        start := i * cols
        end := start + cols
        reshaped[i] = data[start:end]
    }
    return reshaped
}

// PrintMatrix displays a Dense matrix in a readable format (for debugging)
func PrintMatrix(title string, matrix mat.Matrix) {
    fmt.Println(title)
    fmt.Printf("%v\n", mat.Formatted(matrix, mat.Prefix(""), mat.Excerpt(0)))
}
