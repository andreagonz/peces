type Conjunto struct {
	vector []bool
	fitness float64
	peso float64
	suma int
}

type Conjuntos struct {
	peces []Conjunto
	numeros []int
	baricentro int
	mejor Conjunto
	n int
	iteracion int
}

func (c Conjunto) ObtenFitness() float64 {
	return c.fitness
}

func (c Conjunto) ObtenPeso() float64 {
	return c.peso
}

func (c * Conjunto) CalculaFitness() {
	c.fitness = 0
}

func (c Conjuntos) ObtenBaricentro() {
	return c.baricentro
}

func (c Conjuntos) ObtenPeces() {
	return c.peces
}

func (c Conjuntos) CondicionParo() {
	b := c.mejor.suma == c.n
	b = b || c.iteracion == 10000
	return b
}

func (c * Conjuntos) InicializarCardumen(t int, r * rand.Rand) {
	(*c).Peces = make([]Pez, t)
	for i := 0; i < t; i++ {
		v = make([]bool, t)
		for j := 0; j < t; j++ {
			in := r.Intn(2)
			if in == 0 {
				v[j] = false
			} else v[j] = true
		}
		p := Conjunto{v}
		p.CalculaFitness()
		(*c).peces[i] = p
	}
	CalculaBaricentro(c)
}
