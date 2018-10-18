package basic

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	// 使用短除法 除 2 取模
	result := ""
	if n == 0 {
		return "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
	    result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileName string) {
	if file, err := os.Open(fileName); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
	
	
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(3231551),
		convertToBin(0),
	)
	
	printFile("/Users/jw.hu/Desktop/abc.txt")
}
