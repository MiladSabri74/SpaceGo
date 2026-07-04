package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepareTimeForEachLayer ...int) int {
	prepareTime := 2
	if len(prepareTimeForEachLayer) != 0 {
		if prepareTimeForEachLayer[0] != 0 {
			prepareTime = prepareTimeForEachLayer[0]
		}
	}
	return len(layers) * prepareTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	souces := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			souces += 0.2
		}
	}
	return noodles, souces
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(firendList []string, ownList []string) {
	if len(ownList) == 0 || len(firendList) == 0 {
		return
	}
	(ownList)[len(ownList)-1] = firendList[len(firendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	factor := float64(portions) / 2.0
	for i, q := range quantities {
		scaled[i] = q * factor
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
