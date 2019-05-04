

const fs = require('fs')
const path = require('path')

const writeFile = (file, data) => new Promise((resolve, reject) =>
	fs.appendFile(file, data,'utf8', err => err? reject(err): resolve()))

const copyall_promise_write = (srcFiles, dest) =>
	Promise.all(srcFiles.map(file => new Promise(resolve =>
		fs.readFile(file, (err, data) => {
			if (err) {
				console.log('Reading ${file} ${err}')
				resolve()
			} else {
				console.log(file + ' read, writing it now')
				writeFile(path.join(dest), data).then( data => {
					console.log('Writinggg' + file + err)
					resolve()
				}).catch(err => {
					console.log(err)
					resolve()
				})
			}
		}))))
copyall_promise_write(['a.txt','b.txt','c.txt'], 'test.txt').then( ()=>
	console.log('ALL doneee'))

//