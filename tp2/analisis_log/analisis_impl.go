package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDALista "tdas/lista"
	TDAAbb "tdas/pruebaAbb"
	TDAHash "tdas/pruebaHash"
	"time"
)

var compararVisitantes func(Ip, Ip) int = CompararIps

type datosServidor struct {
	visitantes TDAAbb.DiccionarioOrdenado[Ip, bool]
	recursos   TDAHash.Diccionario[string, int]
}

func GenerarDatos() AnalisisLog {

	return &datosServidor{TDAAbb.CrearABB[Ip, bool](compararVisitantes), TDAHash.CrearHash[string, int]()}
}

func (servidor *datosServidor) CargarArchivo(archivoLog string) (informe string) {

	conexionesLogActual := TDAHash.CrearHash[Ip, TDALista.Lista[time.Time]]()
	ipSospechosas := TDAAbb.CrearABB[Ip, bool](compararVisitantes)

	archivo, _ := os.Open(archivoLog)

	lecturaArchivo := bufio.NewScanner(archivo)

	for lecturaArchivo.Scan() {
		ipCargada, esSospechosaDeDos := servidor.cargarNuevaConexion(lecturaArchivo.Text(), conexionesLogActual)

		if esSospechosaDeDos {
			ipSospechosas.Guardar(ipCargada, true)
		}
	}

	return obtenerInforme(ipSospechosas)

}

func obtenerInforme(ipSospechosas TDAAbb.Diccionario[Ip, bool]) string {

	informeIpSospechosas := make([]string, 0)

	ipSospechosas.Iterar(func(ip Ip, dato bool) bool {

		informeIpSospechosas = append(informeIpSospechosas, fmt.Sprint("DoS: ", ObtenerStringDeIp(ip), "\n"))
		return true
	})

	return strings.Join(informeIpSospechosas, "")

}

func (servidor *datosServidor) cargarNuevaConexion(infoConexion string, conexionesLogActual TDAHash.Diccionario[Ip, TDALista.Lista[time.Time]]) (ipConexion Ip, esSospechosaDeDos bool) {

	datosLog := strings.Fields(infoConexion)

	ipConectada := ObtenerIpDeString(datosLog[0])
	horarioConexion, _ := time.Parse("2006-01-02T15:04:05+00:00", datosLog[1])
	recursoSolicitado := datosLog[3]

	if conexionesLogActual.Pertenece(ipConectada) {
		if seDetectoDoS(conexionesLogActual.Obtener(ipConectada), horarioConexion) {
			esSospechosaDeDos = true
		}
	} else {
		ultimasConexiones := TDALista.CrearListaEnlazada[time.Time]()
		ultimasConexiones.InsertarPrimero(horarioConexion)
		conexionesLogActual.Guardar(ipConectada, ultimasConexiones)
		servidor.visitantes.Guardar(ipConectada, true)
	}

	cantidadDeVisitas := 0

	if servidor.recursos.Pertenece(recursoSolicitado) {
		cantidadDeVisitas = servidor.recursos.Obtener(recursoSolicitado)
	}

	servidor.recursos.Guardar(recursoSolicitado, cantidadDeVisitas+1)

	return ipConectada, esSospechosaDeDos

}

func seDetectoDoS(ultimasConexiones TDALista.Lista[time.Time], horarioConexion time.Time) (esSospechosaDeDoS bool) {

	if ultimasConexiones.Largo() == 5 {
		return true
	}

	if ultimasConexiones.Largo() == 4 {
		diferenciaDeTiempo := horarioConexion.Sub(ultimasConexiones.VerPrimero())

		if diferenciaDeTiempo.Seconds() >= 2 {
			ultimasConexiones.BorrarPrimero()
		} else {
			esSospechosaDeDoS = true
		}
	}

	ultimasConexiones.InsertarUltimo(horarioConexion)

	return esSospechosaDeDoS

}
