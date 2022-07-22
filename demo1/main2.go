// Объявляет пакет, которому принадлежит код
package main

import (
	"fmt"
	"strings"
)

// Делает пакет fmt (формат) доступным для использования

// Объявляет функцию под названием main
func main() {
	// Выводит текст Hello, playground на экран
	//fmt.Println("Hello, playground")

	fmt.Println("Вы находитесь в темной пещере.")

	var command = "выйти наружу"
	var exit = strings.Contains(command, "наружу")

	fmt.Println("Вы покидаете пещеру:", exit)
}
