package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"io"
	"log"
	"net/http"
)

type NewCryptoPage struct {
	Title string
	New string
	Encrypt_text string
	Decrypt_text string
	İndex_list []string
}

type Tag struct {
	ID   int    "json:'id'"
	Baslik string "json:'baslik'"
	İcerik string "json:'icerik'"
	Sifreli_metin string "json:'sifreli_metin'"
	Anahtar string "json:'sifreli_metin'"
	Sifresiz_metin string "json:'Decod_Text'"
	index_dizi[] string "json:'index_dizi'"
}

var tag Tag
var encrypt_value string
var decrypt_1 string
var CIPHER_KEY =[]byte("0123456789012345")
var CIPHER_KEY_0 =[]byte("0000000000000000")
var CIPHER_KEY_1 =[]byte("1111111111111111")

func main() {
	database()
	processer()
	web_server()
}
func processer()  {
	add_database(CIPHER_KEY_0)
	add_database(CIPHER_KEY_1)
	add_database(CIPHER_KEY)
	deneme(CIPHER_KEY_0)
	deneme(CIPHER_KEY_1)
	deneme(CIPHER_KEY)
}
func web_server()  {
	http.HandleFunc("/", newsCryptoHandler)
	http.HandleFunc("/about", index_handler)
	http.ListenAndServe(":8080", nil) // listen on what port?   ... can serve on tls with ListenAndServeTLS ... need something in server args, we'll put nil for now

}
func add_database(CIPHER_KEY []byte)  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/cryptology")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	var msg_1="Melik KUTLUAY"
		for j:=0;j<=3 ;j++  {
			if msg_1=="Melik KUTLUAY" {
				msg_1="Melik KUTLUAY"
				encrypt_1 := printer(CIPHER_KEY,msg_1)
				decrypt_1,_ = decrypt(CIPHER_KEY,msg_1)
				//fmt.Println("decrypt_1:000:",decrypt_1)
				insert, err := db.Query("INSERT INTO blog (baslik,icerik,sifreli_metin,cipher_key) VALUES('Seneryo_1','"+msg_1+"','"+encrypt_1+"','"+string(CIPHER_KEY)+"')")
				if err!=nil {
					panic(err.Error())
				}
				defer insert.Close()
				msg_1="aaaaaaaaaa"
			}else if msg_1=="aaaaaaaaaa" {
				msg_1="aaaaaaaaaa"
				encrypt_1 := printer(CIPHER_KEY,msg_1)
				decrypt_1,_ = decrypt(CIPHER_KEY,msg_1)
				//fmt.Println("decrypt_1:000",decrypt_1)
				insert, err := db.Query("INSERT INTO blog (baslik,icerik,sifreli_metin,cipher_key) VALUES('Seneryo_2','"+msg_1+"','"+encrypt_1+"','"+string(CIPHER_KEY)+"')")
				if err!=nil {
					panic(err.Error())
				}
				defer insert.Close()
				msg_1="aaaaabaaaa"
			}else if msg_1=="aaaaabaaaa" {
				msg_1="aaaaabaaaa"
				encrypt_1 := printer(CIPHER_KEY,msg_1)
				decrypt_1,_ = decrypt(CIPHER_KEY,msg_1)
				//fmt.Println("decrypt_1:000",decrypt_1)
				insert, err := db.Query("INSERT INTO blog (baslik,icerik,sifreli_metin,cipher_key) VALUES('Seneryo_3','"+msg_1+"','"+encrypt_1+"','"+string(CIPHER_KEY)+"')")
				if err!=nil {
					panic(err.Error())
				}
				defer insert.Close()
				msg_1="Zor kaderi geriye çevirmen"
			}else {
				msg_1="Zor kaderi geriye çevirmen"
				encrypt_1 := printer(CIPHER_KEY,msg_1)
				decrypt_1,_ = decrypt(CIPHER_KEY,msg_1)
				//fmt.Println("decrypt_1:000",decrypt_1)
				insert, err := db.Query("INSERT INTO blog (baslik,icerik,sifreli_metin,cipher_key) VALUES('Seneryo_4','"+msg_1+"','"+encrypt_1+"','"+string(CIPHER_KEY)+"')")
				if err!=nil {
					panic(err.Error())
				}
				defer insert.Close()
			}
		}
}
func deneme(key []byte,) {
	if decrypt_metin,err:=decrypt(key,tag.Sifreli_metin);err!=nil{
		log.Println(err)
		fmt.Println("sonuçç:",decrypt_metin)
	}else {
		log.Printf("Decrypted: %s\n",decrypt_metin)
	}
	return
}
func database()  {

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/cryptology")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	fmt.Println("succesfully connected mysql database :)")

	results,_:=db.Query("SELECT * FROM blog")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.Baslik, &tag.İcerik, &tag.Sifreli_metin,&tag.Anahtar)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
}
func newsCryptoHandler(w http.ResponseWriter, r *http.Request) {

	p:=NewCryptoPage{Title:tag.Baslik,New:tag.İcerik,Encrypt_text:tag.Sifreli_metin,Decrypt_text:tag.Sifresiz_metin}
	t,_:=template.ParseFiles("template/index.html")
	t.Execute(w,p)

}
func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to About Page :)</h1>")
}
func printer(CIPHER_KEY []byte, msg string)string {
	if encrypted,err := encrypt(CIPHER_KEY, msg);err!=nil{
		log.Println(err)
	}else {
		log.Printf("Cipher Key: %s\n",CIPHER_KEY)
		log.Printf("Encrypted: %s\n",encrypted)
		encrypt_value=encrypted
		if decrypted,err := decrypt(CIPHER_KEY,encrypted);err!=nil{
			log.Println(err)
		}else {
			log.Printf("Decrypted: %s\n",decrypted)
			//decrypt_value=decrypted
		}
	}
	return encrypt_value
}
func encrypt(key []byte, message string) (encmess string, err error)  {
	fmt.Println("message :",message,"\n")
	plainText:=[]byte(message)
	block,err:=aes.NewCipher(key)
	if err !=nil{
		return
	}
	cipherText:=make([]byte,aes.BlockSize+len(plainText))
	iv:=cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader,iv); err != nil {
		return
	}
	stream:=cipher.NewCFBEncrypter(block,iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:],plainText)
	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}
func decrypt(key []byte, securemess string) (decodedmess string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Cipher text block size too short !")
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	decodedmess = string(cipherText)
	fmt.Println("decoder message:",decodedmess)
	return
}
