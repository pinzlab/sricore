# SRICore

**SRICore** es una librería para generar, validar y consultar comprobantes electrónicos y datos del contribuyente en el SRI de Ecuador.

## 📦 Instalación

Puedes instalar **SRICore** directamente desde el repositorio usando Go Modules:

```bash
go get github.com/pinzlab/sricore
```


## Cliente SRIOnline

**SRICore** expone un cliente HTTP llamado SRIOnline para interactuar con los endpoints públicos del SRI (Servicio de Rentas Internas). Este cliente permite consultar:

- La existencia de un RUC
- Información general del contribuyente
- Establecimientos registrados asociados al RUC

### Cómo instanciar el cliente

Antes de llamar a cualquiera de las funciones, debes crear una instancia del cliente SRIOnline:

```go
package main

import "github.com/pinzlab/sricore/ws"

func main() {
	service := ws.NewSRIOnline()

	// Usar las funciones con el cliente:
	// - service.CheckRUC(...)
	// - service.GetContributors(...)
	// - service.GetEstablishments(...)

}

```
### Verificar si un RUC existe

```go
exists, err := service.CheckRUC("9999999999001")
if err != nil {
	log.Fatal(err)
}

if exists {
	fmt.Println("El RUC existe en el SRI.")
} else {
	fmt.Println("El RUC no existe.")
}
```

### Obtener información del contribuyente

```go
contributors, err := service.GetContributors("9999999999001")
if err != nil {
	log.Fatal(err)
}

for _, contributor := range contributors {
	fmt.Printf("Razón social: %s\n", contributor.BusinessName)
	fmt.Printf("Estado: %s\n", contributor.Status)
	fmt.Printf("Actividad económica: %s\n", contributor.EconomicActivity)
}

```

### Obtener establecimientos registrados

```go
establishments, err := service.GetEstablishments("9999999999001")
if err != nil {
	log.Fatal(err)
}

for _, establishment := range establishments {
	fmt.Printf("Establecimiento: %s\n", establishment.Number)
	fmt.Printf("Dirección: %s\n", establishment.Address)
	fmt.Printf("Estado: %s\n", establishment.Status)
}

```