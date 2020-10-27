package main

import(
	"fmt"
	"./archivos_mul"
)

func main()  {
	var dato1, dato2 string
	var dato3 uint
	op := 0

	contenido := archivos_mul.ContenidoWeb{
		Mul: []archivos_mul.Multimedia{},
	}
	
	for i := 0; i < 1; i++ {
		fmt.Println("\nMenu\n1.- Capturar imagen\n2.- Capturar audio\n3.- Capturar video\n4.- Mostrar elementos guardados\n0.- Salir")
		fmt.Scan(&op)
		if op == 1{
			fmt.Println("\nImagen")
			fmt.Print("Titulo: ")
			fmt.Scan(&dato1)
			fmt.Print("Formato: ")
			fmt.Scan(&dato2)
			fmt.Print("Canales: ")
			fmt.Scan(&dato3)
			archivo := archivos_mul.Imagen{Titulo: dato1, Formato: dato2, Canales: dato3}
			contenido.Mul = append(contenido.Mul,&archivo)
			i--
		}
		if op == 2{
			fmt.Println("\nAudio")
			fmt.Print("Titulo: ")
			fmt.Scan(&dato1)
			fmt.Print("Formato: ")
			fmt.Scan(&dato2)
			fmt.Print("Duracion en segundos: ")
			fmt.Scan(&dato3)
			archivo := archivos_mul.Audio{Titulo: dato1, Formato: dato2, Duracion: dato3}
			contenido.Mul = append(contenido.Mul,&archivo)
			i--
		}
		if op == 3{
			fmt.Println("\nVideo")
			fmt.Print("Titulo: ")
			fmt.Scan(&dato1)
			fmt.Print("Formato: ")
			fmt.Scan(&dato2)
			fmt.Print("Frames: ")
			fmt.Scan(&dato3)
			archivo := archivos_mul.Video{Titulo: dato1, Formato: dato2, Frames: dato3}
			contenido.Mul = append(contenido.Mul,&archivo)
			i--
		}
		if op == 4{
			contenido.Mostrar()
			i--
		}
	}
	fmt.Println("\nHasta pronto\n")
}