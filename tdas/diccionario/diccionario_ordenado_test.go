package diccionario_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	TDADiccionarioOrdenado "tdas/diccionario"

	"github.com/stretchr/testify/require"
)

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que el diccionario vacío no tiene claves")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que el diccionario con un elemento tiene esa clave únicamente")
	dic := TDADiccionarioOrdenado.CrearABB[string, int](funcion_cmp)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardaar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario y comprueba su correcto funcionamiento")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	for i := 0; i < len(claves); i++ {
		require.False(t, dic.Pertenece(claves[i]))
		dic.Guardar(claves[i], valores[i])
		require.EqualValues(t, i+1, dic.Cantidad())
		require.True(t, dic.Pertenece(claves[i]))
		require.EqualValues(t, valores[i], dic.Obtener(claves[i]))
	}
}

func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos elementos en el diccionario y luego los borra")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	clave1 := "Rojo"
	clave2 := "Verde"
	clave3 := "Azul"
	valor1 := "#FF0000"
	valor2 := "#00FF00"
	valor3 := "#0000FF"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(t, len(claves), dic.Cantidad())

	for i := 0; i < len(claves); i++ {
		require.True(t, dic.Pertenece(claves[i]))
		dic.Borrar(claves[i])
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[i]) })
		require.EqualValues(t, len(claves)-(i+1), dic.Cantidad())
	}
}

func TestDiccionarioActualizar(t *testing.T) {
	t.Log("Guarda un elemento en el diccionario y luego lo actualiza")
	dic := TDADiccionarioOrdenado.CrearABB[string, int](funcion_cmp)
	clave := "Edad"
	valorOriginal := 25
	valorActualizado := 30

	dic.Guardar(clave, valorOriginal)
	require.EqualValues(t, valorOriginal, dic.Obtener(clave))

	dic.Guardar(clave, valorActualizado)
	require.EqualValues(t, valorActualizado, dic.Obtener(clave))
}

func TestReutlizacionDeBoorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionarioOrdenado.CrearABB[int, string](funcion_cmp)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dic := TDADiccionarioOrdenado.CrearABB[avanzado, int](funcion_cmp)

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionarioOrdenado.CrearABB[string, *int](funcion_cmp)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := TDADiccionarioOrdenado.CrearABB[string, *int](funcion_cmp)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionarioOrdenado.CrearABB[string, int](funcion_cmp)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionarioOrdenado.CrearABB[string, int](funcion_cmp)
	dic.Guardar(clave0, 7)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	dic.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := TDADiccionarioOrdenado.CrearABB[string, int](funcion_cmp)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])
	iter := dic.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	claves := []string{"A", "B", "C"}
	dic.Guardar(claves[0], "")
	dic.Guardar(claves[1], "")
	dic.Guardar(claves[2], "")

	dic.Iterador()
	iter2 := dic.Iterador()
	iter2.Siguiente()
	iter3 := dic.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestPruebaIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionarioOrdenado.CrearABB[int, int](funcion_cmp)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIterarConRango(t *testing.T) {
	t.Log("Iterar con un rango específico debería incluir solo los elementos dentro del rango")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionarioOrdenado.CrearABB[string, int](funcion_cmp)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	// Definimos el rango de claves desde "Gato" hasta "Vaca"
	desde := "Gato"
	hasta := "Vaca"

	clavesEnRango := []string{}

	dic.IterarRango(&desde, &hasta, func(clave string, _ int) bool {
		clavesEnRango = append(clavesEnRango, clave)
		return true
	})

	expectedClaves := []string{"Gato", "Hamster", "Perro", "Vaca"}
	require.ElementsMatch(t, expectedClaves, clavesEnRango)
}

