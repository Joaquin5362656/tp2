package tp2

type AnalisisLog interface {
	CargarArchivo(archivoLog string) (informe string)

	//	Visitantes(desde Ip, hasta Ip)

	VerMasVisitados(n int)
}
