package lista_test

import (
	"fmt"
	"strings"
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {

	listaInt := TDALista.CrearListaEnlazada[int]()

	require.True(t, listaInt.EstaVacia(), "Una lista recien creada esta vacia")
}

func TestAgregarBorrarElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.BorrarPrimero(), "Se espera que el primer elemento insertado sea el ultimo en ser borrado")
	require.Equal(t, 1, lista.BorrarPrimero(), "Se espera que el segundo elemento insertado sea el primero en ser borrado")
	require.True(t, lista.EstaVacia(), "La lista debera estar vacia despuees de desapilar todos los elementos")
}

func TestPruebaVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	cantidadElementos := 10000
	for i := 1; i <= cantidadElementos; i++ {
		lista.InsertarPrimero(i)
	}

	for i := cantidadElementos; i > 0; i-- {
		require.Equal(t, i, lista.VerPrimero(), "El tope de la lista debería ser el elemento correcto en cada iteracion")
		require.Equal(t, i, lista.BorrarPrimero(), "Se espera que el elemento borrado sea el correcto")
	}
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrar todos los elementos")
}

func TestAccionesInvalidasEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "La lista recin creada deberia estar vacia")

	require.Panics(t, func() { lista.BorrarPrimero() }, "Debe producirse un panic al intentar borrar un elemento de una lista vacia")
	require.Panics(t, func() { lista.VerPrimero() }, "Debe producirse un panic al intentar ver el primer elemento de una lista vacia")
}

func TestEstaVaciaEnListaRecienCreada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "La lista recien creada deberia estar vacia")

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia(), "La lista a la que se insertaron elementos no deberia estar vacia")
}

func TestAgregarDiferentesTiposDeDatos(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaInt.InsertarPrimero(1)
	listaInt.InsertarPrimero(2)
	require.Equal(t, 2, listaInt.BorrarPrimero())
	require.Equal(t, 1, listaInt.BorrarPrimero())
	require.True(t, listaInt.EstaVacia())

	listaString := TDALista.CrearListaEnlazada[string]()
	listaString.InsertarPrimero("Hola")
	listaString.InsertarPrimero("MUNDO")
	require.Equal(t, "MUNDO", listaString.BorrarPrimero())
	require.False(t, listaString.EstaVacia())
	require.Equal(t, "Hola", listaString.BorrarPrimero())
	require.True(t, listaString.EstaVacia())

	listaFloat := TDALista.CrearListaEnlazada[float64]()
	listaFloat.InsertarPrimero(3.14)
	listaFloat.InsertarPrimero(2.71)
	require.NotEqual(t, 3.14, listaFloat.BorrarPrimero())
	require.Equal(t, 3.14, listaFloat.BorrarPrimero())
}

func TestIteradorInternoTodosLosElementos(t *testing.T) {

	var (
		unNumero      = []*int{proElemento(5)}
		variosNumeros = []*int{proElemento(8), proElemento(2), proElemento(3), proElemento(4), proElemento(12)}
		mismosNumeros = []*int{proElemento(3), proElemento(3), proElemento(3), proElemento(3), proElemento(3), proElemento(3)}
		variosString  = []*string{proElemento("andres"), proElemento("lorenzo"), proElemento("carla"), proElemento("manuel")}
	)
	var (
		listaUnNumero      = TDALista.CrearListaEnlazada[*int]()
		listaVariosNumeros = TDALista.CrearListaEnlazada[*int]()
		listaMismosNumeros = TDALista.CrearListaEnlazada[*int]()
		listaVariosString  = TDALista.CrearListaEnlazada[*string]()
	)
	var (
		aumentarUno = func(numero *int) bool {
			*numero++
			return true
		}
		eliminarVocales = func(nombre *string) bool {
			*nombre = strings.ReplaceAll(*nombre, "a", "")
			*nombre = strings.ReplaceAll(*nombre, "e", "")
			*nombre = strings.ReplaceAll(*nombre, "i", "")
			*nombre = strings.ReplaceAll(*nombre, "o", "")
			*nombre = strings.ReplaceAll(*nombre, "u", "")
			return true
		}
	)

	insertarArrayAListaConPunteros(unNumero, listaUnNumero)
	listaUnNumero.Iterar(aumentarUno)
	visitarArray(unNumero, aumentarUno)
	require.Equalf(t, unNumero, borrarPrimerosNElementos(listaUnNumero, listaUnNumero.Largo()), "Se puede iterar un unico elemento de la lista aplicando la funcion indicada")

	insertarArrayAListaConPunteros(variosNumeros, listaVariosNumeros)
	listaVariosNumeros.Iterar(aumentarUno)
	visitarArray(variosNumeros, aumentarUno)
	require.Equalf(t, variosNumeros, borrarPrimerosNElementos(listaVariosNumeros, listaVariosNumeros.Largo()), "Iterar con iterador interno sobre todos los elementos aplica la funcion correctamente a cada uno de ellos")

	insertarArrayAListaConPunteros(mismosNumeros, listaMismosNumeros)
	listaMismosNumeros.Iterar(aumentarUno)
	visitarArray(mismosNumeros, aumentarUno)
	require.Equalf(t, mismosNumeros, borrarPrimerosNElementos(listaMismosNumeros, listaMismosNumeros.Largo()), "Se puede iterar una lista con mismos elementos, comportandose de la misma forma para cada elemento")

	insertarArrayAListaConPunteros(variosString, listaVariosString)
	listaVariosString.Iterar(eliminarVocales)
	visitarArray(variosString, eliminarVocales)
	require.Equalf(t, variosString, borrarPrimerosNElementos(listaVariosString, listaVariosString.Largo()), "Se puede iterar una lista de strings, aplicando la funcion a cada elemento en el orden y forma correcta")
}

