package main

type Datum struct {
	x Real
	y Real
}

func randomizeData() {
	gap := (dataXMax - dataXMin) / Real(numData)
	for i := 0; i < numData; i++ {
		x := dataXMin + Real(i)*gap
		y := randomReal(dataYMin, dataYMax)
		data[i] = Datum{x, y}
	}
	displayData()
}

func displayData() {
	// for i := range data {
	// 	//fmt.Printf("data[%d] = %+f,%+f\n", i, data[i].x, data[i].y)
	// }
}
