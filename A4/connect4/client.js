const socket = io()
const state = {}
const board = [];
var symbol='-'

var mouseOverHandler = function mouseOverHandler() {
		console.log('you moused over');
	};
var clickhandler = function clickhandler() {
		console.log('you clicked');
	};

var handleClick1 = function handleClick1() {
	let row = 5
	while(row>-1)
	{
		if (board[row][0]!=='- ')
			row-=1
		else
		{
			board[row][0]=symbol
			break;
		}
	}
	sendtoClient()
	makeBoard()
	setState()
}
var handleClick2 = function handleClick2() {
	let row = 5
	while(row>-1)
	{
		if (board[row][1]!=='- ')
			row-=1
		else
		{
			board[row][1]=symbol
			break;
		}
	}
	sendtoClient()
	makeBoard()
	setState()
}
var handleClick3 = function handleClick3() {
	let row = 5
	while(row>-1)
	{
		if (board[row][2]!=='- ')
			row-=1
		else
		{
			board[row][2]=symbol
			break;
		}
	}
	sendtoClient()
	makeBoard()
	setState()
}
var handleClick4 = function handleClick4() {
	let row = 5
	while(row>-1)
	{
		if (board[row][3]!=='- ')
			row-=1
		else
		{
			board[row][3]=symbol
			sendtoClient()
			makeBoard()
			setState()
			break;
		}
	}
	
}
var handleClick5 = function handleClick5() {
	let row = 5
	while(row>-1)
	{
		if (board[row][4]!=='- ')
			row-=1
		else
		{
			board[row][4]=symbol
			break;
		}
	}
	sendtoClient()
	makeBoard()
	setState()
}
var handleClick6 = function handleClick6() {
	let row = 5
	while(row>-1)
	{
		if (board[row][5]!=='- ')
			row-=1
		else
		{
			board[row][5]=symbol
			break;
		}
	}
	sendtoClient()
	makeBoard()
	setState()
}
var handleClick7 = function handleClick7() {
	let row = 5
	while(row>-1)
	{
		if (board[row][6]!=='- ')
			row-=1
		else
		{
			board[row][6]=symbol
			break;
		}
	}
	sendtoClient()
	makeBoard()
	setState()
}

const sendtoClient = () => {
	socket.emit('move',board)
}

var reactNode = React.createElement(
		'div',
		{ onClick: clickhandler, onMouseOver: mouseOverHandler },
		'click or mouse over'
	);

const formSubmit = event => {
	event.preventDefault()
	socket.emit('mymsg',state.message)
	console.log("sending msg from " + button.value + ": "+state.message);
	setState({message:''})
}

const initBoard = () => {
	for (let i=0; i<6; i++)
	{
		
		board[i]=[]

		for (let j=0; j<7; j++)
		{
			board[i][j]='- '}
		}
	}

	let grid=[]
	const makeBoard = () => {
		for (let i=0; i<6; i++)
		{
			let row=[]
			for (let j=0; j<7; j++)
			{	
				row[j] =React.createElement('input',{type: 'button',value: board[i][j],disabled: true});
			}
			grid[i] = React.createElement('div', null, row)
		}
	}

	socket.on('wait', data => 
	{
		ReactDOM.render(React.createElement('p', null, data), document.getElementById('root'))
	})

	socket.on('set', data => 
	{
		symbol=data;
	})
	socket.on('left', data => 
	{
		ReactDOM.render(React.createElement('p', null, data), document.getElementById('root'))
	})

	socket.on('start', () => {
		initBoard()
		makeBoard();
		setState()
	})

	socket.on('update', data => 
	{
		Object.assign(board,data)
		console.log(board)
		makeBoard()
		setState()
	})

	const setState = () => {

		ReactDOM.render(
			React.createElement('p',{},
				React.createElement('input',{type: 'submit', id: 'b1',value: 1,onClick: handleClick1}),
				React.createElement('input',{type: 'submit', id: 'b2',value: 2,onClick: handleClick2}),
				React.createElement('input',{type: 'submit', id: 'b3',value: 3,onClick: handleClick3}),
				React.createElement('input',{type: 'button', id: 'b4',value: 4,onClick: handleClick4}),
				React.createElement('input',{type: 'button', id: 'b5',value: 5,onClick: handleClick5}),
				React.createElement('input',{type: 'button', id: 'b6',value: 6,onClick: handleClick6}),
				React.createElement('input',{type: 'button', id: 'b7',value: 7,onClick: handleClick7}),
				React.createElement('div',{},grid),
			// reactNode
			//state.msgList.map(m=>React.createElement('div',null,m))
			),
			document.getElementById('root')
			)
	}




		