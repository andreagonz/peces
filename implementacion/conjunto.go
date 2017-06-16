package peces

import (
	h "github.com/andreagonz/peces/heuristica"
	"math"
	"bytes"
	"strconv"
)

var numeros []int
var suma, maxDif int

// Conjunto representa un subconjunto del conjunto principal
type Conjunto struct {
	Vector []bool
	fitness float64
	peso float64
	dfitness float64
	Suma int
	Tamanio int
}

// CondicionParo es la implementación de la condición de paro
type CondicionParo struct {
}

// CreaConjunto es la implementación para crear un conjunto
type CreaConjunto struct {
}

// Pez crea una solución aleatoria
func (c * CreaConjunto) Pez(v []bool) h.Pez {
	t := 0
	for i := 0; i < len(v); i++ {
		if v[i] {
			t++
		}
	}
	return &Conjunto{Vector : v, Tamanio: t}
}

// Condición regresa el resultado de la condición de paro
func (cp * CondicionParo) Condicion(c h.Cardumen) bool {
	if c.Mejor.(*Conjunto).Suma == suma ||
		c.Iteracion > c.Itmax {
		return false
	}
	return true
}

// SetNumeros asigna el arreglo de números que representa la instancia
// del problema a la variable global numeros
func SetNumeros(n []int) {
	numeros = n
}

// SetSuma asigna a la variable global suma el valor de la suma a buscar
func SetSuma(n int) {
	suma = n
}

// MaxMin encuentra la diferencia máxima de entre las posibles sumas y la suma a buscar
func MaxMin() {
	smin := 0
	smax := 0
	for i := 0; i < len(numeros); i++ {
		if numeros[i] < 0 {
			smin += numeros[i]
		} else {
			smax += numeros[i]
		}
	}
	dmax := math.Abs(float64(smax - suma))
	dmin := math.Abs(float64(smin - suma))
	maxDif = int(math.Max(dmax, dmin))
}

// CambiaBool cambia el bool en la posición i del vector de c y recalcula el fitness
// si b es true se usa elitismo, si es false no se usa elitismo
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
	if c.Tamanio == 0 {
		c.CambiaBool(i, false);
	}

}

// ObtenBool obtiene el bool en la posición i del vector de c
func (c Conjunto) ObtenBool(i int) bool {
	return c.Vector[i]
}

// CalculaFitness calcula el fitness de c
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

// Fitness regresa el fitness de c
func (c Conjunto) Fitness() float64 {
	return c.fitness
}

// DFitness regresa la diferencia entre el fitness de la iteración anterior y la actual de c
func (c Conjunto) DFitness() float64 {
	return c.dfitness
}

// AsignaPeso asigna el peso p a c
func (c * Conjunto) AsignaPeso(p float64) {
	c.peso = p
}

// Peso regresa el peso de c
func (c Conjunto) Peso() float64 {
	return c.peso
}

// Copia crea un conjunto copia de c y lo regresa
func (c Conjunto) Copia() h.Pez {
	v := make([]bool, len(c.Vector))
	for i := 0; i < len(c.Vector); i++ {
		v[i] = c.Vector[i]
	}
	nuevo := Conjunto{v, c.fitness, c.peso, c.dfitness, c.Suma, c.Tamanio}
	return &nuevo
}

// Str regresa una representación en cadena de c
// si b es true, la cadena incluye el vector
// si b es false la cadena no incluye al vector
func (c Conjunto) Str(b bool) string {
	if b {
		j := 0
		var buffer bytes.Buffer
		buffer.WriteString("Suma: " + strconv.Itoa(c.Suma) + ", Tamaño: " + strconv.Itoa(c.Tamanio) + ", Fitness: " + strconv.FormatFloat(c.fitness, 'f', -1, 64) + "\n")
		for i := 0; i < len(c.Vector); i++ {
			if c.Vector[i] {
				s := strconv.Itoa(numeros[i])
				if j < c.Tamanio - 1 {
					s += ", "
				}
				buffer.WriteString(s)
				j++
			}
		}
		return buffer.String()
	}
	return "\nMejor suma encontrada: " + strconv.Itoa(c.Suma) + "\nTamaño de conjunto: " + strconv.Itoa(c.Tamanio) + "\nFitness: " + strconv.FormatFloat(c.fitness, 'f', -1, 64)
}
