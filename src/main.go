package main

import (
    "fmt"
    "net/http"
    "os"
    "gonum.org/v1/gonum/mat"
    "pca-project/src/pca"

    "github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

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

    // Step 5: Serve static files
    // Serve the "web" directory (HTML, CSS, JS files)
    http.Handle("/", http.FileServer(http.Dir("./web")))

    // Serve the "data" directory (to access reduced_data.csv)
    http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("./data"))))

    fmt.Println("Server is running at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
