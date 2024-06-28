package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	nodoRaiz  parClaveValor[K, V]
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	funcCmp  func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, funcCmp: funcion_cmp}
}

func crearNodoAbb[K comparable, V any](nuevoElemento parClaveValor[K, V]) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, nodoRaiz: nuevoElemento}
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {
	ramaABuscar := buscarRama(clave, &arbol.raiz, arbol.funcCmp)
	return ramaABuscar != nil && *ramaABuscar != nil
}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {

	ramaEncontrada := buscarRama(clave, &arbol.raiz, arbol.funcCmp)

	if *ramaEncontrada == nil {
		*ramaEncontrada = crearNodoAbb(parClaveValor[K, V]{clave, dato})
		arbol.cantidad++
	} else {
		nodoAModificar := *ramaEncontrada
		nodoAModificar.nodoRaiz.dato = dato
	}

}

func (arbol *abb[K, V]) Obtener(clave K) V {

	ramaEncontrada := buscarRama(clave, &arbol.raiz, arbol.funcCmp)

	if ramaEncontrada == nil || *ramaEncontrada == nil {
		panic("La clave no pertenece al diccionario")
	}
	nodoEncontrado := *ramaEncontrada
	return nodoEncontrado.nodoRaiz.dato
}

func (arbol *abb[K, V]) Borrar(clave K) V {

	ramaABorrar := buscarRama(clave, &arbol.raiz, arbol.funcCmp)

	if ramaABorrar == nil || *ramaABorrar == nil {
		panic("La clave no pertenece al diccionario")
	}

	datoBorrado := (*ramaABorrar).nodoRaiz.dato

	if (*ramaABorrar).derecho != nil && (*ramaABorrar).izquierdo != nil {
		borrarCon2Hijos(ramaABorrar)
	} else {
		borrarConMenosDe2Hijos(ramaABorrar)
	}

	arbol.cantidad--

	return datoBorrado
}

func borrarCon2Hijos[K comparable, V any](ramaAModificar **nodoAbb[K, V]) {

	nodoAModificar := *ramaAModificar

	ramaAModificar = buscarRamaSucesorInmediato(&(*ramaAModificar).derecho)

	nodoAModificar.nodoRaiz.clave = (*ramaAModificar).nodoRaiz.clave
	nodoAModificar.nodoRaiz.dato = (*ramaAModificar).nodoRaiz.dato

	borrarConMenosDe2Hijos(ramaAModificar)

}

func borrarConMenosDe2Hijos[K comparable, V any](ramaAModificar **nodoAbb[K, V]) {
	(*ramaAModificar) = (*ramaAModificar).hallarHijoNoNulo()
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *abb[K, V]) Iterar(funcion func(clave K, dato V) bool) {
	if arbol.raiz != nil {
		arbol.raiz.iterarRangoRec(nil, nil, funcion, arbol.funcCmp)
	}
}

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if arbol.raiz != nil {
		arbol.raiz.iterarRangoRec(desde, hasta, visitar, arbol.funcCmp)
	}
}

func (nodoArbol *nodoAbb[K, V]) iterarRangoRec(desde *K, hasta *K, visitar func(clave K, dato V) bool, comparar func(K, K) int) bool {
	if nodoArbol == nil {
		return true
	}
	if desde == nil || comparar(*desde, nodoArbol.nodoRaiz.clave) <= 0 {
		if !nodoArbol.izquierdo.iterarRangoRec(desde, hasta, visitar, comparar) {
			return false
		}
	}
	if (desde == nil || comparar(*desde, nodoArbol.nodoRaiz.clave) <= 0) &&
		(hasta == nil || comparar(*hasta, nodoArbol.nodoRaiz.clave) >= 0) {
		if !visitar(nodoArbol.nodoRaiz.clave, nodoArbol.nodoRaiz.dato) {
			return false
		}
	}
	if hasta == nil || comparar(*hasta, nodoArbol.nodoRaiz.clave) >= 0 {
		if !nodoArbol.derecho.iterarRangoRec(desde, hasta, visitar, comparar) {
			return false
		}
	}
	return true
}

