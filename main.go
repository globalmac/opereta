package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"flag"
	"strconv"
	"os"
	"encoding/json"
)

// The structure for the response in JSON format
type JsonResponse struct {
	Operation  string  `json:"operation"`
	Data       string  `json:"data"`
}

// Hide usage hints with empty arguments
var Usage = func() {
	os.Exit(1)
}

var phone     = flag.Int64("encrypt", 0, "Phone is required!")
var hash      = flag.String("decrypt", "", "Hash is required!")
var wantsJson = flag.Bool("json", false, "Output as JSON")

/* === Some helpers for base64 & AES Padding === */

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}
	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func cipherPad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func cipherUnpad(src []byte) ([]byte) {
	length := len(src)
	unPadding := int(src[length-1])
	if unPadding > length {
		os.Exit(1)
	}
	return src[:(length - unPadding)]
}

/* === Basic encryption and decryption functions === */

// Encrupt phone number
func encrypt(key []byte, phone string) (string) {

	// Create and return a new AES cipher
	block, err := aes.NewCipher(key)

	if err != nil {
		os.Exit(1)
	}

	// AES padding
	msg := cipherPad([]byte(phone))
	cipherText := make([]byte, aes.BlockSize+len(msg))
	iv := cipherText[:aes.BlockSize]

	// Read block in buffer
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		os.Exit(1)
	}

	// Read the cipher stream
	cfb := cipher.NewCFBEncrypter(block, iv)

	// Encrypt the bytes
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(msg))

	// Clear base64
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(cipherText))
	return finalMsg
}

// Encrypt hash phone number
func decrypt(key []byte, hash string) (string) {

	// Create and return a new AES cipher
	block, err := aes.NewCipher(key)

	if err != nil {
		os.Exit(1)
	}

	decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(hash))

	if err != nil {
		os.Exit(1)
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		os.Exit(1)
	}

	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(msg, msg)

	unpadMsg := cipherUnpad(msg)

	return string(unpadMsg)
}

func init() {

	flag.Parse()

}

func main() {

	// Encryption key - need to change!!! 
	key := []byte("XCnnvCRSds47EfJS9Q6N4aeqkdCp2tnE")

	if int64(*phone) > 0 {

		// Encrypt phone number

		encryptText := encrypt(key, strconv.Itoa(int(*phone)))

		if *wantsJson {

			data, _ := json.Marshal(&JsonResponse{
				Operation  :  "encrypt",
				Data       :  string(encryptText),
			})

			fmt.Println(string(data))

		} else {

			fmt.Println(encryptText)

		}

		os.Exit(0)

	} else if *hash != "" {

		// Decrypt phone number hash

		decryptText := decrypt(key, *hash)

		if *wantsJson {

			data, _ := json.Marshal(&JsonResponse{
				Operation  :  "decrypt",
				Data       :  string(decryptText),
			})

			fmt.Println(string(data))

		} else {

			fmt.Println(decryptText)

		}

		os.Exit(0)

	}

	os.Exit(1)

}
