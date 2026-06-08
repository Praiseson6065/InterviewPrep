package main

import "fmt"


type FileSystemEntity interface{
	GetName() string
	GetSize() int
	ShowDetails()
}


type FileEntity struct{
	name string
	size int
}

func (file *FileEntity) GetName() string{
	return file.name
}

func (file *FileEntity) GetSize() int{
	return file.size
}

func (file *FileEntity) ShowDetails() {
	
	fmt.Printf("%s ,FileName : %s , Size : %d\n",file.name,file.name,file.size)
}

func NewFile(name string,size int) *FileEntity{
	return &FileEntity{
		name: name,
		size:size,
	}
}

type DirectoryEntity struct{
	name string
	children []FileSystemEntity
}

func NewDirectory(name string) *DirectoryEntity{
	return &DirectoryEntity{
		name: name,
	}
}

func (directory *DirectoryEntity) Add(entity FileSystemEntity) {

	directory.children = append(directory.children, entity)
}

func (directory *DirectoryEntity) GetName() string{
	return directory.name
}

func (directory *DirectoryEntity) ShowDetails() {

	fmt.Printf("%s/ ,Directory : %s ,Size : %d\n",directory.name,directory.name,directory.GetSize())
	for _,child := range directory.children {
		fmt.Printf("%s/",directory.name)
		child.ShowDetails()

	}
}

func (directory *DirectoryEntity) GetSize() int{
	
	
	totalSize :=0

	var Child FileSystemEntity

	for _,Child = range directory.children {

		totalSize+=Child.GetSize()

	}
	return totalSize
}


func main(){
	resume:= NewFile("Resume.pdf",54)
	notes := NewFile("Notes.txt", 50)
	photo := NewFile("Photo.jpg", 200)


	documents:=NewDirectory("Documents")
	documents.Add(resume)
	documents.Add(notes)

	root:=NewDirectory("root")
	root.Add(documents)
	root.Add(photo)

	root.ShowDetails()
	fmt.Printf("Total Size: %d KB\n", root.GetSize())

}





