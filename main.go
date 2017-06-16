package main

import(
	h "github.com/andreagonz/peces/heuristica"
	i "github.com/andreagonz/peces/implementacion"
	u "github.com/andreagonz/peces/util"
	"fmt"
	"os"
	"flag"
	"math"
	"math/rand"
	"github.com/pkg/errors"
	"encoding/json"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron/bootstrap"
	"github.com/asticode/go-astilog"
)

//go:generate go-bindata -pkg $GOPACKAGE -o resources.go resources/...
func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Uso: ./peces <archivo.ss> <params.txt>")
	} else {
		bgui := false
		suma, conjunto := u.LeeConjunto(args[0])
		params := u.LeeParametros(args[1])
		for x := 0; x < len(args); x++ {
			if args[x] == "-gui" {
				bgui = true
			}
		}

		seed := int64(params[0])
		itmax := int(params[1])
		tcardumen := int(params[2])
		stepi := params[3]
		pind := params[4]
		thresc := params[5]
		thresv := params[6]
		r := rand.New(rand.NewSource(seed))
		
		cparo := i.CondicionParo{}
		crea := i.CreaConjunto{}

		i.SetSuma(suma)
		i.SetNumeros(conjunto)
		i.MaxMin()

		// Si no se va a utilizar interfaz gráfica
		if !bgui {
			fmt.Print("Tamaño de conjunto inicial: ")
			fmt.Println(len(conjunto))
			fmt.Print("Suma a buscar: ")
			fmt.Println(suma)
			h.BFSS(itmax, tcardumen, len(conjunto), stepi, pind, thresc, thresv, seed, &cparo, &crea)
		} else {
			// Se crea la interfaz gráfica
			flag.Parse()
			astilog.SetLogger(astilog.New(astilog.FlagConfig()))
			
			if err := bootstrap.Run(bootstrap.Options{
				AstilectronOptions: astilectron.Options{
					AppName:            "BFSS-SS",
				},
				AdaptWindow: func(w *astilectron.Window) {
					if er := w.Maximize(); er != nil {
						astilog.Fatal(errors.Wrap(er, "minimizing window failed"))
					}
				},
				CustomProvision: func(baseDirectoryPath string) error {
					return nil
				},
				Homepage: "index.html",
				MessageHandler: func(w *astilectron.Window, m bootstrap.MessageIn) {
					switch m.Name {
					case "ready":
						type P struct {
							Tconjunto int `json:"tconjunto"`
							Suma int `json:"suma"`
							Npeces int `json:"npeces"`
						}
						p := P{len(conjunto), suma, tcardumen}
						if err := json.Unmarshal(m.Payload, &p); err != nil {
							astilog.Errorf("Unmarshaling %s failed", m.Payload)
							return
						}
						w.Send(bootstrap.MessageOut{Name: "ready", Payload: p})
						break
						
					case "empezar":
						type P struct {
							Iteracion int `json:"iteracion"`
							Msuma int `json:"msuma"`
							Mfitness float64 `json:"mfitness"`
							Peces []float64 `json:"peces"`
							Subconjunto []bool `json:"subconjunto"`
							Subtamanio int `json:"subtamanio"`
							Diferencia int `json:"diferencia"`
						}
						var c h.Cardumen
						c.Tvector = len(conjunto)
						c.Itmax = itmax
						c.Iteracion = 0
						si := stepi
						tv := thresv
						p := P{c.Iteracion, 0, 0.0, make([]float64, tcardumen), make([]bool, c.Tvector), 0, 0}
						h.InicializarCardumen(&c, tcardumen, &crea, r)
						for cparo.Condicion(c) {
							h.MovimientoIndividual(&c, si, pind, r)
							h.AlimentaPeces(&c)
							h.MovColectivoInstintivo(&c, thresc, r)
							h.MovColectivoVolitivo(&c, tv, r)
							si -= stepi / float64(itmax)
							tv -= thresv / float64(itmax)
							c.Iteracion++
							p.Iteracion++
							p.Msuma = c.Mejor.(*i.Conjunto).Suma
							p.Mfitness = c.Mejor.Fitness()
							for x := 0; x < len(c.Peces); x++ {
								p.Peces[x] = (1.0 - c.Peces[x].Fitness()) * 100
							}
							p.Subconjunto = c.Mejor.(*i.Conjunto).Vector
							p.Subtamanio = c.Mejor.(*i.Conjunto).Tamanio
							p.Diferencia = int(math.Abs(float64(p.Msuma - suma)))
							w.Send(bootstrap.MessageOut{Name: "iteracion", Payload: p})
						}
						w.Send(bootstrap.MessageOut{Name: "terminado", Payload: nil})
						u.EscribeArchivo(c.Mejor.Str(true), "subconjunto.res")
						break
					}
				},
				RestoreAssets: RestoreAssets,
				WindowOptions: &astilectron.WindowOptions{
					Center: astilectron.PtrBool(true),
					Height: astilectron.PtrInt(800),
					Width:  astilectron.PtrInt(1200),
				},
			}); err != nil {
				astilog.Fatal(err)
			}
		}
	}
}
