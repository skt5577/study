package composite

import (
	"container/list"
)

type file interface {
	Path() string
	Size() float32
}

type concreteFile struct {
	path string
	size float32
}
type directory struct {
	path  string
	files list.List
}

func (file *concreteFile) Path() string {
	return file.path
}
func (file *concreteFile) Size() float32 {
	return file.size
}

func (directory *directory) Path() string {
	return directory.path
}

func (directory *directory) Size() float32 {
	var totalSize float32
	for e := directory.files.Front(); e != nil; e = e.Next() {
		child := (e.Value).(file)
		totalSize += child.Size()
	}
	return totalSize
}

func (directory *directory) addFile(file file) {
	directory.files.PushBack(file)
}

func (directory *directory) deleteFile(file file) {
	var element *list.Element
	for e := directory.files.Front(); e != nil; e = e.Next() {
		if e.Value == file {
			element = e
			break
		}
	}
	directory.files.Remove(element)
}
