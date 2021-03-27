package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func classifyData(k int) {

	dataset := "class_data.txt"
	rawData, err := base.ParseCSVToInstances(dataset, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	cls := knn.NewKnnClassifier("euclidean", "linear", k)
	train, test := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(train)

	p, err := cls.Predict(test)
	if err != nil {
		fmt.Println(err)
		return
	}

	confusionMat, err := evaluation.GetConfusionMatrix(test, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(evaluation.GetSummary(confusionMat))
}
