package cola_prioridad

const (
	factorAchicar     = 4
	factorRedimension = 2
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: []T{},
		cant:  0,
		cmp:   funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	copia := make([]T, len(arreglo))
	copy(copia, arreglo)

	heap := &colaConPrioridad[T]{
		datos: copia,
		cant:  len(copia),
		cmp:   funcion_cmp,
	}
	heap.heapify()
	return heap
}

func (cola *colaConPrioridad[T]) EstaVacia() bool {
	return cola.Cantidad() == 0
}

func (cola *colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}

func (cola *colaConPrioridad[T]) Encolar(elem T) {
	if cola.cant >= len(cola.datos) {
		if cola.cant == 0 {
			cola.redimensionar(factorRedimension)
		} else {
			cola.redimensionar(factorRedimension * cola.cant)
		}
	}
	cola.datos[cola.cant] = elem
	cola.cant++
	upHeap(cola.datos, cola.cant-1, cola.cmp)
}

func (cola *colaConPrioridad[T]) VerMax() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.datos[0]
}

func (cola *colaConPrioridad[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	maxElem := cola.datos[0]
	cola.datos[0] = cola.datos[cola.cant-1]
	cola.cant--
	downHeap(cola.datos, cola.cant, 0, cola.cmp)
	cola.datos = cola.datos[:cola.cant]

	if factorAchicar*cola.cant <= len(cola.datos) && cola.cant > 0 {
		cola.redimensionar(cola.cant / factorRedimension)
	}

	return maxElem
}

func (cola *colaConPrioridad[T]) heapify() {
	heapify(cola.datos, cola.cant, cola.cmp)
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, len(elementos), funcion_cmp)

	for i := len(elementos) - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downHeap(elementos, i, 0, funcion_cmp)
	}
}

func heapify[T any](datos []T, n int, cmp func(T, T) int) {
	for i := n/2 - 1; i >= 0; i-- {
		downHeap(datos, n, i, cmp)
	}
}

func downHeap[T any](datos []T, n, pos int, cmp func(T, T) int) {
	for {
		hijoIzq := 2*pos + 1
		hijoDer := 2*pos + 2
		mayor := pos

		if hijoIzq < n && cmp(datos[hijoIzq], datos[mayor]) > 0 {
			mayor = hijoIzq
		}
		if hijoDer < n && cmp(datos[hijoDer], datos[mayor]) > 0 {
			mayor = hijoDer
		}
		if mayor == pos {
			break
		}
		datos[pos], datos[mayor] = datos[mayor], datos[pos]
		pos = mayor
	}
}

func upHeap[T any](datos []T, pos int, cmp func(T, T) int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if cmp(datos[pos], datos[padre]) <= 0 {
			break
		}
		datos[pos], datos[padre] = datos[padre], datos[pos]
		pos = padre
	}
}

func (cola *colaConPrioridad[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, cola.datos)
	cola.datos = nuevosDatos
}
