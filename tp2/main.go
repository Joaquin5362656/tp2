package tp2

import (
	"bufio"
	"os"
	Analisis "tp2/tp2/analisis_log"
	LecturaComandos "tp2/tp2/comandos"
)

func main() {

	lectura := bufio.NewScanner(os.Stdin)

	comandoValido := true

	analisisLogs := Analisis.GenerarDatos()

	for comandoValido && lectura.Scan() {

		comando := LecturaComandos.CargarComando(lectura)

		comandoValido, informeError := comando.EsComandoValido()

		if comandoValido {
			comando.EjecutarComando(analisisLogs)
		} else {
			print(informeError)
		}
	}

}
