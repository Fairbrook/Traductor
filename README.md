# Proyecto de la clase de Seminario de traductores de lenjuage

## Cómo correr el proyecto

Para el sistema operativo Windows 64bits puede ejecutar el archivo [.exe](https://github.com/Fairbrook/Traductor/releases/tag/v0.1-alpha.7) que se encuentra en los release del repositorio, si está en otro sistema operativo o no funciona el ejecutable deberá compilarlo para la plataforma

### Compilador

Si bien el programa puede realizar la traducción por si solo, para el último paso de la compilación es necesario tener instalado en el sistema el compilador [MASM32](http://masm32.com/download.htm).
Y su binario deberá estar en la variable `%PATH%` del sistema

⚠ **Es necesario ejecutar la aplicación dentro del mismo disco en el que MASM32 fue instalado**

## Cómo utilizar el proyecto

Mientras ingresa el código en el panel izquierdo, la tabla de símbolos aparecerá a la derecha. En caso de existir un error, aparecerá el mensaje en la sección de la derecha, en vez de la tabla de símbolos.
![pantalla](https://i.ibb.co/9b26n4L/image.png)

Una vez se esté conforme con el código escrito, el programa cuenta con 3 formas distintas de visualizar la traducción.
En el menú _File_ está la opción _Traducir a ensamblador_ el cual realizará la traducción pertiente y abrirá el código resultante en el programa que se tenga registrado como por defecto

![pantalla](https://i.ibb.co/QfJyjhd/image.png)

En el menú _Compilacion_ se encuentran las otras dos opciones, _Compilar_ y _Compilar y ejecutar_

![pantalla](https://i.ibb.co/z2yDL26/image.png)

Ambos realizan una compilación del código, la diferencia radica en que, como su nombre lo indica, la segunda opción ejecuta el código compilado y lo muestra en la barra lateral derecha, como se muestra en la imagen
![pantalla](https://i.ibb.co/pyHfPkf/image.png)

Cabe mencionar que todos los subproductos de la compilación son generados dentro de la carpeta que se encuentra ejecutandose la aplicación

![pantalla](https://i.ibb.co/LgZym3L/image.png)

## Introducción

El propóstio de ese proyecto es la creación de un traductor entre un subset del lenguaje de programación C y ensamblador. Para realizar esta traducción y analizar la entrada del usuario se utilizan diferentes etapas descritas a continuación

## Analizador léxico

Tipos admitidos

| Símbolo             | Descirpción                         |
| ------------------- | ----------------------------------- |
| Identificador       | letra(letra\|digito)\*              |
| Entero              | digito+                             |
| Real                | entero.entero                       |
| Adicion             | + -                                 |
| Multiplicacion      | \* /                                |
| Asignacion          | =                                   |
| Relacional          | < > <= >= != ==                     |
| And                 | &&                                  |
| Or                  | \|\|                                |
| Not                 | !                                   |
| Parentesis          | (,)                                 |
| Llave               | {,}                                 |
| Punto y coma        | ;                                   |
| Palabras reservadas | if, while, return, else, int, float |

Esta priemra entrega realiza el analizador léxico mediante una pseudo tabla de estados, se le llama seudo porque engloba los casos de múltiples caracteres en funciones como _isDigit_ y _isAlpha_ para determinar si un carcter es numérico o una letra

Esta tabla de estados es utilizada por un autómata que analiza la cadena de entrada hasta llegar a un caracter para el cual no se tiene un estado siguiente, una vez detenido regresa al estado 0 y reinicia el proceso

## Analizador Sintáctico

El analizador sintáctico permite identificar secuencia de lexemas válidos para el lenguaje determinado
Por lo general se dfine con una serie de reglas que permiten a un autómata, a partir de la salida de Analizador Léxico verificar la secuencia de los tokens

En específico, para esta primera entrega, se utilizaron las siguientes reglas

- R1 -> Aceptación
- R2 -> id + id | E
- R3 -> id

Dando como resultado la siguiente tabla de estados

|     | id  | +   | $   | E   |
| --- | --- | --- | --- | --- |
| 0   | d2  |     |     | 1   |
| 1   |     |     | R1  |     |
| 2   |     | d3  | R2  |     |
| 3   | d2  |     |     | 4   |
| 4   |     |     | R1  |     |

Tomando en cuenta que en el programa la repesentación numérica de los caracteres es

- id = 0
- \+ = 5
- $ = 23
- E = 103

Si bien la salida de este módulo es un símple válido o inválido, se decidió mostrar todo el proceso que sigue el analizador en la barra derecha con el fin de ejemplificar el progreso

Para la segunda entrega, la pila está compuesta por objetos en vez de cadenas, estos objetos son árboles y cuando se realiza una reducción, se guarda en el árbol correspondiente resultante los segmentos usados para su construcción

## Modulo 4 - Gramática del compilador

En esta entrega se hace uso de una serie de reglas extendidas para el compilador, lo cual permite identificar cada segmento del programa gracias al analizador sintáctico que identifica cada regla

## Modulo 5 - Árbol sintáctico

La creación de un árbol sintáctico consiste en almacenar los componentes de cada reducción en un árbol permitiendo la posterior creación de una tabla de símbolos

Para que quede más claro se utilizó el sigueinte código de ejemplo

```
int a;
int suma(int a, int b){
    return a+b;
}

int main(){
    float a;
    int b;
    int c;
    c = a+b;
    c = suma(8,9);
}
```

__El cual da como resultado [este árbol sintáctico](https://laughing-ardinghelli-16d744.netlify.app/)__

## Modulo 6 - Tabla de símbolos

Mediante el árbol obtenido del analizador sintáctico, podemos generar una tabla de símbolos, esta tabla permit realizar las validaciones necesarias para el contexto en cada parte del programa. Por ejemplo, si tenemos el siguiente código

```
int suma(int a, int b){
     a+b;
}

int main(){
    int a, b;
    a = 3;
    b = 10;
    int c;
    c = suma(a,b);
}
```

generará esta tabla de símbolos:

![tabla](https://i.ibb.co/n8ztCnT/Screenshot-2.png)

Este análisis se logra mediante el corrido del arbol sintáctico y la utilizacion de estructuras "map" o diccionarios

Cada vez que se encuentra una deficion en el programa se añade un registro a la tabla de símbolos, indicando si es una función, una variable y su tipo de datos. En caso de ser una función se guardan los parámetros que require.

Otra observación es que se están utilizando anidación de tablas, lo que significa que cuando el programa entra en un bloque nuevo (funcion) genera una subtabla que se vincula al regitro de la funcion en la tabla principal

Cuando el programa encuentra una expresión detecta que tipo es y en caso de ser un identificador utiliza la tabla del contexto actual para encontrar la defición original. En caso de no existir se envía el error correspondiente

### Funciones Print

La forma más sencilla de saber si un programa corre correctamente es su salida estándard, por lo tanto era necesario que el programa fuera capaz de _"Imprimir"_. Para este objetivo, se inyectaron en la tabla de símbolos tres funciones que permiten realizar la función de tipos sin necesidad de cambiar la lógica del analizador semántico. Estas tres funciones son:

| Funcion | Prototipo      | Descripción                                                    |
| ------- | -------------- | -------------------------------------------------------------- |
| printS  | printS(char\*) | Recibe una cadena en formato de C y la muestra en pantalla     |
| printF  | printF(float)  | Recibe una expresión de coma flotante y la muestra en pantalla |
| printI  | printI(int)    | Recibe una expresión entera y la muestra en pantalla           |

Estas tre funciones están reslpaldadas por la función printf de una de las librerías estándard de __MASM32__, lo que nos permite delegarle la responsabilidad del manejo de la salida estándard, uno de los puntos por lo que se escogió este compilador

## Módulo 7 - Traductor

El último módulo y producto final de la aplicación es el traductor a ensamblador, este módulo como se puede imaginar, es el encargado de tomar la tabla de símbolos junto con un _Arbol decorado_ del analizador semántico y convertirlo en código ensamblador para su compilación en **MASM32**

Sería demasiado extenso explicar toda la lógica que sigue este paso, por lo que un ejemplo es más adecuado. Si se ingresa el sigueinte código, por ejemplo:

```
int suma(int a, int b){
	return a+b;
}

int main(){
	int a;
	a=suma(1,3);
	printI(a);
	printS("\nhola mundo!");
	return 0;
}
```

Se obtendría el siguiente código ensamblador:

```
.386
.model flat, stdcall
option casemap:none

INCLUDE \masm32\include\masm32rt.inc

.data
_feax real8 0.0
_fhelper dword 0

.code

;---------- Funcion suma -----------
suma proc suma_a:DWORD,suma_b:DWORD
;locals
;fin locals

mov eax, suma_b
mov ebx, eax
mov eax, suma_a
add eax, ebx
ret
suma endp

;---------- Funcion main -----------
main proc
;locals
local main_a:DWORD
;fin locals

finit
; Llamada a la función suma
mov eax, 3
push eax
mov eax, 1
push eax
call suma
mov main_a, eax
mov eax, main_a
printf("%u",eax)
printf("\nhola mundo!")
mov eax, 0
; Retorno al SO
invoke ExitProcess, eax
main endp

end main
```

## Tecnologías utilizadas

| Tecnologia                                                  | Utilizacion                                                                                                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| [golang](https://golang.org/)                               | Funcionalidad principal (backend)                                                                                                               |
| [electron](https://www.electronjs.org/)                     | Desarrollo de aplicaciones nativas con ui desarrollado en html, css y js                                                                        |
| [asilectron](https://github.com/asticode/astilectron)       | Proporciona sockets TCP para la comunicación entre Electron y cualquier lenguaje de programación con el fin de escuchar los eventos de electron |
| [go-asilectron](https://github.com/asticode/go-astilectron) | Api de asilectron desarrollada para go                                                                                                          |
| [quill](https://quilljs.com/)                               | Como editor de texto para el texto de entrada                                                                                                   |
| [AlpineJs](https://alpinejs.dev/)                           | Librería de JS para facilitar el comportamiento reactivo del UI                                                                                 |
| [MASM32](http://masm32.com/)                                | Entorno de desarrollo para la compilación de ensamblador en Windows                                                                             |