func TestIteradorInternoHastaDevolverFalse(t *testing.T) {

	var (
		unNumeroCumpleCond            = []*int{proElemento(5)}
		unNumeroNoCumple              = []*int{proElemento(2)}
		numerosImparesMenosUno        = []*int{proElemento(3), proElemento(7), proElemento(13), proElemento(4), proElemento(11), proElemento(1)}
		positivosInicioNegativosFinal = []*int{proElemento(3), proElemento(6), proElemento(7), proElemento(-3), proElemento(-8), proElemento(-31)}
		primerNumeroNoCumple          = []*int{proElemento(6), proElemento(5), proElemento(2), proElemento(7), proElemento(8)}
	)
	var (
		listaUnNumeroCumpleCond            = TDALista.CrearListaEnlazada[*int]()
		listaUnNumeroNoCumple              = TDALista.CrearListaEnlazada[*int]()
		listaNumerosImparesMenosUno        = TDALista.CrearListaEnlazada[*int]()
		listaPositivosInicioNegativosFinal = TDALista.CrearListaEnlazada[*int]()
		listaPrimerNumeroNoCumple          = TDALista.CrearListaEnlazada[*int]()
	)
	var (
		multiplicarx2TodosLosElementos = func(numero *int) bool {
			*numero = *numero * 2
			return true
		}
		multiplicarx2HastaEncontrarPar = func(numero *int) bool {
			if *numero%2 == 1 {
				*numero = *numero * 2
				return true
			} else {
				return false
			}
		}
		invertirSignoATodos = func(numero *int) bool {
			*numero = *numero * -1
			return true
		}
		pasarAElementosNegativosHastaEncontraUno = func(numero *int) bool {

			if *numero <= 0 {
				return false
			}

			*numero = *numero * -1
			return true
		}
	)

	insertarArrayAListaConPunteros(unNumeroCumpleCond, listaUnNumeroCumpleCond)
	listaUnNumeroCumpleCond.Iterar(multiplicarx2HastaEncontrarPar)
	visitarArray(unNumeroCumpleCond, multiplicarx2TodosLosElementos)
	require.Equalf(t, unNumeroCumpleCond, borrarPrimerosNElementos(listaUnNumeroCumpleCond, listaUnNumeroCumpleCond.Largo()), "Se itera una lista con un solo elemento que cumple la condicion aplicandole la funcion correspondiente")

	insertarArrayAListaConPunteros(unNumeroNoCumple, listaUnNumeroNoCumple)
	listaUnNumeroNoCumple.Iterar(multiplicarx2HastaEncontrarPar)
	visitarArray(unNumeroNoCumple, multiplicarx2HastaEncontrarPar)
	require.Equalf(t, unNumeroNoCumple, borrarPrimerosNElementos(listaUnNumeroNoCumple, listaUnNumeroNoCumple.Largo()), "Se itera una lista con un unico elemento que devuelve false y se cumple lo indicado en la funcion correctamente")

	insertarArrayAListaConPunteros(numerosImparesMenosUno, listaNumerosImparesMenosUno)
	listaNumerosImparesMenosUno.Iterar(multiplicarx2HastaEncontrarPar)
	visitarArray(numerosImparesMenosUno, multiplicarx2HastaEncontrarPar)
	require.Equalf(t, numerosImparesMenosUno, borrarPrimerosNElementos(listaNumerosImparesMenosUno, listaNumerosImparesMenosUno.Largo()), "Se iteran correctamente aplicando la funcion a cada elemento hasta encontrar uno que devuelva false, dejando de iterar elementos")

	insertarArrayAListaConPunteros(positivosInicioNegativosFinal, listaPositivosInicioNegativosFinal)
	listaPositivosInicioNegativosFinal.Iterar(pasarAElementosNegativosHastaEncontraUno)
	visitarArray(positivosInicioNegativosFinal, invertirSignoATodos)
	require.NotEqualf(t, positivosInicioNegativosFinal, borrarPrimerosNElementos(listaPositivosInicioNegativosFinal, listaPositivosInicioNegativosFinal.Largo()), "Al iterar una lista donde un elemento devuelve false, no se iteran todos los elementos y no se le aplica la funcion a todos estos")

	visitarArray(positivosInicioNegativosFinal, invertirSignoATodos)

	insertarArrayAListaConPunteros(positivosInicioNegativosFinal, listaPositivosInicioNegativosFinal)
	listaPositivosInicioNegativosFinal.Iterar(pasarAElementosNegativosHastaEncontraUno)
	visitarArray(positivosInicioNegativosFinal, pasarAElementosNegativosHastaEncontraUno)
	require.Equalf(t, positivosInicioNegativosFinal, borrarPrimerosNElementos(listaPositivosInicioNegativosFinal, listaPositivosInicioNegativosFinal.Largo()), "Al iterar una lista donde un elementos devuelve false, se iteran correctamente los primeros elementos hasta encontrar el que devuelve false")

	insertarArrayAListaConPunteros(primerNumeroNoCumple, listaPrimerNumeroNoCumple)
	listaPrimerNumeroNoCumple.Iterar(multiplicarx2HastaEncontrarPar)
	require.Equal(t, primerNumeroNoCumple, borrarPrimerosNElementos(listaPrimerNumeroNoCumple, listaPrimerNumeroNoCumple.Largo()), "En una lista que se itera, si el primer elemento devuelve false, deja de iterar no aplicando la funcion a ningun elemento y dejando la lista igual que al inicio")
}

