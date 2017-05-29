package main

import(
	h "github.com/andreagonz/peces/heuristica"
	i "github.com/andreagonz/peces/implementacion"
	p "github.com/andreagonz/peces/prueba"
	"fmt"
)

func main() {
	// s := -5
	// c := []int{15, 22, 14, 26, 32, 9, 16, 8, -43, -30, 0, -1, 2, 23, -12, 4, -4, 33, -33}
	
	// c := p.Conjunto
	// s := 5385
	
	c := p.Conjunto2
	s := -910492
	
	itmax := 100
	tcardumen := 100
	stepi := 0.5
	thresc := 0.5
	thresv := 0.5
	seed := int64(4)	
	cparo := i.CondicionParo{0, itmax}
	crea := i.CreaConjunto{}
	fmt.Print("Conjunto inicial ")
	fmt.Println(c)
	fmt.Print("Tamaño de conjunto: ")
	fmt.Println(len(c))
	fmt.Print("Suma: ")
	fmt.Println(s)
	i.SetSuma(s)
	i.SetNumeros(c)
	i.MaxMin()
	h.BFSS(itmax, tcardumen, len(c), stepi, thresc, thresv, seed, &cparo, &crea)
}
