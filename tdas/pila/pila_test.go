package pila_test

import (
	"strconv"
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func apilarDatos[T any](datos []T, pila TDAPila.Pila[T]) {

	for _, dato := range datos {
		pila.Apilar(dato)
	}
}

func desapilarDatos[T any](cantidadADesapilar int, pila TDAPila.Pila[T]) {

	for cantidadADesapilar > 0 {
		pila.Desapilar()
		cantidadADesapilar--
	}
}

func pilaDesapiladaEnOrden[T int | string | *TDAPila.Pila[int] | *any](datosEnOrdenEsperado []T, pila TDAPila.Pila[T]) bool {

	if len(datosEnOrdenEsperado) == 0 {
		return true
	}

	if pila.Desapilar() != datosEnOrdenEsperado[0] {
		return false
	}

	return pilaDesapiladaEnOrden(datosEnOrdenEsperado[1:], pila)
}

func TestPilaVacia(t *testing.T) {

	pilaInt := TDAPila.CrearPilaDinamica[int]()
	t.Run("Pila vacia con int", func(t *testing.T) {
		require.True(t, pilaInt.EstaVacia(), "Una pila recien creada debe estar vacia")
	})

	pilaString := TDAPila.CrearPilaDinamica[string]()

	t.Run("Pila vacia con string", func(t *testing.T) {
		require.True(t, pilaString.EstaVacia(), "Una pila recien creada debe estar vacia")
	})

	pilaPilas := TDAPila.CrearPilaDinamica[TDAPila.Pila[int]]()

	t.Run("Pila vacia de pilas", func(t *testing.T) {
		require.True(t, pilaPilas.EstaVacia(), "Una pila recien creada debe estar vacia")
	})
}

func TestAccionesInvalidasEnPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "No hubo un panic error con ver tope en una pila vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "No se puede desapilar un elemento de pila vacia recien creada tirando panic error")
}

func TestApilarElementos(t *testing.T) {

	{
		pila := TDAPila.CrearPilaDinamica[int]()
		pila.Apilar(5)
		require.False(t, pila.EstaVacia(), "Una pila en el que se apilo un elemento no esta vacia")
		require.Equal(t, 5, pila.VerTope(), "El tope de una pila en la que se apilo un dato es el dato apilado")
	}
	{
		pila := TDAPila.CrearPilaDinamica[int]()
		apilarDatos([]int{6, 7, 9}, pila)
		require.Equal(t, 9, pila.VerTope(), "El tope de una pila en la que se apilaron varios datos es el ultimo dato apilado")
	}
	{
		pila := TDAPila.CrearPilaDinamica[int]()
		apilarDatos([]int{3, 13, 23, 46}, pila)
		pila.Apilar(2)
		require.Equal(t, 2, pila.VerTope(), "Si se apila un elemento nuevo a una pila con datos, su tope es el ultimo apilado")
	}

}

func TestDesapilarUnicoElemento(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(5)
	datoDesapilado := pila.Desapilar()
	require.Equal(t, 5, datoDesapilado, "Desapilar el unico dato de una pila devuelve el dato indicado")
	require.True(t, pila.EstaVacia(), "Desapilar un dato de una pila con solo un elemento deja una pila vacia")

}

func TestVerTope(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(6)
	require.Equal(t, 6, pila.VerTope(), "El tope de una pila donde se apilo un unico elemento es el elemento apilado")

	apilarDatos([]int{-13, 4, 7, -4, 11, 19}, pila)
	require.Equal(t, 19, pila.VerTope(), "El tope de una pila donde se apilaron varios elementos es el ultimo elemento apilado")

	require.Equal(t, pila.VerTope(), pila.Desapilar(), "El valor desapilado es el mismo que el del tope de la pila")

	apilarDatos([]int{6, 9, 15}, pila)
	topeAnterior := pila.VerTope()
	pila.Desapilar()
	require.False(t, pila.VerTope() == topeAnterior, "Desapilar un elemento modifica el tope de la pila")
	require.Equal(t, 9, pila.VerTope(), "El tope de una pila a la que se apilaron varios elementos y se desapilo el ultimo, es el anterior a este")

}

