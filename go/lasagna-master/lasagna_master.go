package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
	var refAvgTime int = 2

	if avgTime != 0 {
		refAvgTime = avgTime
	}

	return len(layers) * refAvgTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodlesQty int, sauceQty float64) {
	const noodlesWeight int = 50
	const sauceVolume float64 = 0.2

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodlesQty += noodlesWeight
		}

		if layers[i] == "sauce" {
			sauceQty += sauceVolume
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(ingredients1, ingredients2 []string) {
	ingredients2[len(ingredients2)-1] = ingredients1[len(ingredients1)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, nbPortions int) []float64 {
	if nbPortions == 2 {
		return quantities
	}

	var quantitiesFor1Portion []float64
	var res []float64

	for i := 0; i < len(quantities); i++ {
		quantitiesFor1Portion = append(quantitiesFor1Portion, quantities[i]/2)
	}

	if nbPortions == 1 {
		return quantitiesFor1Portion
	}

	for i := 0; i < len(quantitiesFor1Portion); i++ {
		res = append(res, quantitiesFor1Portion[i]*float64(nbPortions))
	}

	return res
}
