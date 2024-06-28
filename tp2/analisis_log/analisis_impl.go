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
	_MAXIMASCONEXIONESDOS     = 5
	_TIEMPOENTRECONEXIONESDOS = 2
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

	fmt.Printf("Visitantes:\n")

	servidor.visitantes.IterarRango(&desde, &hasta, func(actual Ip, dato bool) bool {
		fmt.Printf("\t%s\n", ObtenerStringDeIp(actual))
		return true
	})
}

func (servidor *datosServidor) VerMasVisitados(n int) {
	masVisitados := TDAHeap.CrearHeap(func(recurso1, recurso2 parRecursoVisitas) int {
		return recurso1.visitas - recurso2.visitas
	})

	servidor.recursos.Iterar(func(recurso string, visitas int) bool {
		if masVisitados.Cantidad() < n {
			masVisitados.Encolar(parRecursoVisitas{recurso, visitas})
		} else if visitas > masVisitados.VerMax().visitas {
			masVisitados.Desencolar()
			masVisitados.Encolar(parRecursoVisitas{recurso, visitas})
		}
		return true
	})

	resultados := make([]parRecursoVisitas, masVisitados.Cantidad())
	for i := masVisitados.Cantidad() - 1; i >= 0; i-- {
		resultados[i] = masVisitados.Desencolar()
	}

	fmt.Printf("Sitios más visitados:\n")
	for _, recurso := range resultados {
		fmt.Printf("\t%s - %d\n", recurso.recurso, recurso.visitas)
	}
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

	if ultimasConexiones.Largo() == _MAXIMASCONEXIONESDOS {
		return true
	}

	if ultimasConexiones.Largo() == _MAXIMASCONEXIONESDOS-1 {
		diferenciaDeTiempo := horarioConexion.Sub(ultimasConexiones.VerPrimero())

		if diferenciaDeTiempo.Seconds() >= _TIEMPOENTRECONEXIONESDOS {
			ultimasConexiones.BorrarPrimero()
		} else {
			esSospechosaDeDoS = true
		}
	}

	ultimasConexiones.InsertarUltimo(horarioConexion)

	return esSospechosaDeDoS

}