func TestIteradorExternoOperacionesInvalidasYFuncionamientoBasico(t *testing.T) {

	var (
		array = []*int{proElemento(3), proElemento(9), proElemento(12), proElemento(4)}
	)
	var (
		listaVacia   = TDALista.CrearListaEnlazada[*int]()
		listaNoVacia = TDALista.CrearListaEnlazada[*int]()
	)

	iterador := listaVacia.Iterador()
	require.False(t, iterador.HaySiguiente(), "Un iterador creado de una lista vacia no tiene ningun elemento siguiente para iterar")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() }, "Hubo panic error al tratar de ver el actual de un iterador creado de una lista vacia")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() }, "Hubo panic error al tratar de avanzar al siguiente elemento en un iterador creado de una lista vacia")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() }, "Hubo panic error al tratar de borrar un elemento con un iterador creado de una lista vacia")

	insertarArrayAListaConPunteros(array, listaNoVacia)
	iterador = listaNoVacia.Iterador()
	require.True(t, iterador.HaySiguiente(), "Un iterador creado de una lista no vacia tiene un elemento siguiente para iterar")
	require.Equal(t, listaNoVacia.VerPrimero(), iterador.VerActual(), "Ver el elemento actual de Un iterador que no avanzo a siguiente posicion es el primer elemento de la lista")

	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}

	require.False(t, iterador.HaySiguiente(), "Un iterador puede avanzar al siguiente elemento hasta que no haya siguiente elemento")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() }, "Hubo panic error al tratar de ver el actual de un iterador creado de una lista no vacia en la que ya se avanzo por todos los elementos")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() }, "Hubo panic error al tratar de avanzar al siguiente elemento en un iterador creado de una lista no vacia en la que ya se avanzo por todos los elementos")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() }, "Hubo panic error al tratar de borrar un elemento con un iterador creado de una lista no vacia en la que ya se avanzo con todos los elementos")
}

