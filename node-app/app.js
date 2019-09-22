const express = require('express')
const { readFile, writeFile } = require('fs')
const path = require('path')
const crypto = require('crypto')
const { promisify } = require('util')
const axios = require('axios')

const app = express()
const readFileProm = promisify(readFile)
const writeFileProm = promisify(writeFile)

const filepath = path.resolve(`${__dirname}/arquivo-node.txt`)

const handler = async (_req, res) => {
  try {
    // Lê os dados de um arquivo específico, 
    // concatena aos dados uma string randômica
    // e salva o arquivo com os novos dados
    const data = await readFileProm(filepath, 'utf-8')
    const content = `${data}\n${generateRandomString(256)}`
    await writeFileProm(filepath, content)
    
    // Usa os dados randômicos gerados anteriormente
    // para criar um novo arquivo numa pasta específica
    // O nome do arquivo é gerado randômicamente
    // await generateRandomFile(content)

    // Chama uma api http com dados mock,
    // escreve no stdout as informações retornadas
    // await logMockApiResponse()

    // Retorna os dados no response body
    return res.status(200).send(content)

  } catch (err) {
    console.log(err)
    return res.status(404).end()
  }
}

const generateRandomString = (bytes) => crypto.randomBytes(bytes).toString('hex')

const generateRandomFile = (content) => writeFileProm(`./contents/${generateRandomString(10)}.txt`, content)

const logMockApiResponse = async () => {
  const response = await axios({ 
    method: 'get', 
    url: 'http://5d879522cd71160014aaeac7.mockapi.io/api/v1/users'
  })

  response.data
    .forEach(user => console.log(`\nid: ${user.id}\nNome: ${user.name}\n`))

  return true
}

app.get('/', handler)
// Start server
const port = process.env.APP_PORT || 8080
app.listen(port, () => console.log(`Server listening in ${port}`))