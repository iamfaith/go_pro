package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	//_ "github.com/mattn/go-sqlite3"
	//"database/sql"
	_ "encoding/json"
	"os"
)

var (
	memoryDb          *bolt.DB
	QuestionBucket = "Question"
)


//func readSqlite() {
//	db, err := sql.Open("sqlite3", "./data.db")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	rows, err := db.Query("select quiz, answer from questions")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var quiz string
//		var ans string
//		err = rows.Scan(&quiz, &ans)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(quiz, ans)
//	}
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func initDb() {
	var err error
	memoryDb, err = bolt.Open("questions.data", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	memoryDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(QuestionBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func ShowAllQuestions() {
	var kv = map[string]string{}
	memoryDb.View(func(tx *bolt.Tx) error {

		userFile := "data.txt"
		fout,err := os.Create(userFile)
		defer fout.Close()
		if err != nil {
			fmt.Println(userFile,err)
		}

		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(QuestionBucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {


			//fmt.Printf("%s %s\n", k, v)
			fmt.Fprintf(fout, "%s %s\n", k, v)
			kv[string(k)] = string(v)
		}
		return nil
	})

}


func CountQuestions() int {
	var i int
	memoryDb.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(QuestionBucket))
		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			i++
		}
		return nil
	})
	fmt.Println(i)
	return i
}


func main() {
	fmt.Println("Hello, World!")
	initDb()
	CountQuestions()
	ShowAllQuestions()
	memoryDb.Close()
	//readSqlite()
}
