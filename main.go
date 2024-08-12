package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputLine := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение")
		input, _ := inputLine.ReadString('\n')
		input = strings.TrimSpace(input)

		result, err := calculator(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(1)
		}

		fmt.Println("Результат:", result)
	}

}

var listOFRomanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var listOFArabicNumerals = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
	6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV",
	16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
}

func calculator(input string) (string, error) {

	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		return "", errors.New("некорректный формат ввода")
	}

	firstNum, secondNum := tokens[0], tokens[2]
	mathOperation := tokens[1]

	if isRoman(firstNum) && isRoman(secondNum) {
		return calculateRoman(firstNum, secondNum, mathOperation)
	} else if isArabic(secondNum) && isArabic(firstNum) {
		return calculateArabic(firstNum, secondNum, mathOperation)
	} else {
		return "", errors.New("используются одновременно разные системы счисления")
	}

}

func calculateArabic(firstNum, secondNum, mathOperation string) (string, error) {
	var result int
	first, err := strconv.Atoi(firstNum)
	if err != nil {
		return "", errors.New("некорректный ввод числа")
	}

	second, err := strconv.Atoi(secondNum)
	if err != nil {
		return "", errors.New("некорректный ввод числа")
	}

	if first < 1 || first > 10 || second < 1 || second > 10 {
		return "", errors.New("числа должны быть от 1 до 10")
	}

	switch mathOperation {

	case "+":
		result = first + second

	case "-":
		result = first - second

	case "*":
		result = first + second

	case "/":
		result = first + second
	default:
		return "", errors.New("некорректная операция")
	}

	return strconv.Itoa(result), nil

}

func calculateRoman(firstNum, secondNum, mathOperation string) (string, error) {

	first, err := listOFRomanNumerals[firstNum]
	if !err {
		return "", errors.New("некорректный ввод римского числа")
	}

	second, err := listOFRomanNumerals[secondNum]
	if !err {
		return "", errors.New("некорректный ввод римского числа")
	}

	if first < 1 || second > 10 || second < 1 || first > 10 {
		return "", errors.New("числа должны быть от 1 до 10")
	}

	var result int
	switch mathOperation {
	case "+":
		result = first + second

	case "-":
		result = first - second

	case "*":
		result = first + second

	case "/":
		result = first + second
	default:
		return "", errors.New("некорректная операция")
	}

	if result < 1 {
		return "", errors.New("результат меньше единицы в римской системе")
	}

	return listOFArabicNumerals[result], nil
}

func isRoman(s string) bool {
	_, ok := listOFRomanNumerals[s]
	return ok
}

func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
