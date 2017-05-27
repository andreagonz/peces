package peces

import (
	"math/rand"
	"math"
)

type Pez interface {
	ObtenPeso() float64
	AsignaPeso(float64)
	ObtenPosicionDif() float64
	ObtenFitnessDif() float64
	Mover(float64, bool, * rand.Rand) // Variable de movimiento, Elitismo?, Generador aleatorio
	Str()
}

type Cardumen interface {
 	ObtenPeces() []Pez
	InicializarCardumen(int, *rand.Rand)
	Mejor() Pez
	GuardarMejor(Pez)
	CalculaBaricentro()
	MovVolitivo(* rand.Rand)
}

func MovimientoIndividual(c * Cardumen, s float64,  r * rand.Rand) {
	for x := 0; x < len((*c).ObtenPeces()); x++ {
		mov := (r.Float64() * 2 - 1) * s
		(*c).ObtenPeces()[x].Mover(mov, true, r)
	}
}

func MovColectivoInstintivo(c * Cardumen, r * rand.Rand) {
	a := 0.0
	b := 0.0
	for x := 0; x < len((*c).ObtenPeces()); x++ {
		df := (*c).ObtenPeces()[x].ObtenPosicionDif() * (*c).ObtenPeces()[x].ObtenFitnessDif()
		a += (*c).ObtenPeces()[x].ObtenFitnessDif()
		b += df
	}
	i := a / b
	for x := 0; x < len((*c).ObtenPeces()); x++ {	
		(*c).ObtenPeces()[x].Mover(i, false, r)
	}
}

func MovColectivoVolitivo(c * Cardumen, r * rand.Rand) {
	(*c).CalculaBaricentro()
	(*c).MovVolitivo(r)
}

func AlimentaPeces(c * Cardumen) {
	mfd := 0.0
	for x := 0; x < len((*c).ObtenPeces()); x++ {
		if mfd < math.Abs((*c).ObtenPeces()[x].ObtenPeso()) {
			mfd = math.Abs((*c).ObtenPeces()[x].ObtenPeso())
		}
	}	
	for x := 0; x < len((*c).ObtenPeces()); x++ {
		w := (*c).ObtenPeces()[x].ObtenPeso() + (*c).ObtenPeces()[x].ObtenFitnessDif() / math.Abs(mfd)
		(*c).ObtenPeces()[x].AsignaPeso(w)
	}
}

func FSS(itmax int, t int, stepi float64, stepv float64, seed int64) {
	var c Cardumen
	r := rand.New(rand.NewSource(seed))
	si := stepi
	sv := stepv
	c.InicializarCardumen(t, r)
	for i := 0; i < itmax; i++ {
		MovimientoIndividual(&c, stepi, r)
		AlimentaPeces(&c)
		MovColectivoInstintivo(&c, r)
		MovColectivoVolitivo(&c, r)
		si -= stepi / float64(itmax)
		sv -= stepv / float64(itmax)
	}
}
