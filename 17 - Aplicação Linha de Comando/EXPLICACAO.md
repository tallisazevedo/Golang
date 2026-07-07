# Explicação do Projeto — Aplicação de Linha de Comando

Este documento explica, **linha por linha**, como o projeto funciona — seguindo a
ordem real em que o código é executado (e não a ordem dos arquivos na pasta).

---

## 1. Visão geral do fluxo

```
Terminal (você digita)
        │
        ▼
   main.go  →  chama app.Gerar()  →  monta a estrutura da CLI
        │
        ▼
   aplicacao.Run(os.Args)  →  a lib urfave/cli interpreta o que foi digitado
        │
        ▼
   identifica o comando (ip | servidores) e a flag (--host)
        │
        ▼
   chama a função Action correspondente (buscarIps | buscaServidores)
        │
        ▼
   usa o pacote padrão "net" para consultar o DNS
        │
        ▼
   imprime o resultado no terminal
```

Guarde esse diagrama — vamos percorrer cada uma dessas caixas em detalhe.

---

## 2. `go.mod` — a "certidão de nascimento" do módulo

```go
module linha-de-comando

go 1.26.4

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli v1.22.17 // indirect
)
```

| Linha | Explicação |
|---|---|
| `module linha-de-comando` | Nome do módulo. É o "endereço" que outros arquivos usam para importar pacotes internos do projeto (veja `"linha-de-comando/app"` no `main.go`). |
| `go 1.26.4` | Versão mínima do compilador Go exigida. |
| `require (...)` | Lista de dependências externas que o projeto baixa (via `go get`) e trava em uma versão específica, para builds reproduzíveis. |
| `// indirect` | Indica que você **não importa** esse pacote diretamente no seu código — ele é usado internamente pelo `urfave/cli` (para gerar texto de ajuda em Markdown, por exemplo). O Go rastreia isso automaticamente. |

> Note que `urfave/cli` está marcado como `indirect`, mesmo sendo importado diretamente em `app.go`. Isso normalmente é ajustado rodando `go mod tidy`, mas não afeta o funcionamento do programa.

---

## 3. `main.go` — o ponto de entrada

```go
package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partidas")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
```

### Linha a linha

**`package main`**
Todo executável Go precisa de um pacote chamado `main`. É esse pacote que o comando `go build` transforma em um binário. Se fosse uma biblioteca (para outros importarem), o nome seria outro (ex: `package app`).

**`import (...)`**
Bloco de importação. Cada linha é um pacote usado neste arquivo:
- `"fmt"` → pacote padrão para formatação/impressão de texto (`Println`, `Printf`, etc).
- `"linha-de-comando/app"` → **pacote interno** do próprio projeto. O caminho é `<nome-do-módulo>/<pasta>`, por isso bate com `module linha-de-comando` do `go.mod` + a pasta `app/`.
- `"log"` → pacote padrão para logging; aqui é usado só para `log.Fatal`.
- `"os"` → pacote padrão de interação com o sistema operacional; aqui é usado só para `os.Args`.

**`func main() { ... }`**
A função `main`, dentro do pacote `main`, é o ponto de partida obrigatório de qualquer programa Go — é a primeira função executada quando você roda o binário. Não recebe parâmetros e não retorna nada.

**`fmt.Println("Ponto de partidas")`**
Apenas imprime uma linha de texto no terminal, para sinalizar que o programa começou. Não tem relação com a lógica da CLI — é só um log manual de debug/aprendizado.

**`aplicacao := app.Gerar()`**
- `app.Gerar()` chama a função `Gerar` exportada pelo pacote `app` (veja seção 4). Ela devolve um `*cli.App` (um **ponteiro** para uma struct `cli.App`, já configurada com nome, comandos e flags).
- `:=` é o operador de **declaração curta**: cria a variável `aplicacao` e já infere o tipo (`*cli.App`) a partir do valor retornado. Equivale a `var aplicacao *cli.App = app.Gerar()`, mas mais idiomático em Go.

**`if erro := aplicacao.Run(os.Args); erro != nil { log.Fatal(erro) }`**
Este é um dos padrões mais comuns em Go, então vale destrinchar:

- `os.Args` é um `[]string` (slice de strings) contendo todos os argumentos passados no terminal, incluindo o nome do próprio binário na posição `0`. Ex: se você roda `./app ip --host google.com`, então `os.Args == ["./app", "ip", "--host", "google.com"]`.
- `aplicacao.Run(os.Args)` é o método do `urfave/cli` que **parseia** esse slice, descobre qual comando foi chamado e executa a `Action` correspondente. Ele retorna um `error` (que é `nil` se tudo correu bem).
- `if erro := X; erro != nil { ... }` é a sintaxe de **if com inicialização**: você declara `erro` *dentro* do próprio `if`, e ele só existe dentro do escopo do bloco. É a forma idiomática em Go de checar erros sem "poluir" o resto da função com uma variável `erro` que só serviria para essa checagem.
- `log.Fatal(erro)` imprime o erro no terminal e encerra o programa imediatamente com código de saída `1` (equivalente a `fmt.Println(erro); os.Exit(1)`).

