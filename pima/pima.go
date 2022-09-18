package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

// func readCSVFromUrl(url string) ([][]string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()
// 	reader := csv.NewReader(resp.Body)
// 	reader.Comma = ';'
// 	data, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

func main() {
	// csv_name := "datasets/iris_headers.csv"
	csv_name := "../datasets/pima.csv"
	// csv_name := "https://raw.githubusercontent.com/mwaskom/seaborn-data/master/iris.csv"
	rawData, err := base.ParseCSVToInstances(csv_name, true)
	// rawData, err := readCSVFromUrl(csv_name)
	if err != nil {
		panic(err)
	}

	// fmt.Println(rawData)

	//Initialises a new KNN classifier
	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	//Do a training-test split
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}
	fmt.Println(predictions)

	// Prints precision/recall metrics
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}
