package main

import (
	"fmt"
)

type bookmarkMap = map[string]string

const DefaultSize = 4

func main() {
	data := make(bookmarkMap, DefaultSize)
Menu:
	for {
		switch takeChoice() {
		case "1":
			printMarks(data)
		case "2":
			addMark(data)
		case "3":
			deleteMark(data)
		case "4":
			break Menu
		default:
			fmt.Println("Неизвестный пункт программы")
		}
	}
}

func takeChoice() string {
	var choice string
	fmt.Println("Меню:")
	fmt.Println("- 1. Посмотреть закладки")
	fmt.Println("- 2. Добавить закладку")
	fmt.Println("- 3. Удалить закладку")
	fmt.Println("- 4. Выход")
	fmt.Scan(&choice)
	return choice
}

func printMarks(data bookmarkMap) {
	fmt.Println("Закладки:")
	for key, value := range data {
		fmt.Printf("Закладка ключ %s значение %s\n", key, value)
	}
}

func addMark(data bookmarkMap) {
	var key, value string
	fmt.Print("Введите ключ: ")
	fmt.Scan(&key)
	fmt.Print("Введите значение: ")
	fmt.Scan(&value)
	data[key] = value
	fmt.Println("Закладка добавлена!")
}

func deleteMark(data bookmarkMap) {
	var key string
	fmt.Print("Введите какую закладку удалить: ")
	fmt.Scan(&key)
	delete(data, key)
	fmt.Println("Закладка удалена!")
}
