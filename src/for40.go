package src

import (
	"fmt"
	"log"

	"github.com/vladalexeev-git/homework_31.03/tools"
)

//Функция выводит целые числа от A до B включительно; при этом число A выводиться 1 раз, число A + 1 - 2 раза и т. д.
func PrintValueNumPlusNum() {
	fmt.Println(tools.Desc40)
	var a, b int
	fmt.Println("Введите A и B:")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		log.Println(err)
	}
	
	for i:=0;a<=b; a++{
		i++
		for j := 0; j<i; j++ {
			fmt.Print(a)
		}
		
	}
	fmt.Println("")
}