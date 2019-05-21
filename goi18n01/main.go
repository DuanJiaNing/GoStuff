package main

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	k1 = "%s went to %s."
	k2 = "%s has been stolen."
	k3 = "How are you?"

	// 1 增加 key
)

func init() {

	// 2 对应语言翻译

	// 初始化放入默认的 catalog
	message.SetString(language.Greek, k1, "%s πήγε στήν %s.")
	message.SetString(language.AmericanEnglish, k1, "%s is in %s.")

	message.SetString(language.Greek, k2, "%s κλάπηκε.")
	message.SetString(language.AmericanEnglish, k2, "%s has been stolen.")

	message.SetString(language.Greek, k3, "Πώς είστε?.")
}

func main() {
	p := message.NewPrinter(language.Greek)
	p.Printf(k1, "Ο Πέτρος", "Αγγλία")
	fmt.Println()

	p.Printf(k2, "Η πέτρα")
	fmt.Println()

	p = message.NewPrinter(language.AmericanEnglish)
	p.Printf(k1, "Peter", "England")
	fmt.Println()
	p.Printf(k2, "The Gem")
}
