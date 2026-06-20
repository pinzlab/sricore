# ¿Cómo validar cédula y RUC en Ecuador?

En este documento encontrarás una guía clara para validar correctamente los números de **cédula** y **RUC** en Ecuador. Se detalla la estructura del RUC, los distintos tipos de contribuyentes y los algoritmos que se utilizan para verificar su validez: el **Módulo 10**, aplicado a cédulas y RUCs de personas naturales, y el **Módulo 11**, usado para validar RUCs de entidades públicas, jurídicas y extranjeras.

> ⚠️ **Fuente:** Este contenido se basa en el artículo original de Bryan Suárez publicado en [Medium](https://medium.com/@bryansuarez/c%C3%B3mo-validar-c%C3%A9dula-y-ruc-en-ecuador-b62c5666186f), donde se explica detalladamente el proceso de validación de cédula y RUC en Ecuador.


## Validación de la cédula

El número de cédula ecuatoriano contiene **10 dígitos**, donde:

1. Los dos primeros dígitos indican la provincia de emisión (de 01 a 24).
2. El tercer dígito debe estar en el rango de 0 a 5.
3. Del cuarto al noveno dígito forman un número secuencial.
4. El décimo dígito es el dígito verificador, calculado con el algoritmo *Módulo 10*.


```mermaid
stateDiagram

    ruc: 06 - 0 - 123456 - 0

    ruc --> provinceCode
    provinceCode: 06
    provinceCode --> province
    province: Código de provincia

    ruc --> contributorCode
    contributorCode: 0
    contributorCode --> contributor
    contributor: En rango de [0 a 5]

    ruc --> serialCode
    serialCode: 123456
    serialCode --> serial
    serial: Número consecutivo

    ruc --> checkerNumber
    checkerNumber: 0
    checkerNumber --> checker
    checker: Dígito verificador
    
```


### Algorítmo `Módulo 10`

Este algoritmo se aplica para validar cédulas y RUCs de personas naturales. Consiste en multiplicar los primeros nueve dígitos por un patrón alternado de coeficientes `2` y `1`, sumando los resultados y verificando que el décimo dígito (verificador) sea correcto.

Pasos:
1. Multiplicar los dígitos impares por 2 y los pares por 1.
2. Si algún producto es mayor que 9, restar 9.
3. Sumar todos los resultados.
4. Restar esa suma del siguiente múltiplo de 10.
5. El resultado debe coincidir con el último dígito de la cédula.


```mermaid
block-beta
columns 1
  
    block:MUL
        columns 11

        name1["Cédula"] ba1<[" "]>(right) 
        dni1["0"] dni2["6"] dni3["0"] dni4["1"] dni5["2"] dni6["3"] dni7["4"] dni8["5"] dni9["6"]

        space:2       
        mop1["*"] mop2["*"] mop3["*"] mop4["*"] mop5["*"] mop6["*"] mop7["*"] mop8["*"] mop9["*"]

        name2["Coeficiente"] ba2<["siempre"]>(right) 
        coe1["2"] coe2["1"] coe3["2"] coe4["1"] coe5["2"] coe6["1"] coe7["2"] coe8["1"] coe9["2"]

    end
    space

    block:DIFF
        columns 11

        name3["Producto"] ba3<[" "]>(right)
        min1["0"] min2["6"] min3["0"] min4["1"] min5["4"] min6["3"] min7["8"] min8["5"] min9["12"]

        space:2       
        sop1["-"] sop2["-"] sop3["-"] sop4["-"] sop5["-"] sop6["-"] sop7["-"] sop8["-"] sop9["-"]

        name4["Diferencia"] ba4<[" "]>(right)
        sub1["0"] sub2["0"] sub3["0"] sub4["0"] sub5["0"] sub6["0"] sub7["0"] sub8["0"] sub9["9"]
    
    end
    space
    space 
    
    block:RES
        columns 11

        name5["Resultado"] ba5<[" "]>(right)
        rst1["0"] rst2["6"] rst3["0"] rst4["1"] rst5["4"] rst6["3"] rst7["8"] rst8["5"] rst9["3"] 

        space:11

        name7["Sumatoria"] ba7<[" "]>(right)
        space dv0["30"] ba6<[" "]>(right) dv1["40 - 30 = 10"] ba8<[" "]>(right) dv2["0"] ba9<[" "]>(right) dni["0601234560"]
    
    end


    MUL -- "El producto se resta 9 si es mayor a 10" --> DIFF
    DIFF -- "Sumar todos los dígitos resultantes. \n Restar el resultado de la decena superior. \n Si el resultado es 10 el digito validador será 0 " --> RES
  
  

    classDef alpha stroke-width:0px
    class name1,name2,name3,name4 alpha
    class mop1,mop2,mop3,mop4,mop5,mop6,mop7,mop8,mop9 alpha
    class sop1,sop2,sop3,sop4,sop5,sop6,sop7,sop8,sop9 alpha

```

## Validación del RUC

La validación del RUC depende del tipo de contribuyente. Los métodos varían si el RUC pertenece a una persona natural, una entidad pública o una persona jurídica.


### Persona natural

Para personas naturales, los primeros 10 dígitos del RUC deben ser una cédula válida (validada con el algoritmo Módulo 10). El RUC termina en un código de establecimiento de tres dígitos (`001` usualmente).

Características:
- El tercer dígito está entre `0` y `5`.
- El décimo dígito se valida con *Módulo 10*.
- Los tres últimos dígitos (`001`) indican el establecimiento principal.


```mermaid
stateDiagram

    ruc: 06 - 0 - 123456 - 0 - 001

    ruc --> provinceCode
    provinceCode: 06
    provinceCode --> province
    province: Código de provincia

    ruc --> contributorCode
    contributorCode: 0
    contributorCode --> contributor
    contributor: En rango de [0 a 5]

    ruc --> serialCode
    serialCode: 123456
    serialCode --> serial
    serial: Número consecutivo

    ruc --> checkerNumber
    checkerNumber: 0
    checkerNumber --> checker
    checker: Dígito verificador

    ruc --> estabCode
    estabCode: 001
    estabCode --> estab
    estab: Número de establecimiento
    
```


### Públicos

Los RUCs de entidades del sector público tienen el **tercer dígito igual a 6**, y su validación se realiza con el algoritmo *Módulo 11*.

Características:
- El dígito verificador (noveno) se valida con el algoritmo Módulo 11 usando los coeficientes `[3, 2, 7, 6, 5, 4, 3, 2]`.
- El dígito resultante se resta de 11 para obtener el verificador. Si el resultado es 11, se usa 0; si es 10, no es válido.
- El código de establecimiento debe ser `0001`.


```mermaid
stateDiagram

    ruc: 17 - 6 - 00015 - 5 - 0001

    ruc --> provinceCode
    provinceCode: 17
    provinceCode --> province
    province: Código de provincia

    ruc --> contributorCode
    contributorCode: 6
    contributorCode --> contributor
    contributor: Siempre 6

    ruc --> serialCode
    serialCode: 00015
    serialCode --> serial
    serial: Número consecutivo

    ruc --> checkerNumber
    checkerNumber: 5
    checkerNumber --> checker
    checker: Dígito verificador

    ruc --> estabCode
    estabCode: 0001
    estabCode --> estab
    estab: Número de establecimiento
    
```


### Jurídicos y extranjeros sin cédula

Este tipo de RUC corresponde a sociedades privadas o personas extranjeras sin cédula ecuatoriana. El **tercer dígito es 9**, y también se valida con el algoritmo *Módulo 11*.

Características:
- El dígito verificador (décimo) se valida con coeficientes `[4, 3, 2, 7, 6, 5, 4, 3, 2]`.
- La suma se resta de 11 para obtener el verificador. Si el resultado es 11, se usa 0; si es 10, no es válido.
- El código de establecimiento también es `001`.


```mermaid
stateDiagram

    ruc: 17 - 9 - 001167 - 4 - 001

    ruc --> provinceCode
    provinceCode: 17
    provinceCode --> province
    province: Código de provincia

    ruc --> contributorCode
    contributorCode: 9
    contributorCode --> contributor
    contributor: Siempre 9

    ruc --> serialCode
    serialCode: 001167
    serialCode --> serial
    serial: Número consecutivo

    ruc --> checkerNumber
    checkerNumber: 4
    checkerNumber --> checker
    checker: Dígito verificador

    ruc --> estabCode
    estabCode: 001
    estabCode --> estab
    estab: Número de establecimiento
    
```