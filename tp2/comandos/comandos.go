package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	Analisis "tp2/tp2/analisis_log"
)

const (
	comando1 = "agregar_archivo"
	comando2 = "ver_visitantes"
	comando3 = "ver_mas_visitados"
)

type ComandosLog []string

func CargarComando(lecturaEntrada *bufio.Scanner) (commandoCargado ComandosLog) {

	logALeer := lecturaEntrada.Text()

	return strings.Fields(logALeer)

}

func (lectura ComandosLog) obtenerComando() (comandoLeido string) {
	return lectura[0]
}

func (lectura ComandosLog) obtenerParametros() []string {
	return lectura[1:]
}

func (lectura ComandosLog) EsComandoValido() (esValido bool, informeError string) {

	switch lectura.obtenerComando() {
	case comando1:
		if len(lectura.obtenerParametros()) == 1 && existeArchivoLog(lectura.obtenerParametros()[0]) {
			esValido = true
		}
	case comando2:
		if len(lectura.obtenerParametros()) == 2 {
			esValido = true
		}
	case comando3:
		if len(lectura.obtenerParametros()) == 1 {
			esValido = true
		}
	default:
		esValido = false
	}

	if !esValido {
		informeError = fmt.Sprintf("Error en comando %s\n", lectura.obtenerComando())
	}

	return esValido, informeError
}

func existeArchivoLog(archivo string) bool {

	if _, err := os.Stat(archivo); os.IsNotExist(err) {
		return false
	}
	return true
}

func (lectura ComandosLog) EjecutarComando(infoServidor Analisis.AnalisisLog) {

	switch lectura.obtenerComando() {
	case comando1:
		informeAnalisis := infoServidor.CargarArchivo(lectura.obtenerParametros()[0])
		fmt.Print(informeAnalisis)
	case comando3:
		numeroMasVisitados, _ := strconv.Atoi(lectura.obtenerParametros()[0])
		infoServidor.VerMasVisitados(numeroMasVisitados)
	default:
		break
	}

	fmt.Print("OK\n")
}
