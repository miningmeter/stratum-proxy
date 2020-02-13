package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

/*
ValidateAddr - валидация адреса.

@param string addr   адрес в формате addr:port.
@param bool   canDNS addr может быть DNS именем.

@return bool Если адрес валидный.
*/
func ValidateAddr(addr string, canDNS bool) bool {
	regex := "^[\\d\\.]+:\\d{1,5}$"
	if canDNS {
		regex = "^[^:]+:\\d{1,5}$"
	}
	matched, err := regexp.MatchString(regex, addr)
	if !matched || err != nil {
		return false
	}
	parts := strings.Split(addr, ":")
	if canDNS {
		if !ValidateDNS(parts[0]) && !ValidateIPV4(parts[0]) {
			return false
		}
	} else {
		if !ValidateIPV4(parts[0]) {
			return false
		}
	}
	if !ValidatePort(parts[1]) {
		return false
	}

	return true
}

/*
ValidateDNS - валидация адреса DNS.

@param string dns DNS-имя для проверки.

@return bool если имя верное.
*/
func ValidateDNS(dns string) bool {
	matched, err := regexp.MatchString("^([a-z0-9]+(-[a-z0-9]+)*\\.)+[a-z]{2,}$", dns)
	if !matched || err != nil {
		return false
	}

	return true
}

/*
ValidateIPV4 - валидация IPV4.

@param string ipv4 IP-адрес.

@return bool Если IP-адрес валидный.
*/
func ValidateIPV4(ipv4 string) bool {
	matched, err := regexp.MatchString("^(?:\\d{1,3}\\.){3}\\d{1,3}$", ipv4)
	if !matched || err != nil {
		return false
	}
	parts := strings.Split(ipv4, ".")
	if len(parts) != 4 {
		return false
	}
	for _, x := range parts {
		if i, err := strconv.Atoi(x); err == nil {
			if i < 0 || i > 255 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

/*
ValidatePort - валидация порта.

@return bool Если порт валидный.
*/
func ValidatePort(port string) bool {
	matched, err := regexp.MatchString("^\\d{1,5}$", port)
	if !matched || err != nil {
		return false
	}
	if i, err := strconv.Atoi(port); err == nil {
		if i < 1 || i > 65535 {
			return false
		}
	} else {
		return false
	}
	return true
}

/*
ValidateHexString - валидация 16-тиричной строки произвольной ненулевой длины.

@param string str строка для проверки.

@return bool Если строка валидна.
*/
func ValidateHexString(str string) bool {
	return rHexStr.MatchString(str)
}

/*
LogInfo - вывод информационной строки лога.
*/
func LogInfo(format, sid string, v ...interface{}) {
	header := make([]byte, 0, 0)
	if sid != "" {
		header = append(header, fmt.Sprintf("[%s] : ", sid)...)
	}
	if !syslog {
		header = append(header, "\033[0;32mI\033[0m : "...)
	} else {
		header = append(header, "I : "...)
	}

	header = append(header, format...)
	log.Printf(string(header), v...)
}

/*
LogError - вывод строки лога c ошибкой.
*/
func LogError(format, sid string, v ...interface{}) {
	header := ""
	if sid != "" {
		header += fmt.Sprintf("[%s] : ", sid)
	}
	if !syslog {
		header += "\033[0;31mE\033[0m : "
	} else {
		header += "E : "
	}
	log.Printf(header+format, v...)
}
