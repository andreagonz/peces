package peces

import (
	"math/rand"
)

type Pez interface {
	ObtenPeso() float64
	ObtenPosicion() float64
	CalculaFitness()
	ObtenFitness() float64
	AsignarFitness()
}

type Cardumen interface {
 	ObtenPeces() []Pez
	ObtenBaricentro() int
	CondicionParo() bool
	InicializarCardumen()
}

func MovimientoIndividual(c * Cardumen) {
	a := 0
	b := 0
	for x := 0; x < len((*c).Peces) - 1; x++ {
		df := (*c).Peces[x].Fitness - (c*).Peces[x + 1].ObtenFitness()
		a += (*c).Peces[x].ObtenPosicion() - (c*).Peces[x + 1].ObtenPosicion()) * df
		b += df
	}
	i := a / b
	for x := 0; x < len((*c).ObtenPeces); x++ {	
		pez := (c*).ObtenPeces[x]
		pez.AsignarFitness(pez.ObtenFitness() + i)
	}
}

func CalculaBaricentro(c * Cardumen) {
	c.baricentro = 0
}

func AlimentaPeces(c * Cardumen) {
}

func MovColectivoIns(c * Cardumen) {
}

func MovColectivoVol(c * Cardumen) {
	
}

func FSS(seed int, t int) {
	c := Cardumen{}
	r := rand.New(rand.NewSource(seed))	
	InicializarCardumen(&c, t, r)
	for !c.CondicionParo() {
		MovimientoIndividual(&c)
		AlimentaPeces(&c)
		MovColectivoIns(&c)
		MovColectivoVol(&c)
	}
}
