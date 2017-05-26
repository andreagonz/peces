type Conjunto struct {
	Numeros []int
	fitness float64
	peso float64
}

func (c Conjunto) ObtenFit() float64 {
	return c.fitness
}

func (c Conjunto) ObtenPeso() float64 {
	return c.peso
}

func (c *Conjunto) CalculaFitness() {
	c.fitness = 0
}
