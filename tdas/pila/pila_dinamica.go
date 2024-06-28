package pila

/* Definición del struct pila proporcionado por la cátedra. */

const tamanioInicial = 10

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, tamanioInicial), cantidad: 0}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {

	if pila.EstaVacia() {
		panic("La pila esta vacia")
	} else {
		return pila.datos[pila.cantidad-1]
	}

}

func (pila *pilaDinamica[T]) redimensionar(nuevoTamanio int) {

	datosRedimensionados := make([]T, nuevoTamanio)
	copy(datosRedimensionados, pila.datos)
	pila.datos = datosRedimensionados
}

func (pila *pilaDinamica[T]) Apilar(dato T) {

	if pila.cantidad == len(pila.datos) {
		pila.redimensionar(len(pila.datos) * 2)
	}

	pila.datos[pila.cantidad] = dato
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {

	if !pila.EstaVacia() {

		pila.cantidad--
		datoDesapilado := pila.datos[pila.cantidad]

		if pila.cantidad < len(pila.datos)/4 && (len(pila.datos)/2) > tamanioInicial {
			pila.redimensionar(len(pila.datos) / 2)
		}

		return datoDesapilado

	} else {
		panic("La pila esta vacia")
	}
}
