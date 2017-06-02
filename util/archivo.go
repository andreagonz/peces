package peces

import (
	"io/ioutil"
	"strings"
	"strconv"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// EscribeArchivo recibe una cadena y crea un archivo donde la escribe.
func EscribeArchivo(s string, nombre string) {
	d1 := []byte(s)
	err := ioutil.WriteFile(nombre, d1, 0644)
	check(err)
}

// LeeArchivo lee un archivo y lo regresa como cadena.
func LeeArchivo(nom string) string {
	if _, err := os.Stat(nom); os.IsNotExist(err) {
		return ""
	}
	dat, err := ioutil.ReadFile(nom)
	check(err)
	return string(dat)
}

// Entrada recibe un conjunto de numeros en forma de cadena y la 
// regresa como un arreglo de enteros.
func LeeConjunto(s string) (int, []int) {
	r := LeeArchivo(s)
	r = strings.Replace(r, ",", " ", -1)
	r = strings.Replace(r, "\n", " ", -1)
	l := strings.Fields(r)
	res := make([]int, len(l) - 1)
	for i := 1; i < len(l); i++ {
		ind, err := strconv.Atoi(l[i])
		check(err)
		res[i - 1] = ind
	}
	sum, err := strconv.Atoi(l[0])
	check(err)
	return sum, res
}

// Entrada recibe los parametros en forma de cadena y los
// regresa como un arreglo.
func LeeParametros(s string) []float64 {
	r := LeeArchivo(s)
	r = strings.Replace(r, "\n", " ", -1)
	l := strings.Fields(r)
	res := make([]float64, len(l))
	for i := 0; i < len(l); i++ {
		ind, err := strconv.ParseFloat(l[i], 64)
		check(err)
		res[i] = ind
	}
	return res
}
