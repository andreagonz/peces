package main

import(
	h "github.com/andreagonz/peces/heuristica"
	i "github.com/andreagonz/peces/implementacion"
	u "github.com/andreagonz/peces/util"
	// p "github.com/andreagonz/peces/prueba"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Uso: ./peces <archivo.ss> <params.txt>")
	} else {
		suma, conjunto := u.LeeConjunto(args[0])
		params := u.LeeParametros(args[1])

		seed := int64(params[0])
		itmax := int(params[1])
		tcardumen := int(params[2])
		stepi := params[3]
		thresc := params[4]
		thresv := params[5]
		
		cparo := i.CondicionParo{}
		crea := i.CreaConjunto{}
		fmt.Print("Tamaño de conjunto inicial: ")
		fmt.Println(len(conjunto))
		fmt.Print("Suma a buscar: ")
		fmt.Println(suma)
		i.SetSuma(suma)
		i.SetNumeros(conjunto)
		i.MaxMin()
		h.BFSS(itmax, tcardumen, len(conjunto), stepi, thresc, thresv, seed, &cparo, &crea)
	}
}
