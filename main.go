package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"
)

type UserDetails struct {
	Name   string `json:"Name"`
	Addr   string `json:"Address"`
	Mobile string `json:"Mobile"`
	Email  string `json:"Email"`
}

func generateQR(info UserDetails) {
	var content strings.Builder
	var fileName string

	/*
		BEGIN:VCARD
		VERSION:4.0
		FN:{Name}
		ADR:{Addr}
		TEL:{Mobile}
		EMAIL:{Email}
		END:VCARD
	*/

	content.WriteString("BEGIN:VCARD\n")
	content.WriteString("VERSION:4.0\n")
	content.WriteString("FN:" + info.Name + "\n")
	content.WriteString("ADR:" + info.Addr + "\n")
	content.WriteString("TEL:" + info.Mobile + "\n")
	content.WriteString("EMAIL:" + info.Email + "\n")
	content.WriteString("END:VCARD\n")

	fileName = "images/" + info.Name + "-qr.png"

	err := qrcode.WriteFile(content.String(), qrcode.Medium, 100, fileName)

	if err != nil {
		log.Fatal(err)
	}

}

func generateQRRequest(c *gin.Context) {
	var info UserDetails
	if err := c.BindJSON(&info); err != nil {
		log.Fatal(err)
	}
	if len(info.Name) == 0 || len(info.Mobile) == 0 || len(info.Addr) == 0 || len(info.Email) == 0 {
		msg := "ERROR: Incomplete contact details were provided."
		c.IndentedJSON(http.StatusInternalServerError, msg)
		return
	}
	fmt.Println(info)
	generateQR(info)

	c.IndentedJSON(http.StatusOK, info)
}

func main() {
	router := gin.Default()

	router.GET("/generateQR", generateQRRequest)

	router.Run("localhost:8080")

}
