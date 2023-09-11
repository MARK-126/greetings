# Saludos en golang
Este paquete proporciona un ejemplo de como se pueden obtener saludos aleatorios en Golang

## Instalacion
Ejecuta el siguiente comando para instalar el paquete de manera local
```bash
go get -u github.com/MARK-126/greetings

```
## USo
Aquí tienes un ejemplo de como puedes usar el paquete en tu código:

```go
func main() {

message, err := greetings.Hello("Marcos") 
	if err == nil {
		fmt.Println("Ocurrio un error: ", err)
    return
	}
	fmt.Print(message)
} 
```
Este ejemplo importa el paquete github.com/MARK-126/greetings y llama a la funcion Hello que nos devuelve un saludo personalizado. 
Si ocurre un error, se imprime un mensaje de error.
