package tp2

import (
	"bufio"
	"os"
	LecturaComandos "tp2/tp2/comandos"
)

func main() {

	lectura := bufio.NewScanner(os.Stdin)

	comandoValido := true

	for comandoValido && lectura.Scan() {

		comando := LecturaComandos.CargarComando(lectura)

		comandoValido, informeError := comando.EsComandoValido()

		if comandoValido {
			comando.EjecutarComando()
		} else {
			print(informeError)
		}
	}

}
