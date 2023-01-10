
// Usage:
//  go run chapter4/gradient_descent.go chapter4/advertising.csv
package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"
	"path"

	// Frameworks
	"github.com/djthorpe/MachineLearning/util"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	LEARNING_RATE = 0.001
)

///////////////////////////////////////////////////////////////////////////////

func RunMain() int {
	if flag.NArg() < 1 {
		log.Println("Expected file argument", flag.NArg())
		return -1
	}

	table, _ := util.NewTable()
	filename := flag.Arg(0)
	if err := table.ReadCSV(filename, false, true, true); err != nil {
		log.Println("Unable to read CSV:", err)
		return -1
	}

	x_column := "Sales"
	y_column := "TV"
	if flag.NArg() > 1 {
		y_column = flag.Arg(1)
	}

	if x_data, err := table.FloatColumn(x_column, 0); err != nil {
		log.Println("Unable to create X samples:", err)
		return -1
	} else if y_data, err := table.FloatColumn(y_column, 0); err != nil {
		log.Println("Unable to create Y samples:", err)
		return -1
	} else if plot, err := plot.New(); err != nil {
		log.Println("Unable to create plot:", err)
		return -1
	} else {
		plot.X.Label.Text = x_column
		plot.Y.Label.Text = y_column

		if scatter, err := plotter.NewScatter(plot_points(x_data, y_data)); err != nil {
			log.Println("Unable to create plot:", err)
			return -1
		} else if line, err := plotter.NewLine(line_points(x_data, y_data)); err != nil {
			log.Println("Unable to create plot:", err)
			return -1
		} else {
			line.Color = color.RGBA{B: 255, A: 255}
			plot.Add(scatter, line)
			if err := plot.Save(4*vg.Inch, 4*vg.Inch, path.Base(filename)+"_"+x_column+"_"+y_column+".png"); err != nil {
				log.Println("Unable to create plot:", err)
				return -1
			}
		}
	}
	//

	return 0
}