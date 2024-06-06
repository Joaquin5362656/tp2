package main

import (
	"bufio"
	"fmt"
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

		var informeError string

		comandoValido, informeError = comando.EsComandoValido()

		if comandoValido {
			comando.EjecutarComando(analisisLogs)
		} else {
			fmt.Fprintf(os.Stderr, "%s", informeError)
		}
	}

}
