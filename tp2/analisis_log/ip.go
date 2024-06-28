package tp2

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	_CUATRO = 4
)

type Ip [4]int

func ObtenerIpDeString(numeroIp string) Ip {

	digitos := strings.Split(numeroIp, ".")

	var nuevaIp Ip

	for i := range digitos {
		nuevaIp[i], _ = strconv.Atoi(digitos[i])
	}

	return nuevaIp
}

func ObtenerStringDeIp(ipAConvertir Ip) string {

	var digitos [_CUATRO]string

	for i, digito := range ipAConvertir {
		digitos[i] = strconv.Itoa(digito)
	}

	return fmt.Sprintf("%s.%s.%s.%s", digitos[0], digitos[1], digitos[2], digitos[3])
}

func CompararIps(numeroIp1 Ip, numeroIp2 Ip) int {

	diferenciaEntreIp := 0
	digitoAComparar := 0

	for diferenciaEntreIp == 0 && digitoAComparar < _CUATRO {
		diferenciaEntreIp = numeroIp1[digitoAComparar] - numeroIp2[digitoAComparar]
		digitoAComparar++
	}

	return diferenciaEntreIp
}
