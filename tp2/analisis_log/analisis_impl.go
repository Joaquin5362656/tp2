package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
	"time"
)

const (
	_MaximasConexionesDoS     = 5
	_TiempoEntreConexionesDoS = 2
)

var compararVisitantes func(Ip, Ip) int = CompararIps

type datosServidor struct {
	visitantes TDADiccionario.DiccionarioOrdenado[Ip, bool]
	recursos   TDADiccionario.Diccionario[string, int]
}

type parRecursoVisitas struct {
	recurso string
	visitas int
}

func GenerarRegistros() AnalisisLog {

	return &datosServidor{
		visitantes: TDADiccionario.CrearABB[Ip, bool](compararVisitantes),
		recursos:   TDADiccionario.CrearHash[string, int](),
	}
}

func (servidor *datosServidor) CargarArchivo(archivoLog string) (informe string) {

	conexionesLogActual := TDADiccionario.CrearHash[Ip, TDALista.Lista[time.Time]]()
	ipSospechosas := TDADiccionario.CrearABB[Ip, bool](compararVisitantes)

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

func (servidor *datosServidor) Visitantes(desde Ip, hasta Ip) {

	print("Visitantes:\n")

	servidor.visitantes.IterarRango(&desde, &hasta, func(actual Ip, dato bool) bool {
		fmt.Printf("\t%s\n", ObtenerStringDeIp(actual))
		return true
	})
}

func (servidor *datosServidor) VerMasVisitados(n int) {

	masVisitados := TDAHeap.CrearHeap(func(recurso1 parRecursoVisitas, recurso2 parRecursoVisitas) int {
		return recurso2.visitas - recurso1.visitas
	})

	añadirNMasVisitasActual := func(recurso string, visitas int) bool {

		if masVisitados.Cantidad() >= n {
			if visitas > masVisitados.VerMax().visitas {
				masVisitados.Desencolar()
			} else {
				return true
			}
		}

		masVisitados.Encolar(parRecursoVisitas{recurso, visitas})
		return true
	}

	servidor.recursos.Iterar(añadirNMasVisitasActual)

	print("Sitios más visitados:\n")

	mostrarMasVisitados(masVisitados)
}

func mostrarMasVisitados(masVisitados TDAHeap.ColaPrioridad[parRecursoVisitas]) {

	if masVisitados.Cantidad() == 0 {
		return
	}

	menosVisitado := masVisitados.Desencolar()

	mostrarMasVisitados(masVisitados)

	fmt.Printf("\t%s - %d\n", menosVisitado.recurso, menosVisitado.visitas)
}

func obtenerInforme(ipSospechosas TDADiccionario.Diccionario[Ip, bool]) string {

	informeIpSospechosas := make([]string, 0)

	ipSospechosas.Iterar(func(ip Ip, dato bool) bool {

		informeIpSospechosas = append(informeIpSospechosas, fmt.Sprint("DoS: ", ObtenerStringDeIp(ip), "\n"))
		return true
	})

	return strings.Join(informeIpSospechosas, "")

}

func (servidor *datosServidor) cargarNuevaConexion(infoConexion string, conexionesLogActual TDADiccionario.Diccionario[Ip, TDALista.Lista[time.Time]]) (ipConexion Ip, esSospechosaDeDos bool) {

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

	if ultimasConexiones.Largo() == _MaximasConexionesDoS {
		return true
	}

	if ultimasConexiones.Largo() == _MaximasConexionesDoS-1 {
		diferenciaDeTiempo := horarioConexion.Sub(ultimasConexiones.VerPrimero())

		if diferenciaDeTiempo.Seconds() >= _TiempoEntreConexionesDoS {
			ultimasConexiones.BorrarPrimero()
		} else {
			esSospechosaDeDoS = true
		}
	}

	ultimasConexiones.InsertarUltimo(horarioConexion)

	return esSospechosaDeDoS

}
