package peces

import (
	"math/rand"
	"math"
	"container/list"
)

type Pez interface {
	AsignaVector([]bool)
	CambiaBool(int, bool) // Índice de bit a cambiar, MovimientoIndividual?. Debe actualizar fitness.
	ObtenBool(int) bool
	CalculaFitness()
	Fitness() float64
	DFitness() float64
	Copia()
	Str() string
}

type Cardumen struct {
	peces []Pez
	mejor Pez
	tvector int
}

func InicializarCardumen(c * Cardumen, tcardumen int, r * rand.Rand) {
	(*c).peces = make([]Pez, tcardumen)
	for i := 0; i < tcardumen; i++ {
		v = make([]bool, (*c).tvector)
		for j := 0; j < (*c).tvector; j++ {
			in := r.Intn(2)
			if in == 1 {
				v[j] = true
			} 
		}
		var p Pez
		p.AsignaVector()
		p.CalculaFitness()
		(*c).peces[i] = p
		if i == 0 {
			(*c).mejor = p.Copia()
		} else {
			if (*c).mejor.Fitness() < p.Fitness() {
				(*c).mejor = p.Copia()
			}
		}
	}
}

func MovimientoIndividual(c * Cardumen, s float64,  r * rand.Rand) {
	for x := 0; x < len((*c).Peces()); x++ {
		if r.Float64() < s {
			(*c).peces[x].CambiaBool(x, true)
			if (*c).peces[x].Fitness() > mejor.Fitness() {
				mejor = (*c).peces[x].Copia()
			}
		}
	}
}

func ComparaVectores(c * Cardumen, v []bool) {
	for i := 0; i < len((*c).peces); i++ {
		l := list.New()
		for j := 0; j < len(v); j++ {
			if v[j] != (*c).peces[i].ObtenBool(j) {
				l.PushBack(j)
			}
		}
		g := r.Intn(l.Len())
		for e := l.Front(), k := 0; e != nil, k < g; e = e.Next(), k++ {
			if k == g {
				(*c).peces[i].CambiaBool(e, false)
				if (*c).peces[i].Fitness() > (*c).mejor.Fitness() {
					(*c).mejor = (*c).peces[i].Copia()
				}
			}
		}
	}
}

func MovColectivoInstintivo(c * Cardumen, thresc float64, r * rand.Rand) {
	v := make(float64, len((*c).tvector))
	df := 0.0
	for i := 0; i < len((*c).peces); i++ {
		for j := 0; j < len((*c).tvector); j++ {
			if (*c).peces[i].ObtenBool(j) {
				v[j] += float64((*c).peces[i].DFitness())
			}
		}
		df += (*c).peces.DFitness()
	}
	max := 0.0
	for i := 0; i < len(v); i++ {
		v[i] /= df
		if v[i] > max {
			max = v[i]
		}
	}
	tmax := max * thresc
	iv := make(bool, len(v))
	for i := 0; i < len(v); i++ {
		if v[i] > tmax {
			iv[i] = true
		}
	}
	ComparaVectores(c, iv)
}

func MovColectivoVolitivo(c * Cardumen, r * rand.Rand) {
	(*c).CalculaBaricentro()
	(*c).MovVolitivo(r)
}

func AlimentaPeces(c * Cardumen) {
	mfd := 0.0
	for x := 0; x < len((*c).Peces()); x++ {
		if mfd < math.Abs((*c).Peces()[x].Peso()) {
			mfd = math.Abs((*c).Peces()[x].Peso())
		}
	}	
	for x := 0; x < len((*c).Peces()); x++ {
		w := (*c).Peces()[x].Peso() + (*c).Peces()[x].FitnessDif() / math.Abs(mfd)
		(*c).Peces()[x].AsignaPeso(w)
	}
}

func FSS(itmax int, tcardumen int, tvector int, stepi float64, stepv float64, seed int64) {
	var c Cardumen
	c.tvector = tvector
	r := rand.New(rand.NewSource(seed))
	si := stepi
	sv := stepv
	InicializarCardumen(&c, tcardumen, r)
	for i := 0; i < itmax; i++ {
		MovimientoIndividual(&c, si, r)
		AlimentaPeces(&c)
		MovColectivoInstintivo(&c, r)
		MovColectivoVolitivo(&c, r)
		si -= stepi / float64(itmax)
		sv -= stepv / float64(itmax)
	}
}
