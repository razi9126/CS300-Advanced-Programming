const fs = require('fs')
const http = require('http')
const socketio = require('socket.io')

const readFile = f => new Promise((resolve,reject) =>
	fs.readFile(f, (e, d) => e? reject(e):resolve(d)))


const server = http.createServer( async (req,resp) =>
	resp.end( await readFile(req.url.substr(1))))

const io = socketio(server)
io.sockets.on('connection',socket=>
	socket.on('mymsg',data => io.sockets.emit('yourmsg',data)))

server.listen(8000, () => console.log('Started'))
