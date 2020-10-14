package main

import "fmt"

//ReadCloser is readable, closeable interface
type ReadCloser interface {
	Read(n int) []byte
	Close()
}

// Buffer is memory buffer
type Buffer struct {
	Mem      []byte
	Length   int
	CurIndex int
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (buf *Buffer) Read(n int) []byte {
	if buf.CurIndex >= buf.Length {
		return []byte{}
	}
	end := min(buf.CurIndex+n, buf.Length)
	ret := buf.Mem[buf.CurIndex:end]
	buf.CurIndex += min(buf.CurIndex+n, buf.Length)
	return ret
}

//Close will clean the buffer
func (buf *Buffer) Close() {
	buf.CurIndex = buf.Length
}

type Shape interface {
	GetArea() int
}

func AddArea(shape1 Shape, shape2 Shape) int {
	return shape1.GetArea() + shape2.GetArea()
}

type Circle struct {
	Radius int
}

func (c *Circle) GetArea() int {
	return 3 * c.Radius * c.Radius
}

type Square struct {
	Side int
}

func (s *Square) GetArea() int {
	return s.Side * s.Side
}

func main() {
	var reader ReadCloser
	buf := Buffer{Mem: make([]byte, 16), Length: 16, CurIndex: 0}
	reader = &buf
	fmt.Printf("%v.\n", reader.Read(8))

	var c Circle
	c.Radius = 1
	var s Square
	s.Side = 2
	fmt.Printf("%v", AddArea(&c, &s))
}
