package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

//Page sirve para la estrucura de las páginas
type Page struct{
    Title string
    Body []byte
}

func (p *Page) save() error {
    filename := "./data/" + p.Title + ".txt"
    //(nombre,contenido,permisos)
    err := ioutil.WriteFile(filename,p.Body,0600)
    return err
}


func loadPage(title string) *Page {
    filename := "./data/" + title + ".txt"
	body,_ := ioutil.ReadFile(filename)
	page := &Page{Title:title,Body: body}

	return page
}

func main(){
	// page := &Page{Title:"primer",Body:[]byte("Primer articulo con go y archivos yeah!!!")}
	page := loadPage("primer")
	page.save()
	fmt.Println(page.Title,string(page.Body))
}
