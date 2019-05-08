package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	// fmt.Println(t)
	t.SkipNow()
	t.Errorf("Wrong....")
}

func TestPrint1(t *testing.T) {
	// fmt.Println("test1")
	t.Run("hello", func(t *testing.T) {
		fmt.Println("sub test1")
	})
	t.Run("hi", func(t *testing.T) {
		fmt.Println("sub test2")
	})
}

func xTestPrint1() {
	fmt.Println("jashfgkjhasgkjh")
}

func TestMain(m *testing.M) {
	fmt.Println("First....")

	m.Run()
}

func BenchmarkAll(b *testing.B) {
	for n := 0; n > b.N; n++ {
		xTestPrint1()
	}
}
