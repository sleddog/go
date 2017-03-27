/*
https://projecteuler.net/problem=59
XOR decryption

Each character on a computer is assigned a unique code and the preferred standard is ASCII (American Standard Code for Information Interchange). For example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.

A modern encryption method is to take a text file, convert the bytes to ASCII, then XOR each byte with a given value, taken from a secret key. The advantage with the XOR function is that using the same encryption key on the cipher text, restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.

For unbreakable encryption, the key is the same length as the plain text message, and the key is made up of random bytes. The user would keep the encrypted message and the encryption key in different locations, and without both "halves", it is impossible to decrypt the message.

Unfortunately, this method is impractical for most users, so the modified method is to use a password as a key. If the password is shorter than the message, which is likely, the key is repeated cyclically throughout the message. The balance for this method is using a sufficiently long password key for security, but short enough to be memorable.

Your task has been made easy, as the encryption key consists of three lower case characters. Using cipher.txt (right click and 'Save Link/Target As...'), a file containing the encrypted ASCII codes, and the knowledge that the plain text must contain common English words, decrypt the message and find the sum of the ASCII values in the original text.
*/

package main

import (
	"./lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("problem 59")
	/*
		text := 65
		key := 42
		cipher := xor(text, key)
		fmt.Println("text=", text, "key=", key, "cipher=", cipher)
		fmt.Println(charCode("s"))

		a := 65
		fmt.Println("a=", a, "--", codeChar(a))
		a = 42
		fmt.Println("a=", a, "--", codeChar(a))
	*/

	//read file
	lines, _ := lib.GetLines("./p059_cipher.txt")
	//fmt.Println(lines)
	//[x,y,...,n]
	numbers := parseNumbers(lines[0])
	fmt.Println(numbers)
	fmt.Println("len=", len(numbers))

	//enumerate all lowercase letters
	for k1 := 97; k1 < 123; k1++ {
		for k2 := 97; k2 < 123; k2++ {
			for k3 := 97; k3 < 123; k3++ {
				key := codeChar(k1) + codeChar(k2) + codeChar(k3)
				decryptedText := tryKey(key, numbers)
				if testEnglish(decryptedText) {
					fmt.Println("key=", key)
					fmt.Println(decryptedText)
				}
			}
		}
	}

	//found god! LOL
	originalText := tryKey("god", numbers)
	fmt.Println(originalText)
	sum := 0
	for _, v := range originalText {
		sum += charCode(string(v))
	}
	fmt.Println("sum=", sum)

}

func testEnglish(words string) bool {
	//try simple tests for english words
	the := strings.Contains(words, "the")
	and := strings.Contains(words, "and")
	is := strings.Contains(words, "is")
	return the && and && is
}

func keyNumbers(key string, length int) []int {
	//repeat the key to fill up the original text
	repeatCount := length / len(key)
	keyNumbers := []int{}
	for i := 0; i < repeatCount; i++ {
		for _, v := range key {
			keyNumbers = append(keyNumbers, charCode(string(v)))
		}

	}
	remainder := length % len(key)
	for j := 0; j < remainder; j++ {
		keyNumbers = append(keyNumbers, charCode(string(key[j])))
	}
	return keyNumbers
}

func tryKey(key string, encryptedNumbers []int) string {
	keyNumbers := keyNumbers(key, len(encryptedNumbers))

	//for each number in the encrypted numbers, xor with key number, and convert to a string
	text := ""
	for i := 0; i < len(encryptedNumbers); i++ {
		xor := xor(encryptedNumbers[i], keyNumbers[i])
		t := codeChar(xor)
		text += t
	}
	return text
}

func parseNumbers(s string) []int {
	n := []int{}
	for _, a := range strings.Split(s, ",") {
		i, _ := strconv.Atoi(a)
		n = append(n, i)
	}
	return n
}

func xor(t, k int) int {
	return t ^ k
}

func charCode(s string) int {
	c := []rune(s)[0]
	return int(c)
}

func codeChar(i int) string {
	return string(i)

}
