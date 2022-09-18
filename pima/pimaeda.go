package main

import (
	"fmt"

	"github.com/sjwhitworth/golearn/base"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func pimaeda() {
	// csv_name := "datasets/iris_headers.csv"
	csv_name := "../datasets/pima.csv"
	// csv_name := "https://raw.githubusercontent.com/mwaskom/seaborn-data/master/iris.csv"
	rawData, err := base.ParseCSVToInstances(csv_name, true)
	// rawData, err := readCSVFromUrl(csv_name)
	if err != nil {
		panic(err)
	}

	fmt.Println(rawData)

	groupA := plotter.Values{20, 35, 30, 35, 27}
	groupB := plotter.Values{25, 32, 34, 20, 25}
	groupC := plotter.Values{12, 28, 15, 21, 8}

	p, _ := plot.New()

	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	w := vg.Points(20)

	barsA, err := plotter.NewBarChart(groupA, w)
	if err != nil {
		panic(err)
	}
	barsA.LineStyle.Width = vg.Length(0)
	barsA.Color = plotutil.Color(0)
	barsA.Offset = -w

	barsB, err := plotter.NewBarChart(groupB, w)
	if err != nil {
		panic(err)
	}
	barsB.LineStyle.Width = vg.Length(0)
	barsB.Color = plotutil.Color(1)

	barsC, err := plotter.NewBarChart(groupC, w)
	if err != nil {
		panic(err)
	}
	barsC.LineStyle.Width = vg.Length(0)
	barsC.Color = plotutil.Color(2)
	barsC.Offset = w

	p.Add(barsA, barsB, barsC)
	p.Legend.Add("Group A", barsA)
	p.Legend.Add("Group B", barsB)
	p.Legend.Add("Group C", barsC)
	p.Legend.Top = true
	p.NominalX("One", "Two", "Three", "Four", "Five")

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}

}
