package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	MFILEPATH = "meat"
	VFILEPATH = "vege"
)
func main() {
	e := gin.Default()

	e.GET("/", randomOrders)
	log.Println(http.ListenAndServe(":7777", e).Error())
}

func randomOrders(c *gin.Context) {

	mF := make([]string, 0)
	vF := make([]string, 0)
	mFile, err := os.Open(MFILEPATH)
	if err != nil {
		log.Printf("file: %s not found err %s", MFILEPATH, err)
		c.JSON(http.StatusInternalServerError, "open order file err")
		return
	}
	defer mFile.Close()
	vFile, err := os.Open(VFILEPATH)
	if err != nil {
		log.Printf("file: %s not found err %s", VFILEPATH, err)
		c.JSON(http.StatusInternalServerError, "open order file err")
		return
	}
	defer vFile.Close()

	mbr := bufio.NewReader(mFile)
	for {
		b, _, c := mbr.ReadLine()

		if c == io.EOF {
			break
		}

		mF = append(mF, string(b))
	}


	vbr := bufio.NewReader(vFile)
	for {
		b, _, c := vbr.ReadLine()

		if c == io.EOF {
			break
		}

		vF = append(vF, string(b))
	}

	ml := len(mF)
	vl := len(vF)

	rand.Seed(time.Now().Unix())

	res := make(map[string][]string)

	res["meat"] = make([]string, 0)

	res["meat"] = append(res["meat"], mF[rand.Intn(ml)], mF[rand.Intn(ml)], mF[rand.Intn(ml)])

	res["vege"] = make([]string, 0)

	res["vege"] = append(res["vege"], vF[rand.Intn(vl)])


	c.JSON(http.StatusOK, res)
	return
}
