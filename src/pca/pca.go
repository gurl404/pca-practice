package pca

import (
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
            standardized.Set(i, j, (col[i]-mean)/std)
        }
    }

    return standardized
}

// ComputeCovarianceMatrix calculates the covariance matrix of the dataset
func ComputeCovarianceMatrix(data *mat.Dense) *mat.SymDense {
    r, _ := data.Dims()
    cov := mat.NewSymDense(r, nil)
    cov.SymOuterK(1.0/float64(r-1), data)
    return cov
}

// EigenDecomposition performs eigenvalue and eigenvector decomposition of the covariance matrix
func EigenDecomposition(cov *mat.SymDense) ([]float64, *mat.Dense) {
    var eig mat.EigenSym
    ok := eig.Factorize(cov, true)
    if !ok {
        panic("Eigen decomposition failed")
    }

    eigenValues := eig.Values(nil)
    eigenVectors := mat.NewDense(len(eigenValues), len(eigenValues), nil)
    eig.VectorsTo(eigenVectors)

    return eigenValues, eigenVectors
}

// ProjectData projects the standardized dataset onto the top k principal components
func ProjectData(data *mat.Dense, eigenVectors *mat.Dense, k int) *mat.Dense {
    r, _ := data.Dims()
    reduced := mat.NewDense(r, k, nil)

    // Select top k eigenvectors
    reductionMatrix := eigenVectors.Slice(0, eigenVectors.RawMatrix().Rows, 0, k).(*mat.Dense)
    reduced.Mul(data, reductionMatrix)

    return reduced
}

// PCA orchestrates the PCA workflow: standardization, covariance calculation, eigen decomposition, and data projection
func PCA(data *mat.Dense, k int) *mat.Dense {
    // Step 1: Standardize data
    standardized := Standardize(data)

    // Step 2: Compute covariance matrix
    covarianceMatrix := ComputeCovarianceMatrix(standardized)

    // Step 3: Eigen decomposition
    _, eigenVectors := EigenDecomposition(covarianceMatrix)

    // Step 4: Project data
    reducedData := ProjectData(standardized, eigenVectors, k)

    return reducedData
}
