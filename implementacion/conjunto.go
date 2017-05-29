package peces

import (
	h "github.com/andreagonz/peces/heuristica"
	"math"
	"strconv"
)

type Conjunto struct {
	vector []bool
	fitness float64
	peso float64
	dfitness float64
	suma int
	tamanio int
}

type CondicionParo struct {
	Iteracion int
	Itmax int
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
	return &Conjunto{vector : v, tamanio: t}
}

func (c * CondicionParo) Condicion() bool {
	//danger
	// if mejorSuma == suma ||
	// c.Iteracion > c.Itmax {
	if c.Iteracion > c.Itmax {
		return false
	}
	c.Iteracion++
	return true
}

var numeros []int
var suma, max, min, mejorSuma int

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
	min = int(math.Min(float64(smin), float64(mini)))
	max = int(math.Max(float64(smax), float64(maxi)))
}

func (c * Conjunto) CambiaBool(i int, b bool) {
	nf := 0.0
	c.vector[i] = !c.vector[i]
	if c.vector[i] {
		c.tamanio++
		c.suma += numeros[i]
	} else {
		c.tamanio--
		c.suma -= numeros[i]
	}	

	dif := math.Abs(float64(c.suma - suma))
	dif = (dif - float64(min)) / float64((max - min))	
	nf = 1 - dif
	
	if b {
		if nf > c.fitness {
			c.fitness = nf
			c.dfitness = nf - c.fitness
		} else {
			c.dfitness = 0.0
			c.vector[i] = !c.vector[i]
			if c.vector[i] {
				c.tamanio++
				c.suma += numeros[i]
			} else {
				c.tamanio--
				c.suma -= numeros[i]
			}
		}
	} else {
		c.fitness = nf
	}

	if math.Abs(float64(c.suma - suma)) < float64(mejorSuma) {
		mejorSuma = c.suma
	}
}

func (c Conjunto) ObtenBool(i int) bool {
	return c.vector[i]
}

func (c * Conjunto) CalculaFitness() {
	s := 0
	for i := 0; i < len(c.vector); i++ {
		if c.vector[i] {
			s += numeros[i]
		}
	}
	c.suma = s
	dif := math.Abs(float64(s - suma))
	dif = (dif - float64(min)) / float64((max - min))
	c.fitness = 1 - dif
	if math.Abs(float64(c.suma - suma))  < float64(mejorSuma) {
		mejorSuma = c.suma
	}
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
	v := make([]bool, len(c.vector))
	for i := 0; i < len(c.vector); i++ {
		v[i] = c.vector[i]
	}
	nuevo := Conjunto{v, c.fitness, c.peso, c.dfitness, c.suma, c.tamanio}
	return &nuevo
}

func (c Conjunto) Str() string {
	s := "["
	j := 0
	for i := 0; i < len(c.vector); i++ {
		if c.vector[i] {
			s += strconv.Itoa(numeros[i])
			if j < c.tamanio - 1 {
				s += " "
			}
			j++
		}
	}
	s += "] suma: " + strconv.Itoa(c.suma) + ", tamaño: " + strconv.Itoa(c.tamanio)
	return s
}
