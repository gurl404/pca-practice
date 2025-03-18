package main

import (
    "fmt"
    "os"
    "gonum.org/v1/gonum/mat"
    "pca-project/src/pca"

    "github.com/joho/godotenv" 
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    // Get file paths from environment variables
    inputFile := os.Getenv("PCA_INPUT_FILE")
    outputFile := os.Getenv("PCA_OUTPUT_FILE")

    // Step 1: Load data from CSV
    data, err := pca.LoadCSV(inputFile)
    if err != nil {
        panic(fmt.Sprintf("Failed to load CSV: %v", err))
    }

    // Step 2: Convert data to Dense matrix
    dataMatrix := mat.NewDense(len(data), len(data[0]), pca.FlattenData(data))

    // Step 3: Perform PCA
    numComponents := 2 // Adjust as needed
    reducedData := pca.PCA(dataMatrix, numComponents)

    // Step 4: Save reduced data to CSV
    err = pca.SaveCSV(pca.ReshapeData(reducedData.RawMatrix().Data, len(data), numComponents), outputFile)
    if err != nil {
        panic(fmt.Sprintf("Failed to save reduced data: %v", err))
    }

    fmt.Printf("PCA completed successfully! Reduced data saved to %s\n", outputFile)
}
