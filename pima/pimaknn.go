package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rocketlaunchr/dataframe-go"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func pimaknn() {
	// csv_name := "datasets/iris_headers.csv"
	// csv_name := "../datasets/pima.csv"
	url := "https://raw.githubusercontent.com/jameshtwose/example_deliverables/main/classification_examples/pima_diabetes/diabetes.csv"
	// csv_name := "https://raw.githubusercontent.com/mwaskom/seaborn-data/master/iris.csv"
	// rawData, err := base.ParseCSVToInstances(csv_name, true)
	rawData, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}

	s1 := dataframe.NewSeriesFloat64("col1", nil, 1.2, 2.3, 3.4, 4.5, 5.6)
	s2 := dataframe.NewSeriesFloat64("col2", nil, 6.7, 7.8, 8.9, 9.0, 10.1)
	df := dataframe.NewDataFrame(s1, s2)

	fmt.Println(df)

	fmt.Println(rawData[0][0])

	// fmt.Println(rawData)
	fmt.Println(reflect.TypeOf(rawData))

	// //Initialises a new KNN classifier
	// cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	// //Do a training-test split
	// trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	// cls.Fit(trainData)

	// //Calculates the Euclidean distance and returns the most popular label
	// predictions, err := cls.Predict(testData)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(predictions)

	// // Prints precision/recall metrics
	// confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	// if err != nil {
	// 	panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	// }
	// fmt.Println(evaluation.GetSummary(confusionMat))
}
