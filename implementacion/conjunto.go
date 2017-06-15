package peces

import (
	h "github.com/andreagonz/peces/heuristica"
	"math"
	"strconv"
)

var numeros []int
var suma, maxDif int

type Conjunto struct {
	Vector []bool
	fitness float64
	peso float64
	dfitness float64
	Suma int
	Tamanio int
}

type CondicionParo struct {
}

type CreaConjunto struct {
}

func (c * CreaConjunto) Pez(v []bool) h.Pez {
	t := 0
	for i := 0; i < len(v); i++ {
		if v[i] {
			t++
		}
	}
	return &Conjunto{Vector : v, Tamanio: t}
}

func (cp * CondicionParo) Condicion(c h.Cardumen) bool {
	if c.Mejor.(*Conjunto).Suma == suma ||
		c.Iteracion > c.Itmax {
		return false
	}
	return true
}

func SetNumeros(n []int) {
	numeros = n
}

func SetSuma(n int) {
	suma = n
}

func MaxMin() {
	smin := 0
	smax := 0
	mini := math.MaxInt32
	maxi := math.MinInt32
	for i := 0; i < len(numeros); i++ {
		if numeros[i] < 0 {
			smin += numeros[i]
		} else {
			smax += numeros[i]
		}
		if numeros[i] < mini {
			mini = numeros[i]
		}
		if numeros[i] > maxi {
			maxi = numeros[i]
		}
	}
	dmax := math.Abs(float64(smax - suma))
	dmin := math.Abs(float64(smin - suma))
	maxDif = int(math.Max(dmax, dmin))
}

func (c * Conjunto) CambiaBool(i int, b bool) {
	nf := 0.0
	c.Vector[i] = !c.Vector[i]
	if c.Vector[i] {
		c.Tamanio++
		c.Suma += numeros[i]
	} else {
		c.Tamanio--
		c.Suma -= numeros[i]
	}	

	dif := math.Abs(float64(c.Suma - suma))
	dif = dif / float64(maxDif)
	nf = 1 - dif
	
	if b {
		if nf > c.fitness {
			c.fitness = nf
			c.dfitness = nf - c.fitness
		} else {
			c.dfitness = 0.0
			c.Vector[i] = !c.Vector[i]
			if c.Vector[i] {
				c.Tamanio++
				c.Suma += numeros[i]
			} else {
				c.Tamanio--
				c.Suma -= numeros[i]
			}
		}
	} else {
		c.fitness = nf
	}
}

func (c Conjunto) ObtenBool(i int) bool {
	return c.Vector[i]
}

func (c * Conjunto) CalculaFitness() {
	s := 0
	for i := 0; i < len(c.Vector); i++ {
		if c.Vector[i] {
			s += numeros[i]
		}
	}
	c.Suma = s
	dif := math.Abs(float64(s - suma))
	dif = dif / float64(maxDif)
	c.fitness = 1 - dif
}

func (c Conjunto) Fitness() float64 {
	return c.fitness
}

func (c Conjunto) DFitness() float64 {
	return c.dfitness
}

func (c * Conjunto) AsignaPeso(p float64) {
	c.peso = p
}

func (c Conjunto) Peso() float64 {
	return c.peso
}

func (c Conjunto) Copia() h.Pez {
	v := make([]bool, len(c.Vector))
	for i := 0; i < len(c.Vector); i++ {
		v[i] = c.Vector[i]
	}
	nuevo := Conjunto{v, c.fitness, c.peso, c.dfitness, c.Suma, c.Tamanio}
	return &nuevo
}

func (c Conjunto) Str() string {
	return "\nMejor suma encontrada: " + strconv.Itoa(c.Suma) + "\nTamaño de conjunto: " + strconv.Itoa(c.Tamanio) + "\nFitness: " + strconv.FormatFloat(c.fitness, 'f', -1, 64)
}
