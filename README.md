# Proyecto de la clase de Seminario de traductores de lenjuage

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

## Haz click [aquí](https://laughing-ardinghelli-16d744.netlify.app/) para ver el árbol de forma interactiva

El cual da como resultado el siguiente árbol sintáctico
![arbol](https://i.ibb.co/fS4nHHT/Screenshot-2021-10-08-at-13-39-49-Screenshot.png)

## Cómo correr el proyecto

Para el sistema operativo Windows 64bits puede ejecutar el archivo [.exe](https://github.com/Fairbrook/Traductor/releases/tag/v0.1-alpha.5) que se encuentra en los release del repositorio, si está en otro sistema operativo o no funciona el ejecutable deberá compilarlo para la plataforma
Requerira un entorno de [golang](https://golang.org/)
instalar el paquete [go-asilectron](https://github.com/asticode/go-astilectron) y su respectivo [bundler](https://github.com/asticode/go-astilectron-bundler)

## Cómo utilizar el proyecto

Escriba el texto de entrada en el area de la izquierda y espera a que aparezca el resultado en la parte derecha
![pantalla](https://i.ibb.co/rpdG5R4/Screenshot-5.png)

## Tecnologías utilizadas

| Tecnologia                                                  | Utilizacion                                                                                                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| [golang](https://golang.org/)                               | Funcionalidad principal (backend)                                                                                                               |
| [electron](https://www.electronjs.org/)                     | Desarrollo de aplicaciones nativas con ui desarrollado en html, css y js                                                                        |
| [asilectron](https://github.com/asticode/astilectron)       | Proporciona sockets TCP para la comunicación entre Electron y cualquier lenguaje de programación con el fin de escuchar los eventos de electron |
| [go-asilectron](https://github.com/asticode/go-astilectron) | Api de asilectron desarrollada para go                                                                                                          |
| [quill](https://quilljs.com/)                               | Como editor de texto para el texto de entrada                                                                                                   |
| [AlpineJs](https://alpinejs.dev/)                           | Librería de JS para facilitar el comportamiento reactivo del UI                                                                                 |
