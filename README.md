# SRICore

**SRICore** es una librería para generar, validar y consultar comprobantes electrónicos y datos del contribuyente en el SRI de Ecuador.

## 📥 Instalación

Puedes instalar **SRICore** directamente desde el repositorio usando Go Modules:

```bash
go get github.com/pinzlab/sricore
```

## 📦 id

Este paquete proporciona funciones para validar números de identificación ecuatorianos, incluyendo la Cédula (DNI) y diferentes tipos de RUC (Registro Único de Contribuyentes).

Soporta validaciones para:
- Cédula (DNI)
- RUC para Personas Naturales
- RUC para Empresas Privadas
- RUC para Entidades Públicas

La validación se realiza utilizando los algoritmos Modulo 10 y Modulo 11 según lo especificado por la normativa ecuatoriana.





### Validar una Cédula (DNI)

```go
	if err := id.IsDNI("0601234560"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cédula valida")
```

### Validar RUC

```go
	if err := id.IsRUC("0601234560001"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("RUC valido")
```


## 📦 ws

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
exists, err := service.CheckRUC("0601234560001")
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
contributors, err := service.GetContributors("0601234560001")
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
establishments, err := service.GetEstablishments("0601234560001")
if err != nil {
	log.Fatal(err)
}

for _, establishment := range establishments {
	fmt.Printf("Establecimiento: %s\n", establishment.Number)
	fmt.Printf("Dirección: %s\n", establishment.Address)
	fmt.Printf("Estado: %s\n", establishment.Status)
}

```