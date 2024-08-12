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
	21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV",
	26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX",
	31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV",
	36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL",
	41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV",
	46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L",
	51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV",
	56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX",
	61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV",
	66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX",
	71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV",
	76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX",
	81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV",
	86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
	91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV",
	96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
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
		result = first * second

	case "/":
		result = first / second
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
		result = first * second

	case "/":
		result = first / second
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
