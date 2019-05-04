const socket = io()
const state = {}
const buttons = []

//socket.emit('start')
const formSubmit = event => {
	event.preventDefault()
	// WRITE HERE THAT EMIT ID SHIT STUFF
}
myBoard = []
a=0
w = ""
// bttns = []
// const topButtons = () => {
// 	for (let i=0; i<=6;i++)
// 	{
// 		bttns[i] = React.createElement('input', {type: 'submit', value: i, id: 'btn' + i})

// 	}
// }


const setState = updates => {
	Object.assign(state, updates)
	ReactDOM.render(
		React.createElement('form',{onSubmit: formSubmit},
			React.createElement('h1',null,'CONNECT4'),
			React.createElement('p',null,'Instructions:-'),
			React.createElement('div',null,'1. Press the number to insert in its respective column.'),
			React.createElement('div',null,'2. Your aim is to connect 4 of your characters vertically, horizontally or diagonally. First to do so wins'),
			React.createElement('div',null,'3. Have fun.'),
			React.createElement('p',null,''),
			React.createElement('div',null, w),
			React.createElement('input',{type: 'submit',value: 1, id: 1}),
			React.createElement('input',{type: 'submit',value: 2, id: 2}),
			React.createElement('input',{type: 'submit',value: 3, id: 3}),
			React.createElement('input',{type: 'submit',value: 4, id: 4}),
			React.createElement('input',{type: 'submit',value: 5, id: 5}),
			React.createElement('input',{type: 'submit',value: 6, id: 6}),
			React.createElement('input',{type: 'submit',value: 7, id: 7}),
			React.createElement('div', null, myBoard),
			),
		 document.getElementById('root')
		)
}
const create_grid = () => {
	for(i=0;i<6;i++)
	{
		col = []
		for(j=0;j<7;j++)
		{
			col[j] = React.createElement('input',{type: 'submit',value: "  ", id: a, disabled: true})
			a++
		}
		myBoard[i] = React.createElement('div',null,col)
	}	
}
socket.on('wait', data => 
	{
		console.log(data)
		ReactDOM.render(React.createElement('p', null, data), document.getElementById('root'))
	})
socket.on('left', () => 
{
	ReactDOM.render(React.createElement('p', null, "Other player left. Refresh again."), document.getElementById('root'))
})

socket.on('start', () => {
	//topButtons()
	create_grid()
	setState()
})
//setState()