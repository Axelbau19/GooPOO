package main

/*CMD utilizados
go get  repositorioName // para obtener una depedencia o codigo de otra persona en github
go mod init nombre // nombrar tu modulo
go build //para construir proyecto //
go mod tiny // agrega depedencias al proyecto
go clean -modcache // limpia el cache */
import (
	"github.com/donvito/hellomod"
	helloMod2 "github.com/donvito/hellomod/v2"
)

func main() {
	hellomod.SayHello()
	helloMod2.SayHello("Axel")
}
