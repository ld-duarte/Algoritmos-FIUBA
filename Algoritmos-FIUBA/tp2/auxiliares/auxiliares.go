package auxiliares

import (
	"strconv"
	"strings"
)

func EsIp(ip string) bool {
	bloques := strings.Split(ip, ".")
	if len(bloques) != 4 {
		return false
	}
	for _, bloque := range bloques {
		bloqueInt, err := strconv.Atoi(bloque)
		if err != nil || bloqueInt < 0 || bloqueInt > 255 {
			return false
		}
	}
	return true
}

func CompararIpsMax(ip1, ip2 string) int {
	bloques1 := strings.Split(ip1, ".")
	bloques2 := strings.Split(ip2, ".")

	for i := 0; i < 4; i++ {
		num1, _ := strconv.Atoi(bloques1[i])
		num2, _ := strconv.Atoi(bloques2[i])
		if num1 > num2 {
			return 1
		}
		if num1 < num2 {
			return -1
		}
	}
	return 0
}

func CompararIpsMin(ip1, ip2 string) int {
	return CompararIpsMax(ip2, ip1)
}
