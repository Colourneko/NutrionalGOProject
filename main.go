package main

import "fmt"

func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(1),
		Sugars:              SugarGram(19),
		SaturatedFattyAcids: SaturatedFattyAcids(4),
		Sodium:              SodiumMilligram(700),
		Fruits:              FruitsPercent(56),
		Fibre:               FibreGram(5),
		Protein:             ProteinGram(8),
	}, Food)
	fmt.Printf("Nutritional score:%d\n", ns)
	fmt.Printf("NutriScore: %s\n", ns.GetNurtiScore())
}