func TestIteradorExternoVerActual(t *testing.T) {

	var (
		array = []int{6, 11, 4, 2}
	)
	var (
		pruebaLista            = TDALista.CrearListaEnlazada[int]()
		pruebaIteracionEnOrden = TDALista.CrearListaEnlazada[int]()
	)

	insertarArrayALista(array, pruebaLista)
	iterador := pruebaLista.Iterador()

	require.Equal(t, pruebaLista.VerPrimero(), iterador.VerActual(), "El actual de un iterador recien creado es el primer elemento de la lista")
	iterador.Siguiente()
	require.NotEqual(t, pruebaLista.VerPrimero(), iterador.VerActual(), "Al ir al siguiente con un iterador, el actual deja de apuntar al primer elemento")
	require.Equal(t, array[1], iterador.VerActual(), "Al ir al siguiente con un iterador, se avanza al segundo elemento de la lista")

	ultimoLista := iterador.VerActual()
	for iterador.HaySiguiente() {
		ultimoLista = iterador.VerActual()
		iterador.Siguiente()
	}

	require.Equal(t, pruebaLista.VerUltimo(), ultimoLista, "El ultimo elemento que fue iterado es el ultimo elemento en la lista")

	insertarArrayALista(array, pruebaIteracionEnOrden)
	require.Equal(t, listaAArrayConIterador(pruebaIteracionEnOrden), array, "Se itera todos los elementos en mismo orden que insercion en la lista y se aplica funcion correctamente")
}

func TestIteradorExternoInsertarYBorrarUnSoloElemento(t *testing.T) {

	listaCreadaSinElementos := TDALista.CrearListaEnlazada[int]()

	iterador := listaCreadaSinElementos.Iterador()
	iterador.Insertar(6)
	require.Equal(t, iterador.VerActual(), 6, "Insertar un elemento con el iterador en una lista vacia cambia el actual al elemento insertado")
	require.Equal(t, listaCreadaSinElementos.VerPrimero(), iterador.VerActual(), "Insertar un elemento con el iterador de una lista vacia, modifica el primero de la lista al elemento recien insertado")
	require.Equal(t, listaCreadaSinElementos.Largo(), 1, "Insertar un elemento con el iterador de una lista vacia aumenta el largo de la lista en 1")
	require.Equal(t, listaCreadaSinElementos.VerUltimo(), iterador.VerActual(), "Insertar un elemento con el iterador de una lista vacia, modifica el ultimo de la lista al elemento recien insertado")

	elementoEliminado := iterador.Borrar()
	require.Equal(t, elementoEliminado, 6, "Borrar el unico elemento de una lista con el iterador devuelve el elemento indicado")
	require.False(t, iterador.HaySiguiente(), "Al borrar el unico elemento de la lista, el iterador deja de tener siguiente elemento terminando de iterar")
	require.True(t, listaCreadaSinElementos.EstaVacia(), "Borrar el unico elemento de una lista con el iterador deja una lista vacia")
	require.Equal(t, listaCreadaSinElementos.Largo(), 0, "Al borrar el unico elemento de una lista, se modifica el largo de esta quedando en 0")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() }, "Hubo panic error al tratar de ver el actual de un iterador con el que borramos el unico elemento de la lista")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() }, "Hubo panic error al tratar de avanzar al siguiente elemento en un iterador con el que orramos el unico elemento de la lista")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() }, "Hubo panic error al tratar de borrar un elemento con un iterador con el que borramos el unico elemento de la lista")

}

