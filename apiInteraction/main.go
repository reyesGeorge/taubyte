package main

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"

	"fmt"
	"os"

	shell "github.com/ipfs/go-ipfs-api"

	"github.com/gin-gonic/gin"
	// "crypto/aes"
	// "crypto/cipher"
	// "crypto/rand"
	// "encoding/base64"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/:hash", HashGrabber)
	router.POST("/:data", StoreNReadString)

	router.Run("localhost:8000")
}

func HashGrabber(c *gin.Context) {

	hash := c.Param("hash")
	// key := []byte("123456789012345678901234")

	sh := shell.NewShell("localhost:5001")
	cid, err := sh.BlockGet(hash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("added %s", cid)
	fmt.Printf("%T\n", cid)

	// newCid := string(cid)

	// encrypt base64 crypto to original value
	// text := decrypt(key, cid)
	// fmt.Printf(text)

	c.JSON(http.StatusOK, string(cid))

}

func StoreNReadString(c *gin.Context) {

	data := c.Param("data")

	// key := []byte("1234561234561234561234561234561234561234561234")

	// encrypt value to base64
	// cryptoText := encrypt(key, data)

	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader(data))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	// HashGrabber(cid)
	c.JSON(http.StatusOK, string(cid))
}

// // encrypt string to base64 crypto using AES
// func encrypt(key []byte, text string) string {
// 	// key := []byte(keyText)
// 	plaintext := []byte(text)

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// The IV needs to be unique, but not secure. Therefore it's common to
// 	// include it at the beginning of the ciphertext.
// 	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
// 	iv := ciphertext[:aes.BlockSize]
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		panic(err)
// 	}

// 	stream := cipher.NewCFBEncrypter(block, iv)
// 	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

// 	// convert to base64
// 	return base64.URLEncoding.EncodeToString(ciphertext)
// }

// // decrypt from base64 to decrypted string
// func decrypt(key []byte, cryptoText string) string {
// 	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// The IV needs to be unique, but not secure. Therefore it's common to
// 	// include it at the beginning of the ciphertext.
// 	if len(ciphertext) < aes.BlockSize {
// 		panic("ciphertext too short")
// 	}
// 	iv := ciphertext[:aes.BlockSize]
// 	ciphertext = ciphertext[aes.BlockSize:]

// 	stream := cipher.NewCFBDecrypter(block, iv)

// 	// XORKeyStream can work in-place if the two arguments are the same.
// 	stream.XORKeyStream(ciphertext, ciphertext)

// 	return fmt.Sprintf("%s", ciphertext)

// }
