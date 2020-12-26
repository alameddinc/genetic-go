package main

import (
	"fmt"
	"github.com/alameddinc/genetic-go/Genetic"
)

func main() {
	var population Genetic.Population
	population.InitPopulation(Genetic.GenesList, 5, 0.2)
	population.TiltCountLimit = 500
	population.IterationLimit = 2000
	population.Loop()
	fmt.Println((*population.Chromosomes)[0].Genes)
}
