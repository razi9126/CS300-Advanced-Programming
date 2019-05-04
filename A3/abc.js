
 // const json = {};
// var data = {
//     "file":'/root/erfe.txt',
//     "line": [5]
// };
// var data2 = {
//     "file":'a/c/d/razi.txt',
//     "line": [3]
// };
// json['computer']=[data]
// json['computer'].push(data2)
// console.log(json['computer'].length)
// console.log(json['computer'][(json['computer'].length)-1]['file']==='a/c/d/razi.txt')
// console.log(json['computer'])



const fs = require('fs')
const path = require('path')
// const dir = 'E:/Razi/Docs/Lums Notes/Soph2/CS300/A3/a.txt'
const json = {};
const insert_json = (dir) => new Promise((resolve,reject) =>
 {
    var array = fs.readFileSync(dir).toString().split("\n");
    var count=0
    for(i in array) {
        count=count+1
        // console.log(array[i]);   //prints the line
        var array_word=array[i].split(' ');
            var count2=0
            // console.log('============')
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

                        json[array_word[j]].push(data)
                    }

                }
                else //make a new word and its jason object
                {
                    json[array_word[j]]=[]
                    var data2 = {
                        "file": dir,
                        "line": [count]
                    };
                    json[array_word[j]].push(data2)

                }
            }

        }
        if ((count+1)===(array.length))
        {   
            console.log('Parsed completely')
            // resolve()
        }
    }
}
)

const crawl= (dir) => new Promise((resolve,reject) =>
{ 
     arr =fs.readdirSync(dir) 
        
            console.log("Folder " + path.basename(dir) + " found with files = " + arr.length)

            resolve(arr)
        }

        )

// crawl('E:/Razi/Docs/Lums Notes/Soph2/CS300/A3/tests').then((fromCrawl) =>
//     {console.log(fromCrawl)}).catch(err=>{
//     console.log('=======')
//         console.log(err)})

// insert_json('a.txt').then( ()=>
// {
//     console.log("COdv")
// }).catch(() => {
//      console.log('BLAH');
//  });

// const traverse= (dir) => 
// Promise.all( crawl(dir).then((fromCrawl)=>
// {
//     console.log(fromCrawl)
//     fromCrawl.map(file => new Promise((resolve, reject) =>
//     {   
//         let next = path.join(dir,file);
//             console.log(file)
//             if (fs.lstatSync(next).isDirectory()===true)
//             {
//                 traverse(next).then((abc)=>
//                 {
//                     console.log(`Traversed folder: ${next} `)
//                     resolve()
//                 }).catch(err => 
//                 {
//                     console.log(`Error from ${next}`)
//                     console.log(err)
//                     reject()
//                     // resolve()
//                 });

//             }
//             else
//             {
//                 if (path.extname(next) ==='.txt')
//                 {

//                     insert_json(next).then( ()=>
//                     {
//                         console.log(`parsed ${next}`)
//                     }).catch(() => {
//                         console.log(`COULDNT parse ${next}`);
//                         reject()
//                     });
//                 }
//                 resolve()
//             }
           

//         }).catch(err => {
//             console.log('=======')
//             console.log(err)}))
// }).catch(error => {
//     console.log('crawl error')
//     console.log("Fuck JS");
// })
//     )




// s =  traverse('E:/Razi/Docs/Lums Notes/Soph2/CS300/A3/tests',()=>
//  {
//      console.log('HAHAAHAHA')  
//  });


s =  traverse('E:/Razi/Docs/Lums Notes/Soph2/CS300/A3/tests').then(()=>
 {
     console.log('HAHAAHAHA')
 }).catch(() => {
     console.log('BLAH');
    
 });