> **Por que Go não tem `try/catch`?** Porque o idioma da linguagem é tratar erros como **valores retornados normalmente**, checados explicitamente com `if err != nil`. É verboso, mas deixa o fluxo de erros visível e explícito em vez de "escondido" em exceções.

---

## 4. `app/app.go` — onde a CLI é montada

```go
package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)
```

- `package app` → este arquivo pertence ao pacote `app`, por isso é importado em `main.go` como `"linha-de-comando/app"` e usado como `app.Gerar()`.
- `"net"` → pacote padrão para operações de rede (aqui, consultas DNS).
- `"github.com/urfave/cli"` → a biblioteca externa que faz o trabalho pesado de parsear argumentos de linha de comando.

> Repare que o Go organiza importações em dois "grupos" separados por linha em branco: primeiro os pacotes padrão (`fmt`, `log`, `net`), depois os pacotes externos (`urfave/cli`). Isso é só uma convenção de formatação (o `gofmt`/`goimports` costuma fazer isso automaticamente).

### 4.1 A função `Gerar`

```go
// Gerar() vai retornar uma aplicação de linha de comando pronta para ser executada.
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes de servidores na internet"
```

- O comentário logo acima da função, começando com o nome da própria função (`Gerar() vai...`), é a convenção de **doc comment** do Go — ferramentas como `go doc` e o hover do editor exibem esse texto como documentação da função.
- `func Gerar() *cli.App` → função exportada (começa com letra maiúscula, então é visível fora do pacote `app`) que **retorna um ponteiro** para `cli.App`.
- `cli.NewApp()` → função de fábrica da biblioteca que cria e retorna um `*cli.App` com valores padrão.
- `app.Name` / `app.Usage` → atribuição direta de campos da struct. Como `app` é um ponteiro, o Go desreferencia automaticamente (você não precisa escrever `(*app).Name`).

> Cuidado com o nome da variável: aqui `app` é uma variável **local** (`*cli.App`), diferente do **pacote** `app` que dá nome ao arquivo. Dentro desta função, `app` se refere à variável local — não há conflito porque o pacote não precisa ser referenciado por nome dentro de si mesmo.

### 4.2 As flags

```go
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
```

- `[]cli.Flag{...}` → um **slice literal**: cria uma lista de valores do tipo `cli.Flag` (que na verdade é uma *interface* — `cli.StringFlag` é uma das structs que a implementam).
- `cli.StringFlag{ Name: "host", Value: "devbook.com.br" }` → **struct literal com campos nomeados**. Em vez de passar valores na ordem dos campos (`cli.StringFlag{"host", "devbook.com.br"}`), você nomeia cada campo (`Name:`, `Value:`), o que é mais legível e mais seguro contra mudanças de ordem na struct.
- Resultado prático: isso cria a flag `--host`, que pode ser passada no terminal (`--host google.com`); se **não** for passada, assume o valor padrão `"devbook.com.br"`.
- Essa variável `flags` é criada **uma vez** e reaproveitada nos dois comandos abaixo — evita repetir a mesma definição duas vezes.

### 4.3 Os comandos

```go
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscaServidores,
		},
	}

	return app
}
```

- `app.Commands` recebe um slice de `cli.Command`. Cada item é um **subcomando** que o usuário pode chamar no terminal: `./app ip ...` ou `./app servidores ...`.
- Repare que dentro do slice, cada elemento omite o tipo (`cli.Command{...}` vira só `{...}`) — o Go infere o tipo a partir do tipo do slice (`[]cli.Command`). É um atalho permitido para literais de slice/array.
- `Name` → o texto que o usuário digita para escolher o comando.
- `Usage` → texto de ajuda exibido em `./app --help`.
- `Flags: flags` → reaproveita o slice de flags definido na seção anterior.
- `Action: buscarIps` → aqui está o ponto-chave: **está passando a função `buscarIps` como valor**, não chamando ela (repare que não tem `()`). O `urfave/cli` guarda essa referência e só vai *chamar* `buscarIps(c)` no momento em que o usuário digitar `./app ip`. Isso é possível porque, em Go, funções são **valores de primeira classe** — podem ser atribuídas a variáveis, passadas como argumento, guardadas em structs, etc.
- `return app` → devolve o ponteiro `*cli.App` totalmente configurado para quem chamou `Gerar()` (no caso, o `main.go`).

---

## 5. As ações — onde a lógica de rede acontece

### 5.1 `buscarIps`

```go
func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
```

