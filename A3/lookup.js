// ReadME
// M. Razi ul Haq  2020-10-0015
// Phewwwwwwww. Done with this finally. Sir's sneak peak gave away the whole idea :) :) :)

// What works:
// Basically everything apart from the fact that it wont search for a list of words. It takes 
// a SINGLE WORD ONLY
// Valid commands example:
		// node lookup.js index E:/tests/hello.json E:/tests
		// node lookup.js search hello.json razi

// What doesn't works:
// 	Doesnt work for array of words. It was simply some if's and a promiseall.map but
// 	i was too tired by now to implement it



const fs = require('fs')
const path = require('path')

const json = {};

const insert_json = (dir) => new Promise((resolve,reject) =>  //To add words to json
{
	return readFile(dir).then(filedata => {
	var array = filedata.toString().split('\r\n');
	var count=0
	for(i in array) {
		count=count+1
        var array_word=array[i].split(' ');
        var count2=0
        for(j in array_word)
        {
        	count2=count2+1
            // console.log(array_word[j])   //prints each word
            if (array_word[j].length>3)
            {
            	if (json.hasOwnProperty(array_word[j])) 
                 //check if word already present and then check if same file name.
                 // If true then append only line

                 {
                    if (json[array_word[j]][(json[array_word[j]].length)-1]['file']===dir)    //if file already present, then append line number only
                    {
                    	json[[array_word[j]]][(json[array_word[j]].length)-1]['line'].push(count)
                    }
                    else //insert a new jason object for that word
                    {
                    	var data = {
                    		"file": dir,
                    		"line": [count]
                    	};
                    	json[array_word[j]].push(data)}}
                else //make a new word and its jason object
                {
                	json[array_word[j]]=[]
                	var data2 = {
                		"file": dir,
                		"line": [count]
                	};
                	json[array_word[j]].push(data2)
                }}}
        if ((count+1)===(array.length))
        { 
        	resolve()}
    }
})}
)


const readdir = dir => new Promise((resolve,reject) =>
	fs.readdir(dir,(err,file) => err? reject(err) : resolve(file))
	)

const isDirectory = path => new Promise((resolve,reject) =>
	fs.lstat(path,(err,stat) => err? reject(err) : resolve(stat.isDirectory()))
	)

const readFile = file => new Promise((resolve,reject) =>
	fs.readFile(file,'utf-8',(err,data) => err? reject(err) : resolve(data))
	)

const writeFile = (file,data) => new Promise((resolve,reject) =>
	fs.writeFile(file,data,'utf-8',err=> err? reject(err) : resolve())
	)


cbx=[]

//Gets all the txt files first. puts them in the array cbx
function crawl(addr, cb) {
	return readdir(addr).then(path_list => {
		return Promise.all(path_list.map(filepath =>{
			filepath = path.join(addr, filepath);
			return isDirectory(filepath).then(stats => {
				if (stats===true) {
					return crawl(filepath,cbx);
				}
				else 
				{
					cbx.push(filepath);
				}
			});
		}));
	});
}

// function lookup
var lines=[]
const printline= arrdata => new Promise((resolve,reject)=>
{
	return Promise.all(arrdata.map(abc =>
	{
		return readFile(abc['file']).then(filedata => {
			var array = filedata.toString().split('\r\n');
			for(j in abc['line'])
			{
				var count=0
				for(i in array) {
					count=count+1
					if (count===abc['line'][j])
						{	
							 console.log(array[i])
							lines.push(array[i])
							resolve()
						}
				}
			}
		})
	}))
})


function lookup(fname,words){
	var obj;
	return readFile(fname).then(data =>{
		obj=JSON.parse(data)
		// console.log(obj)
		console.log(words)
		if (obj.hasOwnProperty(words))
		{
			var q=words
			var mydata= (obj[q])
			printline(mydata).then(()=>
			{
				return
				// console.log("Printing done")
			// for (h in lines)
			// 	console.log(lines[h])
		})
		}
		else 
		{
			console.log("Word not found")
		}
	})
}

if (process.argv[2]=='index')
{
	let destfile=process.argv[3]
	let pat = process.argv[4]
	// console.log(path)
	// console.log(destfile)
	pat=path.normalize(pat)
	crawl(pat,cbx).then(() =>
	{
		return Promise.all(cbx.map(abc => 
		{
			if (path.extname(abc) =='.txt')  
				{return insert_json(abc)}}
			)
		).then(()=>
		{
			var final=JSON.stringify(json)
			return writeFile(destfile, final).then(()=>
			{
				console.log("Indexed!")
			})
		}
		)})

}
else if (process.argv[2]=='search')
{
	let jsonfile=process.argv[3]
	let word = process.argv[4]
	lookup(jsonfile,word).then(()=>
	{
		console.log('Done reading')
	})
}

//TEST code
// crawl('E:/tests',cbx).then(() =>
// {
// 	return Promise.all(cbx.map(abc => 
// 	{
// 		if (path.extname(abc) =='.txt')  
// 		{return insert_json(abc)}}
// 	)
// 	).then(()=>
// 	{
// 		var final=JSON.stringify(json)
// 		return writeFile("E:/tests/index.json", final).then(()=>
// 		{
// 			console.log("Indexed!")
// 		})
// 	}
// )})


