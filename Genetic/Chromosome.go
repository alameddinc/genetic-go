package Genetic

import (
	"math"
	"math/rand"
	"time"
)

type Chromosome struct {
	Genes *[]Gene
	Score float64
}

func (c *Chromosome) Init(g []Gene, first Gene) {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	tmp := g
	firstGen := first
	random.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})
	rtmp := append(tmp, firstGen)
	for i, j := 0, len(rtmp)-1; i < j; i, j = i+1, j-1 {
		rtmp[i], rtmp[j] = rtmp[j], rtmp[i]
	}
	c.Genes = &rtmp
	c.FitnessFunction()
}

func (c *Chromosome) FitnessFunction() {
	c.Score = 0
	for i := 0; i < len(*c.Genes)-1; i++ {
		X := math.Pow((*c.Genes)[i].X-(*c.Genes)[i+1].X, 2)
		Y := math.Pow((*c.Genes)[i].Y-(*c.Genes)[i+1].Y, 2)
		c.Score += math.Sqrt(X + Y)
	}
}

func (c Chromosome) Crossover(c2 Chromosome) (Chromosome, Chromosome) {
	count := len(*c.Genes)
	var tmpC, tmpC2 Chromosome
	tmpC.Genes = &[]Gene{}
	tmpC2.Genes = &[]Gene{}
	*tmpC.Genes = append(*tmpC.Genes, (*c.Genes)[:count/2]...)
	*tmpC2.Genes = append(*tmpC2.Genes, (*c2.Genes)[:count/2]...)
	for i := 0; i < count; i++ {
		if !tmpC.isExist((*c2.Genes)[i]) {
			*tmpC.Genes = append(*tmpC.Genes, (*c2.Genes)[i])
		}
		if !tmpC2.isExist((*c.Genes)[i]) {
			*tmpC2.Genes = append(*tmpC2.Genes, (*c.Genes)[i])
		}
	}

	tmpC.FitnessFunction()
	tmpC2.FitnessFunction()
	return tmpC, tmpC2
}

func (c Chromosome) Mutation(rate float64) Chromosome {
	count := len(*c.Genes)
	number := int(float64(count) * rate)
	if number == 0 {
		number = 1
	}
	tmp := Chromosome{}
	tmp = c
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < number; i++ {
		in, out := c.RandomGeneID(r, count), c.RandomGeneID(r, count)
		(*tmp.Genes)[in], (*tmp.Genes)[out] = (*tmp.Genes)[out], (*tmp.Genes)[in]
	}
	tmp.FitnessFunction()
	return tmp
}

func (c *Chromosome) isExist(gene Gene) bool {
	for _, v := range *c.Genes {
		if v.Name == gene.Name {
			return true
		}
	}
	return false
}

func (c *Chromosome) RandomGeneID(r *rand.Rand, count int) int {
	return (r.Int() % (count - 1)) + 1
}
