const express = require('express')
const { readFile, writeFile } = require('fs')
const path = require('path')
const crypto = require('crypto')
const { promisify } = require('util')

const app = express()
const readFileProm = promisify(readFile)
const writeFileProm = promisify(writeFile)

const filepath = path.resolve(`${__dirname}/arquivo-node.txt`)

app.get('/', async (_req, res) => {
  try {
    const data = await readFileProm(filepath, 'utf-8')

    const content = `${data}\n${generateRandomString()}`

    await writeFileProm(filepath, content)
    
    // await generateRandomFile(content)

    return res.status(200).send(content)

  } catch (err) {
    console.log(err)
    return res.status(404).end()
  }
})

const generateRandomFile = (content) => writeFileProm(`./contents/${generateRandomString()}.txt`, content)
const generateRandomString = () => crypto.randomBytes(20).toString('hex')


const port = process.env.APP_PORT || 8080
app.listen(port, () => console.log(`Server listening in ${port}`))