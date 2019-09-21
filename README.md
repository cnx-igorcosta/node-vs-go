# node-vs-go

Um benchmark simples para avaliar a diferença de performance entre um server criado com NodeJs com o framework Express e um server em Go.

## Detalhes
As duas aplicações fazem exatamente a mesma coisa: Levantam um servidor, expôem uma rota GET padrão "/" e a cada vez que bate nessa rota, eles abrem um arquivo numa pasta, lêem o conteúdo do arquivo, escrevem uma string randômica concatenando no conteúdo previamente existnte no arquivo e retornam essa string para o response.

Em cada um deles tem um código comentado que serve para criar, além da tarefa normal, um arquivo novo numa pasta com o conteúdo lido do arquivo anterior. Pode ser descomentado para comparar o desempenho das duas aplicações executando mais tarefas "pesadas".

Esses testes servem para simular rotinas de I/O que normalmente têm um custo mais pesado para as aplicações, pode ser testado também fazendo requisições http para outro server, conectando a banco de dados, etc...

## Ferramenta para benchmark
Utilizei a ferramenta (Siege)[https://www.euperia.com/wrote/speed-testing-your-website-with-siege-part-one/] que estressa a aplicação e gera um relatório ao fim do teste com várias informações relevantes.

### Comandos:
siege -c10 -r1 -d10 -v http://localhost:8080

- c10 é o numero de usuários concorrentes a gente quer simular.
- r1 é o número de repetições cada usuário fará.
- d10 é o delay entre cada request de usuário (cada simulação os usuários entram em stand-by por um intervalo entre 0 e 10 segundos). 
- -v é para mostrar as saídas de cada request.
- Por último, a url a ser estressada

## Dependências
NodeJs e npm

## Start dos servidores
### Node server:
```
  cd ./node-app
  npm start
```
** url: http://localhost:8080 **

### Go server
```
  cd ./go-app
  go build ./main.go
  ./go-app
```
** url: http://localhost:8081 **

Obs: Ideal testar cada server isoladamente para não ter nenhum tipo de interferência.