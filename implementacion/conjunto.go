package peces

import (
	h "github.com/andreagonz/peces/heuristica"
	"math"
	"bytes"
	"strconv"
)

var numeros []int
var suma, maxDif int

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

func (cp * CondicionParo) Condicion(c h.Cardumen) bool {
	if c.Mejor.(*Conjunto).suma == suma ||
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
	c.vector[i] = !c.vector[i]
	if c.vector[i] {
		c.tamanio++
		c.suma += numeros[i]
	} else {
		c.tamanio--
		c.suma -= numeros[i]
	}	

	dif := math.Abs(float64(c.suma - suma))
	dif = dif / float64(maxDif)
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
	v := make([]bool, len(c.vector))
	for i := 0; i < len(c.vector); i++ {
		v[i] = c.vector[i]
	}
	nuevo := Conjunto{v, c.fitness, c.peso, c.dfitness, c.suma, c.tamanio}
	return &nuevo
}

func (c Conjunto) Str() string {
	j := 0
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i := 0; i < len(c.vector); i++ {
		if c.vector[i] {
			s := strconv.Itoa(numeros[i])
			if j < c.tamanio - 1 {
				s += " "
			}
			buffer.WriteString(s)
			j++
		}
	}
	buffer.WriteString("] suma: " + strconv.Itoa(c.suma) + ", tamaño: " + strconv.Itoa(c.tamanio))
	return "\nMejor suma encontrada: " + strconv.Itoa(c.suma) + "\nTamaño de conjunto: " + strconv.Itoa(c.tamanio) + "\nFitness: " + strconv.FormatFloat(c.fitness, 'f', -1, 64)
}
