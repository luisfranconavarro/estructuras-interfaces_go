package archivos_mul

import "fmt"

type Multimedia interface{
	Mostrar()
}

type ContenidoWeb struct{
	Mul []Multimedia
}

func (cWeb *ContenidoWeb) Mostrar(){
	fmt.Println()
	for _, datos := range cWeb.Mul {
		datos.Mostrar()
	}
}

type Imagen struct{
	Titulo string
	Formato string
	Canales uint
}

func (i *Imagen) Mostrar(){
	fmt.Println("titulo:",i.Titulo,"	formato:",i.Formato,"	canales:",i.Canales)
}

type Audio struct{
	Titulo string
	Formato string
	Duracion uint
}

func (a *Audio) Mostrar(){
	fmt.Println("titulo:",a.Titulo,"	formato:",a.Formato,"	duracion:",a.Duracion)
}

type Video struct{
	Titulo string
	Formato string
	Frames uint
}

func (v *Video) Mostrar(){
	fmt.Println("titulo:",v.Titulo,"	formato:",v.Formato,"	frames:",v.Frames)
}