func TestApilarYDesapilarMantieneOrden(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	apilarDatos([]int{3, 8, 1, -4, 7}, pila)
	require.True(t, pilaDesapiladaEnOrden([]int{7, -4, 1}, pila), "Apilar varios datos y desapilarlos nos devuelve los datos en el orden indicado")

	apilarDatos([]int{12, 1}, pila)
	require.Equal(t, 1, pila.VerTope(), "El tope de una pila que fue desapilada hasta cierto punto y luego apilada es el ultimo dato apilado")

	require.True(t, pilaDesapiladaEnOrden([]int{1, 12}, pila), "Desapilar datos dejando una pila no vacia y luego apilar datos y desapilarlos nos devuelve los datos en orden indicado")

	var desapilado int

	for !pila.EstaVacia() {
		desapilado = pila.Desapilar()
	}

	require.Equal(t, 3, desapilado, "Desapilar una pila hasta que quede vacia devuelve el primer elemento que fue apilado")

	pila.Apilar(5)
	for i := 0; i < 1000; i++ {
		pila.Apilar((i + 5) * 2)
	}

	desapilarDatos(1000, pila)
	require.Equal(t, 5, pila.Desapilar(), "Desapilar una pila con muchos elementos hasta que quede vacia devuelve el primer elemento que fue apilado")

}

func TestVolumenYRedimension(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < 100000; i++ {
		pila.Apilar(i)
	}

	require.Equal(t, 99999, pila.VerTope(), "Ver el tope de una pila donde se apilaron una gran cantidad de elementos devuelve el ultimo apilado")

	desapilarDatos(100000, pila)

	require.True(t, pila.EstaVacia(), "Desapilamos todos los elementos de una pila con gran cantidad de datos nos deja una pila vacia y no demora mucho")

}

func TestVaciaLuegoDeDesapilar(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	apilarDatos([]int{3, 5, 7, 12, 2, 1, 14, 6, 7}, pila)
	desapilarDatos(9, pila)

	require.True(t, pila.EstaVacia(), "Una pila donde se desapilaron tantos elementos como los que fueron apilados queda vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Ver tope en la pila que fue vaciada sigue tirando panic error")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "No se puede seguir desapilando en la pila vaciada tirando un panic error")

	pila.Apilar(10)
	require.False(t, pila.EstaVacia(), "Apilar un elemento en la lista vaciada hace que esta no quede vacia")
	require.Equal(t, 10, pila.VerTope(), "El tope de la pila vaciada a la que se le apilo un elemento es el ultimo elemento apilado")
}

func TestPilaStrings(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[string]()

	apilarDatos([]string{"hey", "messi", "cuti", "Barbara", "Algoritmo"}, pila)
	require.True(t, pilaDesapiladaEnOrden([]string{"Algoritmo", "Barbara", "cuti"}, pila), "Apilar varios datos y desapilarlos nos devuelve los datos en el orden indicado")

	apilarDatos([]string{"fiuba", "pila"}, pila)
	require.Equal(t, "pila", pila.VerTope(), "El tope de una pila que fue desapilada hasta cierto punto y luego apilada es el ultimo dato apilado")
	require.True(t, pilaDesapiladaEnOrden([]string{"pila", "fiuba"}, pila), "Desapilar datos dejando una pila no vacia y luego apilar datos y desapilarlos nos devuelve los datos en orden indicado")

	var desapilado string

	for !pila.EstaVacia() {
		desapilado = pila.Desapilar()
	}

	require.Equal(t, "hey", desapilado, "Desapilar una pila hasta que quede vacia devuelve el primer elemento que fue apilado")

	pila.Apilar("choclo")
	for i := 0; i < 1000; i++ {
		pila.Apilar(strconv.Itoa(i))
	}

	desapilarDatos(1000, pila)
	require.Equal(t, "choclo", pila.Desapilar(), "Desapilar una pila con muchos elementos hasta que quede vacia devuelve el primer elemento que fue apilado")

}

func TestPilaDePilas(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[*TDAPila.Pila[int]]()

	var (
		pila1 = TDAPila.CrearPilaDinamica[int]()
		pila2 = TDAPila.CrearPilaDinamica[int]()
		pila3 = TDAPila.CrearPilaDinamica[int]()
		pila4 = TDAPila.CrearPilaDinamica[int]()
	)

	apilarDatos([]*TDAPila.Pila[int]{&pila1, &pila2, &pila3, &pila4}, pila)
	apilarDatos([]int{3, 16, 2}, pila1)
	require.True(t, pilaDesapiladaEnOrden([]*TDAPila.Pila[int]{&pila4, &pila3, &pila2}, pila), "Apilar varios datos y desapilarlos nos devuelve los datos en el orden indicado")

	desapilado := *pila.Desapilar()
	require.Equal(t, 2, desapilado.VerTope(), "Desapilar la ultima pila nos devuelve la primera pila apilada y se puede seguir operando con ella")

	require.True(t, pila.EstaVacia(), "Se puede desapilar una pila de pilas hasta que este vacia")

	apilarDatos([]*TDAPila.Pila[int]{&pila3, &pila4}, pila)
	require.Equal(t, &pila4, pila.VerTope(), "Se puede apilar elementos de una pila vaciada y ver tope nos devuelve el ultimo elemento apilado")
}
