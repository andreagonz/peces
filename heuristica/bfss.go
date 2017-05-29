package peces

import (
	"math/rand"
	"fmt"
	"math"
	"container/list"
)

type Pez interface {
	CambiaBool(int, bool) // Índice de bit a cambiar, MovimientoIndividual?. Debe actualizar fitness.
	ObtenBool(int) bool
	CalculaFitness()
	Fitness() float64
	DFitness() float64
	AsignaPeso(float64)
	Peso() float64
	Copia() Pez
	Str() string
}

type Cardumen struct {
	peces []Pez
	mejor Pez
	tvector int
	peso float64	
}

type Paro interface {
	Condicion() bool
}

type Crea interface {
	Pez([]bool) Pez
}

func InicializarCardumen(c * Cardumen, tcardumen int, crea Crea, r * rand.Rand) {
	(*c).peces = make([]Pez, tcardumen)
	for i := 0; i < tcardumen; i++ {
		v := make([]bool, (*c).tvector)
		for j := 0; j < (*c).tvector; j++ {
			in := r.Intn(2)
			if in == 1 {
				v[j] = true
			} 
		}
		p := crea.Pez(v)
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
	for x := 0; x < len((*c).peces); x++ {
		for y := 0; y < (*c).tvector; y++ {
			if r.Float64() < s {
				(*c).peces[x].CambiaBool(y, true)
			}
		}
		if (*c).peces[x].Fitness() > (*c).mejor.Fitness() {
			(*c).mejor = (*c).peces[x].Copia()
		}
	}
}

func ComparaVectores(c * Cardumen, v []bool, r * rand.Rand) {
	for i := 0; i < len((*c).peces); i++ {
		l := list.New()
		for j := 0; j < len(v); j++ {
			if v[j] != (*c).peces[i].ObtenBool(j) {
				l.PushBack(j)
			}
		}
		if l.Len() > 0 {
			g := r.Intn(l.Len())
			k := 0
			for e := l.Front(); e != nil && k < g; e = e.Next() {
				if k == g {
					(*c).peces[i].CambiaBool(e.Value.(int), false)
					if (*c).peces[i].Fitness() > (*c).mejor.Fitness() {
						(*c).mejor = (*c).peces[i].Copia()
					}
				}
				k++
			}
		}
	}
}

func MovColectivoInstintivo(c * Cardumen, thresc float64, r * rand.Rand) {
	v := make([]float64, (*c).tvector)
	df := 0.0
	for i := 0; i < len((*c).peces); i++ {
		for j := 0; j < (*c).tvector; j++ {
			if (*c).peces[i].ObtenBool(j) {
				v[j] += float64((*c).peces[i].DFitness())
			}
		}
		df += (*c).peces[i].DFitness()
	}
	max := 0.0
	for i := 0; i < len(v); i++ {
		v[i] /= df
		if v[i] > max {
			max = v[i]
		}
	}
	tmax := max * thresc
	iv := make([]bool, len(v))
	for i := 0; i < len(v); i++ {
		if v[i] > tmax {
			iv[i] = true
		}
	}
	ComparaVectores(c, iv, r)
}

func MovColectivoVolitivo(c * Cardumen, thresv float64,  r * rand.Rand) {
	b := make([]float64, (*c).tvector)
	w := 0.0
	for i := 0; i < len((*c).peces); i++ {
		for j := 0; j < (*c).tvector; j++ {
			if (*c).peces[i].ObtenBool(j) {
				b[j] += float64((*c).peces[i].Peso())
			}
		}
		w += (*c).peces[i].Peso()
	}
	max := 0.0
	for i := 0; i < len(b); i++ {
		b[i] /= w
		if b[i] > max {
			max = b[i]
		}
	}
	tmax := max * thresv
	iv := make([]bool, len(b))
	for i := 0; i < len(b); i++ {
		if b[i] > tmax {
			if w > (*c).peso {
				iv[i] = true
			}
		} else {
			if w < (*c).peso {
				iv[i] = true
			}
		}
	}
	(*c).peso = w
	ComparaVectores(c, iv, r)
}

func AlimentaPeces(c * Cardumen) {
	mfd := 0.0
	for x := 0; x < len((*c).peces); x++ {
		if mfd < math.Abs((*c).peces[x].Peso()) {
			mfd = math.Abs((*c).peces[x].Peso())
		}
	}
	for x := 0; x < len((*c).peces); x++ {
		w := (*c).peces[x].Peso() + (*c).peces[x].DFitness() / math.Abs(mfd)
		(*c).peces[x].AsignaPeso(w)
	}
}

func BFSS(itmax int, tcardumen int, tvector int, s float64, thresc float64, thresv float64, seed int64, paro Paro, crea Crea) {
	var c Cardumen
	c.tvector = tvector
	r := rand.New(rand.NewSource(seed))
	si := s
	tv := thresv
	InicializarCardumen(&c, tcardumen, crea, r)
	for paro.Condicion() {
		MovimientoIndividual(&c, si, r)
		AlimentaPeces(&c)
		MovColectivoInstintivo(&c, thresc, r)
		MovColectivoVolitivo(&c, tv, r)
		si -= s / float64(itmax)
		tv -= thresv / float64(itmax)
	}
	fmt.Print("Resultado ")
	fmt.Println(c.mejor.Str())
}
