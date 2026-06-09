package main

import "fmt"

type FileSystemEntity interface{
	GetSize() int 
	ShowDetails()
	GetName() string
}

type FileEntity struct{
	Name string
	Size int
}

func (f *FileEntity)GetSize() int {
	return f.Size
}

func (f *FileEntity)GetName() string {
	return "File : "+f.Name
}

func (f *FileEntity)ShowDetails(){
	fmt.Printf("File Name : %s , Size : %d bytes\n",f.Name,f.Size)
}

func NewFile(Name string,Size int) *FileEntity{
	return &FileEntity{
		Name: Name,
		Size: Size,
	}
}


type DirectoryEntiity struct{
	Name string
	Children []FileSystemEntity
}


func (d *DirectoryEntiity)GetName() string {
	return "Directory : " +d.Name
}


func (d *DirectoryEntiity)GetSize() int {
	totalSize:=0
	for _,child := range d.Children {
		totalSize+=child.GetSize()
	}
	return totalSize
}

func (d *DirectoryEntiity)ShowDetails(){
	for _,Child := range d.Children {
		fmt.Printf("Directory Name : %s , Size : %d bytes\n",Child.GetName(),Child.GetSize())
		Child.ShowDetails()
	}
}

func NewDirectory(Name string) *DirectoryEntiity{
	return &DirectoryEntiity{
		Name: Name,
		Children: []FileSystemEntity{},
	}
}

func (d *DirectoryEntiity) Add(Entity FileSystemEntity){
	d.Children = append(d.Children, Entity)
}

type Iterator interface{
	HasNext() bool
	Next() FileSystemEntity
}


type DFSIterator struct{
	stack []FileSystemEntity
}

func NewDFSIterator(root FileSystemEntity) *DFSIterator{
	return &DFSIterator{
		stack: []FileSystemEntity{root},
	}
}

func (it *DFSIterator) HasNext() bool{
	return len(it.stack) > 0
}

func (it *DFSIterator) Next() FileSystemEntity{
	if !it.HasNext(){
		return nil
	}
	n:=len(it.stack)
	current:= it.stack[n-1]
	it.stack = it.stack[:n-1]

	if dir,ok := current.(*DirectoryEntiity); ok {
		
		for i:=len(dir.Children)-1;i>=0;i--{
			it.stack = append(it.stack, dir.Children[i])
		}

	}
	return current


}



func main() {

	resume := NewFile("Resume.pdf",43)
	notes := NewFile("Notes.txt",94)
	goa := NewFile("Goa.jpg",100)
	family := NewFile("Family.jpg",233)
	details:=NewFile("Personal.docs",823)

	root := NewDirectory("Root")
	documents := NewDirectory("Documents")
	pictures := NewDirectory("Pictures")
	vacation := NewDirectory("Vacation")
	personal := NewDirectory("Personal")

	personal.Add(details)

	documents.Add(resume)
	documents.Add(notes)
	documents.Add(personal)

	vacation.Add(goa)

	pictures.Add(vacation)
	pictures.Add(family)

	root.Add(documents)
	root.Add(pictures)



	iterator := NewDFSIterator(root)

	for iterator.HasNext() {
		entity := iterator.Next()
		fmt.Println(entity.GetName())
	}
}