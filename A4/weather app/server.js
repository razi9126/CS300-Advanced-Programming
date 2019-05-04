const fs = require('fs')
const http = require('http')

const readFile = f => new Promise((resolve,reject) =>
	fs.readFile(f, (e, d) => e? reject(e):resolve(d)))


const server = http.createServer( async (req,resp) =>
	resp.end( await readFile(req.url.substr(1))))

server.listen(8000, () => console.log('Someone connected'))
