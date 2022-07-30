package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	inputFile := flag.String("input", "pw.txt", "明文パスワード")
	outputFile := flag.String("output", "pw-hash.txt", "暗号化パスワード")
	cost := flag.Int("cost", 4, "暗号化回数")
	flag.Parse()

	// fmt.Println(*inputFile)
	// fmt.Println(*outputFile)
	// fmt.Println(*cost)

	bytes, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	plainPasswordArray := strings.Split(string(bytes), "\n")

	var hashPasswordArray []string
	for _, password := range plainPasswordArray {
		// ハッシュしたものをskip
		if len(password) == 60 {
			hashPasswordArray = append(hashPasswordArray, password)
			continue
		}
		// 空白行
		if strings.TrimSpace(password) == "" {
			hashPasswordArray = append(hashPasswordArray, "")
			continue
		}
		// パスワードをハッシュ
		hash, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), *cost)
		if err != nil {
			log.Fatal(err)
		}
		hashPasswordArray = append(hashPasswordArray, string(hash))
	}

	os.WriteFile(*outputFile, []byte(strings.Join(hashPasswordArray, "\n")), 0644)
}
