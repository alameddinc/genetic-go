package Genetic

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Population struct {
	Size           int
	Capacity       int
	Chromosomes    *[]Chromosome
	MutationRate   float64
	LastBest       float64
	Iteration      int
	IterationLimit int
	TiltCount      int
	TiltCountLimit int
	VariationRate  float64
}

func (p *Population) InitPopulation(genList []Gene, PopulationSize int, mutationRate float64) {
	p.Size = PopulationSize
	p.Capacity = PopulationSize
	p.MutationRate = mutationRate
	p.Chromosomes = &[]Chromosome{}
	p.ConfiguratorPopulation()
	for i := 0; i < p.Size; i++ {
		tmpChromosome := Chromosome{}
		tmpChromosome.Init(genList[1:], genList[0])
		*p.Chromosomes = append(*p.Chromosomes, tmpChromosome)
	}
	p.Sort()
}

func (p *Population) ConfiguratorPopulation() {
	p.Iteration = 0
	p.TiltCount = 0
	if p.VariationRate == 0 {
		p.VariationRate = 0.3
	}
	if p.IterationLimit == 0 {
		p.IterationLimit = 10000
	}
	if p.TiltCountLimit == 0 {
		p.TiltCountLimit = 1000
	}
	p.LastBest = math.MaxFloat64
}

func (p *Population) Sort() {
	sort.Slice(*p.Chromosomes, func(i, j int) bool {
		return (*p.Chromosomes)[i].Score < (*p.Chromosomes)[j].Score
	})
}

func (p *Population) Crossing(A, B int) {
	newA, newB := (*p.Chromosomes)[A-1].Crossover((*p.Chromosomes)[B-1])
	*p.Chromosomes = append(*p.Chromosomes, newA, newB)
	p.Size = p.Size + 2
}

func (p *Population) Loop() {
	for p.TiltCountLimit > p.TiltCount && p.IterationLimit > p.Iteration {
		p.VariationalGrownUp()
		p.StaticGrownUp()
		p.CleanUp()
		p.Iteration++
	}
}
func (p *Population) VariationalGrownUp() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < int(float64(p.Capacity)*p.VariationRate); i++ {
		A := (r.Int() % len(*p.Chromosomes)) + 1
		B := (r.Int() % len(*p.Chromosomes)) + 1
		C := (r.Int() % len(*p.Chromosomes)) + 1
		for A == B {
			B = (r.Int() % len(*p.Chromosomes)) + 1
		}
		p.Crossing(A, B)
		p.Mutation(C)
	}
}

func (p *Population) StaticGrownUp() {
	// Crossing with the Best Chromosomes
	p.Crossing(1, 2)
	p.Crossing(1, 3)
	p.Crossing(2, 3)
	// Mutation with the Best Chromosomes and Worst Chromosomes
	count := p.Capacity
	p.Mutation(1)
	p.Mutation(2)
	p.Mutation(count)
	p.Mutation(count - 1)
	p.Mutation(count - 2)
}

func (p *Population) Mutation(A int) {
	newA := (*p.Chromosomes)[A-1].Mutation(p.MutationRate)
	*p.Chromosomes = append(*p.Chromosomes, newA)
	p.Size = p.Size + 1
}

func (p *Population) CleanUp() {
	p.Sort()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for p.Size > p.Capacity {
		totalList := p.ScoreResult()
		itemScore := r.Float64() * totalList[len(totalList)-1]
		for k, v := range totalList {
			if v > itemScore {
				*p.Chromosomes = append((*p.Chromosomes)[:k], (*p.Chromosomes)[k+1:]...)
				break
			}
		}
		p.Size--
	}
	p.TiltCount++
	if p.TiltCount%100 == 0 {
		fmt.Println(p.TiltCount)
	}
	if p.LastBest > (*p.Chromosomes)[0].Score {
		p.LastBest = (*p.Chromosomes)[0].Score
		p.TiltCount = 0
		fmt.Printf("%d. Iterasyonda Sonuc : %.5f\n", p.Iteration, p.LastBest)
	}
}

func (p Population) ScoreResult() []float64 {
	total := make([]float64, len(*p.Chromosomes)+1)
	total[0] = 0.0
	i := 1
	for _, v := range *p.Chromosomes {
		total[i] = total[i-1] + v.Score
		i++
	}
	return total[1:]
}
