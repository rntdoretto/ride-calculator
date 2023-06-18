# Code Smells

## 00 - Código original

* criar testes unitários com 100% de cobertura para iniciar as refatorações

## 01 - Simple Refactoring

Nomes

* renomear variável
* renomear função ou método
* renomear classe ou arquivo

Magic number

* introduzir constantes ou variáveis explicativas

Linha em branco dentro de função ou método

* apagar linha em branco

Código morto

* apagar código comentado

## 02 - Structs Refactoring

Longa lista de parâmetros ou parâmetros abertos

* introduzindo struct parâmetro

## 03 - Methods Refactoring

Comentários

* criando variáveis explicativas
* extraindo métodos explicativos

## 04 - Cognitive Complexity Refactoring

* introdução de cláusula guarda
* consolidar comandos condicionais

## 05 - Handle Errors Refactoring

* utilizar um tratamento de erros adequado

## 06 - Object Oriented Programming

* introduzindo objetos

## 07 - Design Pattern Strategy

Design Patterns
    Uma forma conhecida e comum de distribuição de responsabilidades entre classes

Strategy: Ter comportamentos diferentes para o mesmo contrato

## 08 - Design Pattern Strategy + Factory

Design Patterns
    Uma forma conhecida e comum de distribuição de responsabilidades entre classes

Factory (simple): Criar instâncias com base em algum critério

## 09 - Design Pattern Chain Of Responsability

Design Patterns
    Uma forma conhecida e comum de distribuição de responsabilidades entre classes

Chain of Responsability: Encadeamento de comportamento, permitindo que um dos elos (classes encadeadas) DECIDAM se tratam ou não, e se não tratarem, passam para o próximo

## SOLID

S - Single Responsability
    Uma classe/arquivo deve ter um, e somente um, motivo para mudar
O - Open / Closed
    Objetos ou entidades devem estar abertos para extensão, mas fechados para modificaçã
L - Liskov Substitution
    Uma classe derivada deve ser substituível por sua classe base
I - Interface Segregation
    Uma classe não deve ser forçada a implementar interfaces e métodos que não irão utilizar
D - Dependency Inversion
    Dependa de abstrações e não de implementações
