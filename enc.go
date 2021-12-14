package gosharedlibs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func EncryptToken(t string, k string) string {

	token := []byte(t)
	key := []byte(k)

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("1. Encryption error.")
		os.Exit(-12)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println("2. Encryption error.")
		os.Exit(-14)
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println("3. Encryption error.")
		os.Exit(-15)
	}

	result := gcm.Seal(nonce, nonce, token, nil)

	return string(result)
}

func DecryptToken(data []byte, key []byte) string {

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(plaintext)
}

func DumpEncryptedToken(encryptedData []byte) {
	homeFolder, err := HomeFolder()
	if err != nil {
		return
	}
	outdirpath := fmt.Sprintf("%s/%s", homeFolder, ".tikitool")

	fMode := fs.FileMode(uint32(0700))
	err = os.MkdirAll(outdirpath, fs.FileMode(fMode))
	if err != nil {
		fmt.Println(err)
	}
	filename := "ticketool.token"
	fPath := filepath.Join(outdirpath, filename)

	fiLE, err := os.Create(fPath)
	if err != nil {
		fmt.Printf("File creation error : %s", err)
	}

	if _, err := fiLE.Write(encryptedData); err != nil {
		fmt.Printf("File writing error")
	}

	if err := fiLE.Close(); err != nil {
		fmt.Printf("File couldn't close properly.")
	}

}

func ExportEncryptedToken() ([]byte, error) {
	homeFolder, err := HomeFolder()
	if err != nil {
		return nil, err
	}

	outdirpath := fmt.Sprintf("%s%s", homeFolder, "/.ticketool")
	filename := "ticketool.token"
	fPath := filepath.Join(outdirpath, filename)

	tokeEncyptedData, err := ioutil.ReadFile(fPath)
	if err != nil {
		fmt.Printf("File read error: %+v \n", err)
	}

	return tokeEncyptedData, nil
}
