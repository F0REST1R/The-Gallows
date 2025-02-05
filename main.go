package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode/utf8"
)

var hangmanStages = []string{
	`
	_____
	|    |
	|
	|
	|
	|
	=========
	`, `
	_____
	|    |
	|    O
	|
	|
	|
	=========
	`, `
	_____
	|    |
	|    O
	|    |
	|
	|
	=========
	`, `
	_____
	|    |            
	|    O          
	|   /|       
	|           
	|           
	=========
	`, `
	_____
	|    |
	|    O
	|   /|\
	|
	|
	=========
	`, `
	_____
	|    |
	|    O
	|   /|\
	|   /
	|
	=========
	`, `
	_____
	|    |             АХУЕТЬ, ПИЗДА ЕМУ!!!!
	|    O          /
	|   /|\      o
	|   / \     /|\
	|           / \
	=========
	`,
}

var words = []string{
	"компьютер", "программа", "разработка", "интерфейс", "алгоритм", "массив", "переменная",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Добро пожаловать в игру 'Виселица'!")
		fmt.Println("1. Начать новую игру")
		fmt.Println("2. Выйти")
		fmt.Print("Выберите опцию: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			playGame(reader)
		case "2":
			fmt.Println("Выход из игры.")
			return
		default:
			fmt.Println("Неверный выбор, попробуйте снова.")
		}
	}
}

func playGame(reader *bufio.Reader) {
	word := words[rand.Intn(len(words))]
	guessedLetters := make(map[rune]bool)
	mistakes := 0
	maxMistakes := len(hangmanStages) - 1

	for {
		printState(word, guessedLetters, mistakes)
		if mistakes >= maxMistakes {
			fmt.Println("Вы проиграли! Загаданное слово:", word)
			return
		}

		if isWordGuessed(word, guessedLetters) {
			fmt.Println("Поздравляем! Вы выиграли! Слово:", word)
			return
		}

		fmt.Print("Введите букву: ")
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(guess)

		if utf8.RuneCountInString(guess) != 1 {
			fmt.Println("Пожалуйста, введите одну букву.")
			continue
		}

		letter, _ := utf8.DecodeRuneInString(guess)

		if guessedLetters[letter] {
			fmt.Println("Эта буква уже была угадана.")
			continue
		}

		guessedLetters[letter] = true

		if !strings.ContainsRune(word, letter) {
			mistakes++
		}
	}
}

func printState(word string, guessedLetters map[rune]bool, mistakes int) {
	fmt.Println(hangmanStages[mistakes])
	for _, letter := range word {
		if guessedLetters[letter] {
			fmt.Printf("%c ", letter)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
	fmt.Printf("Ошибки: %d/%d\n", mistakes, len(hangmanStages)-1)
}

func isWordGuessed(word string, guessedLetters map[rune]bool) bool {
	for _, letter := range word {
		if !guessedLetters[letter] {
			return false
		}
	}
	return true
}
