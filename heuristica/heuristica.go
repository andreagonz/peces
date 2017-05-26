package peces

import (
	"math/rand"
)

type Pez interface {
	ObtenPeso() float64
	ObtenPosicion() float64
	Aleatorio(*rand.Rand) Pez
	CalculaFit()
	ObtenFitness() float64
}

type Cardumen interface {
	Peces []Pez
	Baricentro int
	CondicionParo() bool
}

func MovimientoIndividual(c * Cardumen) {
	a := 0
	b := 0
	for x := 0; x < len((*c).Peces) - 1; x++ {
		df := ((c*).Peces[x].Fitness - (c*).Peces[x + 1].Fitness)
		a += ((c*).Peces[x].Posicion - (c*).Peces[x + 1].Posicion) * df
		b += df
	}
	i := a / b
	for x := 0; x < len((*c).Peces); x++ {	
		(c*).Peces[x] += i
	}
}

func InicializarCardumen(c * Cardumen, t int, r * rand.Rand) {
	(*c).Peces = make([]Pez, t)
	for i := 0; i < t; i++ {
		p := Pez{}
		p = p.Aleatorio(r)
		p.CalculaFit()
		(*c).Peces[i] = p
	}
}

func FSS(seed int, t int) {
	c := Cardumen{}
	r := rand.New(rand.NewSource(seed))	
	InicializarCardumen(*c, t, r)
	for !c.CondicionParo() {
		
	}
}
