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
	// fmt.Printf("Enc Result: %s\n", string(result))

	return string(result)
}

func DecryptToken(data []byte, key []byte) string {

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1) // TODO: might not exit
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	// ioutil.ReadFile()x

	nonceSize := gcm.NonceSize()

	if len(data) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(plaintext))

	return string(plaintext)
}

func DumpEncryptedToken(encryptedData []byte) {
	// TODO: Check if $HOME/.ticketool directory exists and there is priv/pub key pair
	//  If not, create one

	// fmt.Printf("%+v\n", encryptedData)
	// fmt.Printf("%+v\n", string(encryptedData))

	outdirpath := fmt.Sprintf("%s%s", os.Getenv("HOME"), "/.tikiool")

	fMode := fs.FileMode(uint32(0700))
	err := os.MkdirAll(outdirpath, fs.FileMode(fMode))
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
	//fmt.Println("Existing token is processing...")
	outdirpath := fmt.Sprintf("%s%s", os.Getenv("HOME"), "/.ticketool")
	filename := "ticketool.token"
	fPath := filepath.Join(outdirpath, filename)

	tokeEncyptedData, err := ioutil.ReadFile(fPath)
	if err != nil {
		fmt.Printf("File read error: %+v \n", err)
	}

	return tokeEncyptedData, nil
}