func TestInsertarConIteradorExterno(t *testing.T) {

	var (
		pruebaNumeros = []int{6, 12, 4, 9, 8, 5, 13}
		pruebaString  = []string{"Algoritmo", "datos", "sort", "goland", "suricata", "hakuna batata", "peekaboo"}
		pruebaRunas   = []rune{'?', 'R', '°', '^', 'K', '|', '¸'}
	)
	var (
		numerosAAgregar = [5]int{6, 5, 10, 9, 0}
		stringAAgregar  = [5]string{"pila", "cola", "hash", "grafos", "python"}
		runasAAgregar   = [5]rune{'=', '_', '¬', ';', '%'}
	)

	testInsertarConIteradorExternoGenerico[int](t, pruebaNumeros, numerosAAgregar)
	testInsertarConIteradorExternoGenerico[string](t, pruebaString, stringAAgregar)
	testInsertarConIteradorExternoGenerico[rune](t, pruebaRunas, runasAAgregar)
}

func testInsertarConIteradorExternoGenerico[T any](t *testing.T, datos []T, elementosAInsertar [5]T) {

	var tipoDeDato string = fmt.Sprintf("%T", datos[0])
	listaPruebaInsertar := TDALista.CrearListaEnlazada[T]()

	insertarArrayALista(datos, listaPruebaInsertar)
	iterador := listaPruebaInsertar.Iterador()

	iterador.Insertar(elementosAInsertar[0])
	require.Equal(t, iterador.VerActual(), elementosAInsertar[0], "Insertar en lista no vacia al inicio cambia el actual al elemento insertado con el tipo de dato %s", tipoDeDato)
	require.Equal(t, listaPruebaInsertar.VerPrimero(), iterador.VerActual(), "Insertar un elemento al principio de una lista no vacia modifica el primero de la lista pasando a ser este elemento para el tipo de dato %s", tipoDeDato)

	iterador.Siguiente()
	require.Equal(t, iterador.VerActual(), datos[0], "El siguiente elemento del iterador al que insertamos un elemento al principio es el que antes se encontraba al principio de la lista para el tipo de dato %s", tipoDeDato)

	posicionesAInsertar := []int{2, 5, 6}

	arrayPrevioAInsercion := listaAArrayConIterador(listaPruebaInsertar)
	insertarEnPosiciones(listaPruebaInsertar, posicionesAInsertar, elementosAInsertar[1:5])
	arrayEsperado := resultadoEsperado(arrayPrevioAInsercion, posicionesAInsertar, elementosAInsertar[1:5])
	arrayPosInsercion := listaAArrayConIterador(listaPruebaInsertar)

	require.Equal(t, arrayPosInsercion, arrayEsperado, "Insertar con iterador en distintas posiciones de la lista, inserta correctamente los elementos en la posicion esperada para el tipo de dato %s", tipoDeDato)

	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}

	iterador.Insertar(elementosAInsertar[4])
	require.Equal(t, iterador.VerActual(), elementosAInsertar[4], "Insertar al final de una lista con iterador modifica el actual al elemento insertado para el tipo de dato %s", tipoDeDato)
	require.Equal(t, listaPruebaInsertar.VerUltimo(), iterador.VerActual(), "Al insertar al final de una lista con interador se modifica el ultimo elemento de la lista para el tipo de dato %s", tipoDeDato)

	require.True(t, iterador.HaySiguiente(), "Al insertar al final con iterador externo, el iterador detecta que hay un siguiente elemento a iterar para el tipo de dato %s", tipoDeDato)
	iterador.Siguiente()
	require.False(t, iterador.HaySiguiente(), "Al pasar al siguiente elemento, luego de insertar en ultima posicion con el iterador este detecta que no hay siguiente para el tipo de dato %s", tipoDeDato)

}

func TestBorrarConIteradorExterno(t *testing.T) {

	var (
		pruebaNumeros = [8]int{2, 19, 3, 2, 15, 6, 29, 19}
		pruebaString  = [8]string{"Messi", "Jesus", "sort", "goland", "modulo", "hakuna batata", "peekaboo", "lista"}
		pruebaRunas   = [8]rune{'?', 'R', '°', '^', 'K', '|', '¸', '	'}
	)

	testBorrarConIteradorExternoGenerico[int](t, pruebaNumeros)
	testBorrarConIteradorExternoGenerico[string](t, pruebaString)
	testBorrarConIteradorExternoGenerico[rune](t, pruebaRunas)

	testBorrarConIteradorExternoGenerico(t, [8]int{5, 8, 12, 7, 1, 3, 6, 14})
}

