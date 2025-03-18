## PCA Practice Project
This project demonstrates how to perform Principal Component Analysis (PCA) on numerical datasets using the Go programming language and the gonum library. The project includes modular code for data preprocessing, PCA computation, and debugging.

### Files and Their Roles
1. main.go
The entry point of the project.

Orchestrates the workflow:

Loads the dataset from data/input.csv.

Preprocesses the data and prepares it for PCA.

Calls the PCA implementation to reduce dimensions.

Saves the transformed dataset to data/reduced_data.csv.

Also includes debugging statements to verify that everything runs smoothly.

2. pca/pca.go
Contains the core PCA logic, split into well-defined functions:

Standardize: Normalizes the dataset to have a mean of 0 and standard deviation of 1 for each feature.

ComputeCovarianceMatrix: Calculates the covariance matrix based on the standardized dataset.

EigenDecomposition: Extracts eigenvalues and eigenvectors from the covariance matrix.

ProjectData: Projects the dataset onto the top principal components to reduce dimensionality.

PCA: Combines all the above steps into a seamless workflow.

Includes debugging statements to verify dimensions at every step.

3. pca/utils.go
Provides helper functions for data handling and preprocessing:

LoadCSV: Reads a CSV file and converts it into a numerical dataset.

SaveCSV: Writes a transformed dataset (e.g., PCA output) to a CSV file.

FlattenData: Converts a 2D slice of floats into a 1D slice for matrix operations.

ReshapeData: Converts a 1D slice back into a 2D slice for easier processing.

Includes error handling for invalid or missing data (e.g., handling N/A values in the CSV file).

4. data/input.csv
The raw dataset used for PCA. This file contains numerical features and rows of data points for processing.

The example dataset includes measurements like water temperature, air temperature, pH, and salinity.

5. data/reduced_data.csv
The output file containing the PCA-transformed dataset.

After running the program, this file will include the reduced dataset with only the principal components specified in the code.

6. go.mod
Defines the project as a Go module.

Lists the dependencies required for the project, such as gonum.

7. go.sum
Automatically generated file that contains checksum information for all module dependencies.

### How to Run the Project
Install Dependencies:

Ensure the gonum library is installed by running:


`go get gonum.org/v1/gonum/mat`

Run the Program:

Execute the main.go file to perform PCA:


`go run src/main.go`

Check the Output:

The reduced dataset will be saved to data/reduced_data.csv.

