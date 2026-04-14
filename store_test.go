package main

import (
	"bytes"
	"fmt"
	"testing"
)


func TestPathTransformFunc(t *testing.T){
	key := "momsbestpicture"
	pathname := CASPathTransformFunc(key)
	fmt.Println(pathname)
	expected_path := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathname != expected_path{
		t.Errorf("have %s want %s",pathname,expected_path)
	}
}

func TestStore(t *testing.T){
	opts := StoreOpts{
		TransformFunc: CASPathTransformFunc,

	}
	s := NewStore(opts)
	data := bytes.NewReader([]byte("some jpeg"))
	if err := s.writestream("myspecialpicture",data); err != nil{
		t.Error(err)
	}
}
