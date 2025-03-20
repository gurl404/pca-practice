package pca

import (
    "fmt"
    "gonum.org/v1/gonum/mat"
    "gonum.org/v1/gonum/stat"
)

// Standardize normalizes the dataset by subtracting the mean and dividing by the standard deviation
func Standardize(data *mat.Dense) *mat.Dense {
    r, c := data.Dims()
    standardized := mat.NewDense(r, c, nil)

    for j := 0; j < c; j++ {
        col := mat.Col(nil, j, data)
        mean, std := stat.MeanStdDev(col, nil)
        for i := 0; i < r; i++ {
            if std == 0 { // Handle zero variance
                standardized.Set(i, j, 0)
            } else {
                standardized.Set(i, j, (col[i]-mean)/std)
            }
        }
    }

    return standardized
}

// ComputeCovarianceMatrix calculates the covariance matrix of the dataset
func ComputeCovarianceMatrix(data *mat.Dense) *mat.SymDense {
    rows, cols := data.Dims() 

    // Center the data by subtracting the mean of each column
    centered := mat.NewDense(rows, cols, nil)
    centered.CloneFrom(data)
    for j := 0; j < cols; j++ {
        col := mat.Col(nil, j, data) 
        mean := stat.Mean(col, nil) 
        for i := 0; i < rows; i++ {
            centered.Set(i, j, centered.At(i, j)-mean) 
        }
    }

    // Transpose the centered data for covariance calculation
    transposed := mat.NewDense(cols, rows, nil)
    transposed.CloneFrom(centered.T())

    // Compute the covariance matrix
    cov := mat.NewSymDense(cols, nil)
    // Pass transposed data
    cov.SymOuterK(1.0/float64(rows-1), transposed) 
    return cov
}

// EigenDecomposition performs eigenvalue and eigenvector decomposition of the covariance matrix
func EigenDecomposition(cov *mat.SymDense) ([]float64, *mat.Dense) {
    var eig mat.EigenSym
    ok := eig.Factorize(cov, true)
    if !ok {
        panic("Eigen decomposition failed: ensure covariance matrix is well-formed")
    }

    eigenValues := eig.Values(nil)
    eigenVectors := mat.NewDense(len(eigenValues), len(eigenValues), nil)
    eig.VectorsTo(eigenVectors)

    return eigenValues, eigenVectors
}

// ProjectData projects the standardized dataset onto the top k principal components
func ProjectData(data *mat.Dense, eigenVectors *mat.Dense, k int) *mat.Dense {
    r, c := data.Dims()
    eigenRows, eigenCols := eigenVectors.Dims()

    // Debugging: Print matrix dimensions
    fmt.Printf("Data Matrix Dimensions: %d x %d\n", r, c)
    fmt.Printf("Eigenvector Matrix Dimensions: %d x %d\n", eigenRows, eigenCols)

    reduced := mat.NewDense(r, k, nil)

    // Select top k eigenvectors
    reductionMatrix := eigenVectors.Slice(0, eigenRows, 0, k).(*mat.Dense)

    // Debugging: Print dimensions of the reduction matrix
    redRows, redCols := reductionMatrix.Dims()
    fmt.Printf("Reduction Matrix Dimensions: %d x %d\n", redRows, redCols)

    // Perform matrix multiplication
    reduced.Mul(data, reductionMatrix)

    return reduced
}

// PCA orchestrates the PCA workflow: standardization, covariance calculation, eigen decomposition, and data projection
func PCA(data *mat.Dense, k int) *mat.Dense {
    // Step 1: Standardize data
    standardized := Standardize(data)
    stdRows, stdCols := standardized.Dims()
    fmt.Printf("Standardized Data Dimensions: %d x %d\n", stdRows, stdCols)

    // Step 2: Compute covariance matrix
    covarianceMatrix := ComputeCovarianceMatrix(standardized)
    covRows, covCols := covarianceMatrix.Dims()
    fmt.Printf("Covariance Matrix Dimensions: %d x %d\n", covRows, covCols)

    // Step 3: Eigen decomposition
    _, eigenVectors := EigenDecomposition(covarianceMatrix)
    eigRows, eigCols := eigenVectors.Dims()
    fmt.Printf("Eigenvector Matrix Dimensions: %d x %d\n", eigRows, eigCols)

    // Step 4: Project data
    reducedData := ProjectData(standardized, eigenVectors, k)

    return reducedData
}