func TestIterarSinRango(t *testing.T) {
	t.Log("Iterar sin un rango específico debería incluir todos los elementos del diccionario")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := TDADiccionarioOrdenado.CrearABB[string, int](funcion_cmp)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	todasLasClaves := []string{}

	dic.IterarRango(nil, nil, func(clave string, _ int) bool {
		todasLasClaves = append(todasLasClaves, clave)
		return true
	})

	expectedClaves := []string{"Burrito", "Gato", "Hamster", "Perro", "Vaca"}
	require.ElementsMatch(t, expectedClaves, todasLasClaves)
}
func TestIterarRangoSinHasta(t *testing.T) {
	t.Log("Iterar poniendo un rango de elementos pero sin un hasta (=nil)")
	dic := TDADiccionarioOrdenado.CrearABB[int, int](funcion_cmp)
	claves := []int{1, 2, 3, 4, 5}
	for _, clave := range claves {
		dic.Guardar(clave, clave)
	}

	desde := 3

	clavesEnRango := []int{}

	dic.IterarRango(&desde, nil, func(clave int, _ int) bool {
		clavesEnRango = append(clavesEnRango, clave)
		return true
	})

	expectedClaves := []int{3, 4, 5}
	require.ElementsMatch(t, expectedClaves, clavesEnRango)
}

func TestIteradorEnOrden(t *testing.T) {
	dic := TDADiccionarioOrdenado.CrearABB[int, string](funcion_cmp)
	claves := []int{5, 3, 7, 2, 4, 6, 8}
	for _, clave := range claves {
		dic.Guardar(clave, fmt.Sprintf("valor%d", clave))
	}

	// Iterar sin rango
	iter := dic.Iterador()
	var resultados []int
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultados = append(resultados, clave)
		iter.Siguiente()
	}
	expected := []int{2, 3, 4, 5, 6, 7, 8}
	require.EqualValues(t, expected, resultados)

	// Iterar con rango (desde=3, hasta=6)
	resultados = []int{}
	iter = dic.IteradorRango(&claves[1], &claves[5])
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultados = append(resultados, clave)
		iter.Siguiente()
	}
	expected = []int{3, 4, 5, 6}
	require.EqualValues(t, expected, resultados)

	// Iterar con hasta == nil (desde=3)
	resultados = []int{}
	iter = dic.IteradorRango(&claves[1], nil)
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		resultados = append(resultados, clave)
		iter.Siguiente()
	}
	expected = []int{3, 4, 5, 6, 7, 8}
	require.EqualValues(t, expected, resultados)
}

func TestIterarRango(t *testing.T) {
	diccionario := TDADiccionarioOrdenado.CrearABB[int, string](funcion_cmp)

	claves := []int{5, 3, 7, 2, 4, 6, 8}
	for _, clave := range claves {
		diccionario.Guardar(clave, "valor")
	}

	var resultados []int
	diccionario.IterarRango(nil, nil, func(clave int, valor string) bool {
		resultados = append(resultados, clave)
		return true
	})
	require.Equal(t, []int{2, 3, 4, 5, 6, 7, 8}, resultados)

	resultados = []int{}
	diccionario.IterarRango(&claves[1], &claves[5], func(clave int, valor string) bool {
		resultados = append(resultados, clave)
		return true
	})
	require.Equal(t, []int{3, 4, 5, 6}, resultados)

	resultados = []int{}
	diccionario.IterarRango(&claves[1], nil, func(clave int, valor string) bool {
		resultados = append(resultados, clave)
		return true
	})
	require.Equal(t, []int{3, 4, 5, 6, 7, 8}, resultados)
}

func funcion_cmp[K comparable](clave1, clave2 K) int {
	tipoClave := reflect.TypeOf(clave1)

	switch tipoClave.Kind() {
	case reflect.String:
		c1 := reflect.ValueOf(clave1).String()
		c2 := reflect.ValueOf(clave2).String()
		return strings.Compare(c1, c2)
	case reflect.Int:
		c1 := reflect.ValueOf(clave1).Int()
		c2 := reflect.ValueOf(clave2).Int()
		if c1 < c2 {
			return -1
		} else if c1 > c2 {
			return 1
		}
		return 0
	}
	return 0
}
