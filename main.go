package main

import(
	h "github.com/andreagonz/peces/heuristica"
	i "github.com/andreagonz/peces/implementacion"
)

func main() {
	s := 53
	c := []int{15, 22, 14, 26, 32, 9, 16, 8}
	itmax := 500
	tcardumen := 100
	stepi := 0.5
	thresc := 0.5
	thresv := 0.5
	seed := int64(3)	
	cparo := i.CondicionParo{0, itmax}
	crea := i.CreaConjunto{}
	i.SetSuma(s)
	i.SetNumeros(c)
	i.MaxMin()
	h.BFSS(itmax, tcardumen, len(c), stepi, thresc, thresv, seed, &cparo, &crea)
}
