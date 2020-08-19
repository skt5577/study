package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	directory := directory{path: "directoryPath"}
	var file1 file = &concreteFile{path: "path1", size: 20.5}
	var file2 file = &concreteFile{path: "path2", size: 10.425}
	var file3 file = &concreteFile{path: "path3", size: 1.721}
	var file4 file = &concreteFile{path: "path4", size: 0.5}

	directory.addFile(file1)
	directory.addFile(file2)
	directory.addFile(file3)
	directory.addFile(file4)

	assert.Equal(t, 4, directory.files.Len())
	assert.Equal(t, file1.Size()+file2.Size()+file3.Size()+file4.Size(), directory.Size())

	directory.deleteFile(file3)
	directory.deleteFile(file4)

	assert.Equal(t, 2, directory.files.Len())
	assert.Equal(t, file1.Size()+file2.Size(), directory.Size())
}
