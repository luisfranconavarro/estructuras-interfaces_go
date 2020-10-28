package main

import(
	"fmt"
	"bufio"
	"os"
	"./archivos_mul"
)

func main()  {
	var dato1, dato2 string
	var dato3 uint
	op := 0

	scanner := bufio.NewScanner(os.Stdin)

	contenido := archivos_mul.ContenidoWeb{
		Mul: []archivos_mul.Multimedia{},
	}

	for true {
		fmt.Println("\nMenu\n1.- Capturar imagen\n2.- Capturar audio\n3.- Capturar video\n4.- Mostrar elementos guardados\n0.- Salir")
		fmt.Scanln(&op)
		if op == 1{
			fmt.Println("\nImagen")
			fmt.Print("Titulo: ")
			scanner.Scan()
			dato1 = scanner.Text()
			fmt.Print("Formato: ")
			fmt.Scanln(&dato2)
			fmt.Print("Canales: ")
			fmt.Scanln(&dato3)
			archivo := archivos_mul.Imagen{Titulo: dato1, Formato: dato2, Canales: dato3}
			contenido.Mul = append(contenido.Mul,&archivo)
		}else if op == 2{
			fmt.Println("\nAudio")
			fmt.Print("Titulo: ")
			scanner.Scan()
			dato1 = scanner.Text()
			fmt.Print("Formato: ")
			fmt.Scanln(&dato2)
			fmt.Print("Duracion en segundos: ")
			fmt.Scanln(&dato3)
			archivo := archivos_mul.Audio{Titulo: dato1, Formato: dato2, Duracion: dato3}
			contenido.Mul = append(contenido.Mul,&archivo)
		}else if op == 3{
			fmt.Println("\nVideo")
			fmt.Print("Titulo: ")
			scanner.Scan()
			dato1 = scanner.Text()
			fmt.Print("Formato: ")
			fmt.Scanln(&dato2)
			fmt.Print("Frames: ")
			fmt.Scanln(&dato3)
			archivo := archivos_mul.Video{Titulo: dato1, Formato: dato2, Frames: dato3}
			contenido.Mul = append(contenido.Mul,&archivo)
		}else if op == 4{
			contenido.Mostrar()
		}else if op == 0{
			break
		}else{
			fmt.Println("Opcion incorrecta")
		}
	}
	fmt.Println("\nHasta pronto\n")
}