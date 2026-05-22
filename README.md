# truth-table

CLI em Go para gerar tabelas-verdade a partir de expressões lógicas proposicionais simples.

## Visão geral

O programa recebe uma expressão pela linha de comando, faz o parse da fórmula e imprime a tabela-verdade no terminal.

Atualmente o projeto está em um estado inicial e funciona melhor para expressões curtas, com proposições de uma única letra.

## Requisitos

- Go `1.26.1`

## Como executar

Com `go run`:

```bash
go run . 'p^q'
```

Gerando um binário:

```bash
go build -o truth-table .
./truth-table 'p^q'
```

## Sintaxe suportada

- `~` para negação
- `^` para conjunção
- `|` para disjunção
- `(` e `)` para agrupamento

### Exemplos de expressões

```text
p^q
~p|q
(p|q)
```

## Exemplo de saída

Entrada:

```bash
go run . 'p^q'
```

Saída:

```text
p | q | p^q
 true | true | true |
 true | false | false |
 false | false | false |
 false | true | false |
```

## Estrutura do projeto

- `main.go`: ponto de entrada da CLI e impressão da tabela
- `parser/lexer.go`: tokenização da expressão
- `parser/parser.go`: construção da AST
- `parser/ast.go`: avaliação da expressão lógica
- `parser/props.go`: extração das proposições usadas na fórmula
- `queue/queue.go`: fila genérica usada em rotinas auxiliares

## Limitações atuais

- A expressão deve ser passada como um único argumento.
- Não use espaços na fórmula. Exemplo válido: `p^q`. Exemplo inválido no estado atual: `p ^ q`.
- As proposições devem ser letras únicas, como `p`, `q` e `r`.
- Casos mais complexos com parênteses encadeados e fórmulas com 3 ou mais proposições ainda podem produzir resultados incorretos.

## Próximos passos sugeridos

- suportar identificadores com mais de um caractere
- corrigir a geração das combinações da tabela para mais proposições
- melhorar a validação de expressões inválidas
