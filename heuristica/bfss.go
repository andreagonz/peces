package peces

import (
	"math/rand"
	"fmt"
	"math"
	"container/list"
	// g "github.com/andreagonz/peces/gui"
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
	Peces []Pez
	Mejor Pez
	Tvector int
	peso float64
	Iteracion int
	Itmax int
}

type Paro interface {
	Condicion(Cardumen) bool
}

type Crea interface {
	Pez([]bool) Pez
}

func InicializarCardumen(c * Cardumen, tcardumen int, crea Crea, r * rand.Rand) {
	(*c).Peces = make([]Pez, tcardumen)
	for i := 0; i < tcardumen; i++ {
		v := make([]bool, (*c).Tvector)
		for j := 0; j < (*c).Tvector; j++ {
			in := r.Intn(2)
			if in == 1 {
				v[j] = true
			} 
		}
		p := crea.Pez(v)
		p.CalculaFitness()
		(*c).Peces[i] = p
		if i == 0 {
			(*c).Mejor = p.Copia()
		} else {
			if (*c).Mejor.Fitness() < p.Fitness() {
				(*c).Mejor = p.Copia()
			}
		}
	}
}

func MovimientoIndividual(c * Cardumen, s float64,  p float64, r * rand.Rand) {
	for x := 0; x < len((*c).Peces); x++ {
		for y := 0; y < (*c).Tvector; y++ {
			if r.Float64() < s {
				if r.Float64() < p {
					(*c).Peces[x].CambiaBool(y, false)
				} else {
					(*c).Peces[x].CambiaBool(y, true)
				}
			}
		}
		if (*c).Peces[x].Fitness() > (*c).Mejor.Fitness() {
			(*c).Mejor = (*c).Peces[x].Copia()
		}
	}
}

func ComparaVectores(c * Cardumen, v []bool, r * rand.Rand) {
	for i := 0; i < len((*c).Peces); i++ {
		l := list.New()
		for j := 0; j < len(v); j++ {
			if v[j] != (*c).Peces[i].ObtenBool(j) {
				l.PushBack(j)
			}
		}
		if l.Len() > 0 {
			g := r.Intn(l.Len())
			k := 0
			for e := l.Front(); e != nil && k < g; e = e.Next() {
				if k == g {
					(*c).Peces[i].CambiaBool(e.Value.(int), false)
					if (*c).Peces[i].Fitness() > (*c).Mejor.Fitness() {
						(*c).Mejor = (*c).Peces[i].Copia()
					}
				}
				k++
			}
		}
	}
}

func MovColectivoInstintivo(c * Cardumen, thresc float64, r * rand.Rand) {
	v := make([]float64, (*c).Tvector)
	df := 0.0
	for i := 0; i < len((*c).Peces); i++ {
		for j := 0; j < (*c).Tvector; j++ {
			if (*c).Peces[i].ObtenBool(j) {
				v[j] += float64((*c).Peces[i].DFitness())
			}
		}
		df += (*c).Peces[i].DFitness()
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
	b := make([]float64, (*c).Tvector)
	w := 0.0
	for i := 0; i < len((*c).Peces); i++ {
		for j := 0; j < (*c).Tvector; j++ {
			if (*c).Peces[i].ObtenBool(j) {
				b[j] += float64((*c).Peces[i].Peso())
			}
		}
		w += (*c).Peces[i].Peso()
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
	for x := 0; x < len((*c).Peces); x++ {
		if mfd < math.Abs((*c).Peces[x].Peso()) {
			mfd = math.Abs((*c).Peces[x].Peso())
		}
	}
	for x := 0; x < len((*c).Peces); x++ {
		w := (*c).Peces[x].Peso() + (*c).Peces[x].DFitness() / math.Abs(mfd)
		(*c).Peces[x].AsignaPeso(w)
	}
}

func BFSS(itmax int, tcardumen int, tvector int, s float64, pind float64, thresc float64, thresv float64, seed int64, paro Paro, crea Crea) {
	var c Cardumen
	c.Tvector = tvector
	c.Itmax = itmax
	c.Iteracion = 0
	r := rand.New(rand.NewSource(seed))
	si := s
	tv := thresv
	InicializarCardumen(&c, tcardumen, crea, r)
	for paro.Condicion(c) {
		MovimientoIndividual(&c, si, pind, r)
		AlimentaPeces(&c)
		MovColectivoInstintivo(&c, thresc, r)
		MovColectivoVolitivo(&c, tv, r)
		si -= s / float64(itmax)
		tv -= thresv / float64(itmax)
		c.Iteracion++
	}
	fmt.Print("\nResultado ")
	fmt.Println(c.Mejor.Str())
}
