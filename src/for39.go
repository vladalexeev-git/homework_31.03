package src

import (
	"fmt"
	"log"

	"github.com/vladalexeev-git/homework_31.03/tools"
)

//Функция выводит целые числа от A до B включительно; при этом каждое число выводиться
//	столько раз, каково его значение (например, число 3 выводится 3 раза)
func PrintValueNum() {
	fmt.Println(tools.Desc39)
	var a, b int

	fmt.Println("Введите A и B: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		log.Println(err)
	}
	if a <= 0 || b <= 0 {
		fmt.Println("Числа не подходят")
		return
	}
	
	for ;a<=b; a++{
		for j := 0; j<a; j++ {
			fmt.Print(a)
		}
		
	}
	fmt.Println("")
}