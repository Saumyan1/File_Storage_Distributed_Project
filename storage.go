package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)


func CASPathTransformFunc(key string) string{
	hash := sha1.Sum([]byte(key)) //[20]byte => []byte => [:]
	hashStr := hex.EncodeToString(hash[:])
	blocksize := 5
	sliceLen := len(hashStr) / blocksize
	paths := make([]string, sliceLen)
	for i := 0; i<sliceLen;i++{
		from ,to := i * blocksize,(i*blocksize)+blocksize
		paths[i] = hashStr[from:to]

	}
	return strings.Join(paths, "/")
	

}
//When you put a function inside a struct, you aren't running the code yet. 
//You are storing the logic to be executed later. This is incredibly 
// powerful for making your code flexible.

type PathTransformFunc func(string) string

type StoreOpts struct{
	TransformFunc PathTransformFunc


}
var DefaultPathTransormFunc = func(key string) string{
	return key
}


type Store struct{
	StoreOpts

}

func NewStore(opts StoreOpts) *Store{
	return &Store{
		StoreOpts: opts,
	}
}

func(s *Store) writestream(key string, r io.Reader) error{
	pathName := s.TransformFunc(key)

	if err := os.MkdirAll(pathName, os.ModePerm); err != nil{
		return err
	}

	buf := new(bytes.Buffer)
	io.Copy(buf,r)

	filenameBytes := md5.Sum(buf.Bytes())
	

	filename := hex.EncodeToString(filenameBytes[:])
	pathAndFilename := pathName + "/" +filename
	f,err := os.Create(pathAndFilename)
	if err != nil{
		return err
	}
	n,err := io.Copy(f,buf)
	if err != nil{
		return err
	}
	log.Printf("written (%d) bytes to disk: %s", n,pathAndFilename)
	return nil
}