func testBorrarConIteradorExternoGenerico[T any](t *testing.T, datos [8]T) {

	var tipoDeDato string = fmt.Sprintf("%T", datos[0])
	listaPruebaBorrar := TDALista.CrearListaEnlazada[T]()

	insertarArrayALista(datos[:], listaPruebaBorrar)
	iterador := listaPruebaBorrar.Iterador()

	require.Equal(t, listaPruebaBorrar.VerPrimero(), iterador.Borrar(), "Borrar con el iterador externo recien creado elimina el primer elemento de la lista para el tipo de dato %s", tipoDeDato)
	require.Equal(t, iterador.VerActual(), datos[1], "Al borrar el primer elemento con el iterador externo, el actual pasa al siguiente de la lista para el tipo de dato %s", tipoDeDato)
	require.Equal(t, listaPruebaBorrar.VerPrimero(), iterador.VerActual(), "Al borrar el primer elemento con el iterador externo, el primero de la lista se actualiza al siguiente elemento para el tipo de dato %s", tipoDeDato)

	iterador.Siguiente()
	iterador.Siguiente()
	eliminado := iterador.Borrar()

	require.Equal(t, eliminado, datos[3], "Borrar un elemento en el medio de una lista devuelve el elemento correcto para el tipo de dato %s", tipoDeDato)
	require.Equal(t, iterador.VerActual(), datos[4], "Al borrar un elemento en el medio de una lista con iterador, el actual de este pasa a ser el sigueinte dato en la lista para el tipo %s", tipoDeDato)

	datosPrevioABorrar := listaAArrayConIterador(listaPruebaBorrar)

	posicionesABorrar := []int{1, 2, 4}
	elementosBorrados := borrarEnPosiciones(listaPruebaBorrar, posicionesABorrar)
	require.Equal(t, elementosBorrados, []T{datosPrevioABorrar[1], datosPrevioABorrar[2], datosPrevioABorrar[4]}, "Borrar elementos con el iterador externo devuelve los elementos eliminados en el mismo orden que se recorrio y elimino la lista para el tipo de dato %s", tipoDeDato)

	datosPosBorrar := listaAArrayConIterador(listaPruebaBorrar)
	require.Equal(t, datosPosBorrar, []T{datosPrevioABorrar[0], datosPrevioABorrar[3], datosPrevioABorrar[5]}, "Eliminar elementos en el medio de una lista con iterador elimina en la posicion correcta para el tipo de dato %s", tipoDeDato)

	iteradorHastaElFinal := listaPruebaBorrar.Iterador()
	iteradorHastaUltimo := listaPruebaBorrar.Iterador()
	anteUltimo := iteradorHastaUltimo.VerActual()

	for iteradorHastaElFinal.HaySiguiente() {
		iteradorHastaElFinal.Siguiente()
		if iteradorHastaElFinal.HaySiguiente() {
			anteUltimo = iteradorHastaUltimo.VerActual()
			iteradorHastaUltimo.Siguiente()
		}
	}

	iteradorHastaUltimo.Borrar()
	require.False(t, iteradorHastaUltimo.HaySiguiente(), "Eliminar el ultimo elemento de una lista con iterador externo, hace que ya no haya siguiente elemento a iterar para tipo de dato %s", tipoDeDato)
	require.Equal(t, listaPruebaBorrar.VerUltimo(), anteUltimo, "Eliminar el ultimo elemento de una lista con iterador externo modifica el ultimo de la lista pasando a ser el anterior al eliminado para el tipo de dato %s", tipoDeDato)

}

func proElemento[T any](dato T) *T {

	nuevoProT := new(T)
	*nuevoProT = dato
	return nuevoProT

}

// Le aplica la funcion a todos los elementos de un array o hasta que la
// funcion devuelve false, dejando de aplicar la funcion al resto de los
// elementos
func visitarArray[T any](datos []T, visitar func(T) bool) {

	var seguirRecorriendo bool = true
	var i int = 0

	for seguirRecorriendo && len(datos) > i {
		seguirRecorriendo = visitar(datos[i])
		i++
	}

}