func buscarRamaSucesorInmediato[K comparable, V any](ramaOrigen **nodoAbb[K, V]) **nodoAbb[K, V] {

	raiz := *ramaOrigen

	if raiz.izquierdo == nil {
		return ramaOrigen
	}

	sucesorInmediato := buscarRamaSucesorInmediato(&raiz.izquierdo)

	return sucesorInmediato
}

func (padre *nodoAbb[K, V]) hallarHijoNoNulo() *nodoAbb[K, V] {
	if padre.izquierdo == nil {
		return padre.derecho
	}
	return padre.izquierdo
}

func buscarRama[K comparable, V any](clave K, ramaOrigen **nodoAbb[K, V], comparar func(K, K) int) **nodoAbb[K, V] {

	actual := *ramaOrigen

	if actual == nil || actual.nodoRaiz.clave == clave {
		return ramaOrigen
	}

	var ramaEncontrada **nodoAbb[K, V]

	if comparar(clave, actual.nodoRaiz.clave) < 0 {
		ramaEncontrada = buscarRama(clave, &actual.izquierdo, comparar)
	} else {
		ramaEncontrada = buscarRama(clave, &actual.derecho, comparar)
	}

	return ramaEncontrada
}

type iteradorAbb[K comparable, V any] struct {
	nodosEnOrden TDAPila.Pila[*nodoAbb[K, V]]
	desde        *K
	hasta        *K
	comparar     func(K, K) int
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return arbol.IteradorRango(nil, nil)
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return crearIteradorAbb(arbol.raiz, desde, hasta, arbol.funcCmp)
}

func crearIteradorAbb[K comparable, V any](raiz *nodoAbb[K, V], desde *K, hasta *K, comparar func(K, K) int) *iteradorAbb[K, V] {

	nodosAIterar := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iterAbb := iteradorAbb[K, V]{nodosAIterar, desde, hasta, comparar}

	if iterAbb.desde == nil || iterAbb.hasta == nil || iterAbb.comparar(*iterAbb.desde, *iterAbb.hasta) <= 0 {
		iterAbb.apilarNodosMenores(raiz)
	}

	return &iterAbb
}

func (iterAbb *iteradorAbb[K, V]) HaySiguiente() bool {
	return !iterAbb.nodosEnOrden.EstaVacia()
}

func (iterAbb *iteradorAbb[K, V]) Siguiente() {

	if !iterAbb.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodoIterado := iterAbb.nodosEnOrden.Desapilar()
	iterAbb.apilarNodosMenores(nodoIterado.derecho)

}

func (iterAbb *iteradorAbb[K, V]) VerActual() (K, V) {

	if !iterAbb.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	actual := iterAbb.nodosEnOrden.VerTope()
	return actual.nodoRaiz.clave, actual.nodoRaiz.dato
}

func (iterAbb *iteradorAbb[K, V]) apilarNodosMenores(raiz *nodoAbb[K, V]) {

	if raiz == nil {
		return
	}

	raizEsMenorRango := false
	raizEsMayorRango := false

	if iterAbb.desde != nil && iterAbb.comparar(*iterAbb.desde, raiz.nodoRaiz.clave) > 0 {
		raizEsMenorRango = true
	} else if iterAbb.hasta != nil && iterAbb.comparar(*iterAbb.hasta, raiz.nodoRaiz.clave) < 0 {
		raizEsMayorRango = true
	}

	if !raizEsMenorRango && !raizEsMayorRango {
		iterAbb.nodosEnOrden.Apilar(raiz)
	}

	if raizEsMenorRango {
		iterAbb.apilarNodosMenores(raiz.derecho)
	} else {
		iterAbb.apilarNodosMenores(raiz.izquierdo)
	}

}
