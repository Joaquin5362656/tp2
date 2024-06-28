package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func crearNodoLista[T any](dato T, siguiente *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato, siguiente}
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

// FUncion que crea una lista enlazada
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := crearNodoLista(elemento, nil)

	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := crearNodoLista(elemento, nil)

	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--

	if lista.EstaVacia() {
		lista.ultimo = nil
	}

	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}

	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) (seguirIterando bool)) {

	lista.primero.iterarElemento(visitar)
}

func (actual *nodoLista[T]) iterarElemento(visitar func(T) bool) {

	if actual == nil {
		return
	}

	seguirIterando := visitar(actual.dato)

	if seguirIterando {
		actual.siguiente.iterarElemento(visitar)
	} else {
		return
	}
}

func crearIterador[T any](lista *listaEnlazada[T], actual *nodoLista[T], anterior *nodoLista[T]) *iteradorLista[T] {
	return &iteradorLista[T]{lista, actual, anterior}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return crearIterador[T](lista, lista.primero, nil)
}

func (iterador *iteradorLista[T]) VerActual() T {

	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iterador.actual.dato
}

func (iterador *iteradorLista[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iteradorLista[T]) Siguiente() {

	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	anteriorLista := iterador.actual
	iterador.actual = iterador.actual.siguiente
	iterador.anterior = anteriorLista
}

func (iterador *iteradorLista[T]) Insertar(dato T) {

	anteriorLista := iterador.anterior
	esFinalLista := !iterador.HaySiguiente()
	esInicioLista := anteriorLista == nil

	iterador.actual = crearNodoLista(dato, iterador.actual)

	if esInicioLista {
		iterador.lista.primero = iterador.actual
	} else {
		anteriorLista.siguiente = iterador.actual
	}

	if esFinalLista {
		iterador.lista.ultimo = iterador.actual
	}

	iterador.lista.largo++
}

func (iterador *iteradorLista[T]) Borrar() T {

	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	elementoABorrar := iterador.actual.dato

	esFinalLista := iterador.actual.siguiente == nil
	esInicioLista := iterador.anterior == nil

	iterador.actual = iterador.actual.siguiente

	if esInicioLista {
		iterador.lista.primero = iterador.actual
	} else {
		anteriorLista := iterador.anterior
		anteriorLista.siguiente = iterador.actual
	}

	if esFinalLista {
		iterador.lista.ultimo = iterador.anterior
	}

	iterador.lista.largo--

	return elementoABorrar

}