- `func buscarIps(c *cli.Context)` → esta é a assinatura que o `urfave/cli` espera para uma `Action` (v1 aceita `func(*cli.Context)` ou `func(*cli.Context) error`; aqui foi usada a variante sem retorno).
- `c *cli.Context` → objeto que representa "o que o usuário digitou nesta chamada": valores de flags, argumentos posicionais, etc.
- `c.String("host")` → busca o valor da flag `--host` como string. Se o usuário não passou `--host`, retorna o `Value` padrão definido em `cli.StringFlag` (`"devbook.com.br"`).
- `ips, erro := net.LookupIP(host)` → chama a função da **biblioteca padrão** `net` que faz a resolução DNS do host (busca os endereços IP associados). Retorna dois valores: o slice de IPs encontrados (`[]net.IP`) e um possível erro — padrão extremamente comum em Go (`valor, erro := algumaFuncao()`).
- `if erro != nil { log.Fatal(erro) }` → mesmo padrão de tratamento de erro já visto no `main.go`.
- `for _, ip := range ips { fmt.Println(ip) }` → laço `for...range`, que percorre o slice `ips`. O `range` sempre retorna dois valores (índice e valor); como o índice não é usado aqui, é descartado com `_` (identificador "branco", que diz ao Go "eu sei que existe esse valor, mas não vou usá-lo"). `ip` é cada endereço IP individual, impresso um por linha.

### 5.2 `buscaServidores`

```go
func buscaServidores(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range ips {
		fmt.Println(server.Host)
	}
}
```

Estrutura idêntica à anterior, com duas diferenças:

- `net.LookupNS(host)` → em vez de buscar IPs, busca os registros **NS** (*Name Server*) do domínio — ou seja, quais servidores DNS são responsáveis por resolver aquele domínio. Retorna `[]*net.NS`.
- `server.Host` → cada `server` é um ponteiro para uma struct `net.NS`, que tem o campo `Host` (o nome do servidor). Por isso o `Println` acessa `server.Host` em vez de imprimir `server` diretamente.

> Nota: a variável se chama `ips` nesta função também, embora agora contenha servidores de nome, não IPs — é só uma reutilização de nome pelo instrutor, não afeta o funcionamento.

---

## 6. Juntando tudo: um exemplo de execução

Digamos que você rode no terminal:

```
go run main.go ip --host google.com
```

Passo a passo:

1. `os.Args` vira `["main.go", "ip", "--host", "google.com"]` (aproximadamente).
2. `app.Gerar()` já retornou o `*cli.App` configurado, com os comandos `ip` e `servidores` registrados.
3. `aplicacao.Run(os.Args)` lê os argumentos, identifica o subcomando `ip` e a flag `--host` com valor `"google.com"`.
4. O `urfave/cli` monta um `*cli.Context` com esse valor e chama `buscarIps(c)`.
5. Dentro de `buscarIps`, `c.String("host")` retorna `"google.com"`.
6. `net.LookupIP("google.com")` faz a consulta DNS de verdade e retorna os IPs.
7. Cada IP é impresso no terminal, um por linha.

Se você rodasse só `go run main.go ip` (sem `--host`), o passo 5 retornaria `"devbook.com.br"` — o valor padrão definido na flag.

---

## 7. Resumo dos conceitos de Go vistos neste projeto

| Conceito | Onde aparece | Ideia central |
|---|---|---|
| Pacotes e módulos | `package main`, `package app`, `go.mod` | Organização de código e dependências |
| Ponteiros (`*T`) | `*cli.App`, `*cli.Context` | Passar/retornar referência em vez de cópia |
| Slices (`[]T`) | `os.Args`, `[]cli.Flag`, `[]cli.Command` | Listas dinâmicas |
| Structs literais | `cli.StringFlag{Name: ..., Value: ...}` | Criar valores de structs nomeando os campos |
| Funções como valores | `Action: buscarIps` | Passar uma função sem chamá-la, para ser executada depois |
| Múltiplos retornos | `ips, erro := net.LookupIP(host)` | Padrão idiomático de retorno de erro em Go |
| `if` com inicialização | `if erro := ...; erro != nil` | Escopo local para variável de erro |
| `for...range` | `for _, ip := range ips` | Iteração sobre slices, descartando valores com `_` |
| Tratamento de erro | `log.Fatal(erro)` | Sem exceções — erro é valor, checado explicitamente |

---

## 8. Sobre a biblioteca `urfave/cli`

Resumo do papel dela neste projeto (sem ela, você teria que fazer manualmente com o pacote padrão `flag`):

- Registrar comandos e subcomandos (`ip`, `servidores`).
- Registrar flags reutilizáveis entre comandos (`--host`).
- Gerar automaticamente texto de `--help`.
- Parsear `os.Args` e rotear para a `Action` certa.
- Fornecer o `*cli.Context` como forma padronizada de ler flags dentro da `Action`.

As dependências `go-md2man` e `blackfriday` no `go.mod` são usadas **internamente pelo urfave/cli** para gerar documentação em Markdown/man pages — você nunca as importa ou chama diretamente.