// Devuelve un array donde, en las posiciones indicadas se insertan los elementos
// pasados por parametros y en el resto de las posiciones se agregan los datos
// iniciales pasados por parametro respetando el orden dado por esto.
// Las posiciones a insertar deben estar ordenadas de menor a mayor.
func resultadoEsperado[T any](datosIniciales []T, posicionesAInsertarEnOrden []int, elementosAInsertar []T) []T {

	resultado := make([]T, 0, len(datosIniciales)+len(posicionesAInsertarEnOrden))
	posicion := 0

	for len(posicionesAInsertarEnOrden) > 0 {
		if posicion == posicionesAInsertarEnOrden[0] {
			resultado = append(resultado, elementosAInsertar[0])
			posicionesAInsertarEnOrden = posicionesAInsertarEnOrden[1:]
			elementosAInsertar = elementosAInsertar[1:]
		} else {
			resultado = append(resultado, datosIniciales[0])
			datosIniciales = datosIniciales[1:]
		}
		posicion++
	}

	resultado = append(resultado, datosIniciales...)
	return resultado
}

// Inserta los elementos pasados por parametro en una lista usando el iterador externo
// e insertandolos en las posiciones pasadas por parametro.
// Las posiciones pasadas por parametro deben estar en orden de menor a mayor.
func insertarEnPosiciones[T any](lista TDALista.Lista[T], posicionesAInsertarEnOrden []int, elementosAInsertar []T) {

	iterador := lista.Iterador()
	posicion := 0

	for iterador.HaySiguiente() && len(posicionesAInsertarEnOrden) > 0 {
		if posicion == posicionesAInsertarEnOrden[0] {
			iterador.Insertar(elementosAInsertar[0])
			posicionesAInsertarEnOrden = posicionesAInsertarEnOrden[1:]
			elementosAInsertar = elementosAInsertar[1:]
		}
		posicion++
		iterador.Siguiente()
	}
}

// Borra los elementos de una lista usando el iterador externo en las posiciones
// pasadas por parametro y devuelve los elementos borrados en un array en el mismo
// orden en que se borraron los elementos
func borrarEnPosiciones[T any](lista TDALista.Lista[T], posicionesABorrarEnOrden []int) []T {

	iterador := lista.Iterador()
	elementosBorrados := make([]T, 0, len(posicionesABorrarEnOrden))
	posicion := 0

	for iterador.HaySiguiente() && len(posicionesABorrarEnOrden) > 0 {
		if posicion == posicionesABorrarEnOrden[0] {
			elementosBorrados = append(elementosBorrados, iterador.Borrar())
			posicionesABorrarEnOrden = posicionesABorrarEnOrden[1:]
		} else {
			iterador.Siguiente()
		}
		posicion++
	}

	return elementosBorrados
}

// Inserta en la lista pasada por parametro los elementos que se encuentran en el array
// en el mismo ordena que el dispuesto por este.
func insertarArrayALista[T any](datos []T, lista TDALista.Lista[T]) {

	for _, valor := range datos {
		lista.InsertarUltimo(valor)
	}
}

// insertarArrayAListaConPunteros inserta en la lista de punteros pasada por parametro
// los valores en la posicion de memoria almacenada en el array en el mismo orden
// dispuesto por este.
func insertarArrayAListaConPunteros[T any](datos []*T, lista TDALista.Lista[*T]) {

	for _, valor := range datos {
		lista.InsertarUltimo(proElemento(*valor))
	}

}

// Devuelve un array con el mismo orden que la lista, es decir que el primer elemento de la
// lista es el elemento en la posicion 0 del array, el siguiente al primer elemento de la
// lista es el elemento en la posicion 1 del array y asi sucesivamente hasta que el ultimo
// elemento de la lista sea el elemento en la ultima posicion del array.
func listaAArrayConIterador[T any](lista TDALista.Lista[T]) []T {

	datos := make([]T, 0, lista.Largo())

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		datos = append(datos, iterador.VerActual())
		iterador.Siguiente()
	}

	return datos
}

// borrarPrimerosNElementos saca los primeros N elementos de la lista indicados por parametro,
// devolviendo un array con el mismo orden en que se encontraban en la lista
func borrarPrimerosNElementos[T any](lista TDALista.Lista[T], numeroABorrar int) []T {

	arrayBorrados := make([]T, 0, numeroABorrar)

	for numeroABorrar > 0 {
		arrayBorrados = append(arrayBorrados, lista.BorrarPrimero())
		numeroABorrar--
	}

	return arrayBorrados
}
