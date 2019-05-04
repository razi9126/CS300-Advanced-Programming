const socket = io()
const state = {}
const board = [];
var symbol='-'
var instructions=""
var d =""

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
			sendtoClient()
			// checkwin({r:row, c:0})
			makeBoard()
			socket.emit('checkwin')
			setState()
			break;
		}
	}
};
var handleClick2 = function handleClick2() {
	let row = 5
	while(row>-1)
	{
		if (board[row][1]!=='- ')
			row-=1
		else
		{
			board[row][1]=symbol
			sendtoClient()
			makeBoard()
			socket.emit('checkwin')
			setState()
			break;
		}
	}

};
var handleClick3 = function handleClick3() {
	let row = 5
	while(row>-1)
	{
		if (board[row][2]!=='- ')
			row-=1
		else
		{
			board[row][2]=symbol
			sendtoClient()
			makeBoard()
			socket.emit('checkwin')
			setState()
			break;
		}
	}

};
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
			socket.emit('checkwin')
			setState()
			break;
		}
	}
	
};
var handleClick5 = function handleClick5() {
	let row = 5
	while(row>-1)
	{
		if (board[row][4]!=='- ')
			row-=1
		else
		{
			board[row][4]=symbol
			sendtoClient()
			makeBoard()
			socket.emit('checkwin')
			setState()
			break;
		}
	}

};
var handleClick6 = function handleClick6() {
	let row = 5
	while(row>-1)
	{
		if (board[row][5]!=='- ')
			row-=1
		else
		{
			board[row][5]=symbol
			sendtoClient()
			makeBoard()
			socket.emit('checkwin')
			setState()
			break;
		}
	}

};
var handleClick7 = function handleClick7() {
	let row = 5
	while(row>-1)
	{
		if (board[row][6]!=='- ')
			row-=1
		else
		{
			board[row][6]=symbol
			sendtoClient()
			makeBoard()
			socket.emit('checkwin')
			setState()
			break;
		}
	}

};

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
};

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
		ReactDOM.render(React.createElement('h1', null, data), document.getElementById('root'))
	})
	socket.on('draw', () => 
	{
		ReactDOM.render(React.createElement('h2', null, "Draw"), document.getElementById('root'))
	})

	socket.on('set', data => 
	{
		symbol=data;
	})
	
	socket.on('left', () => 
	{
		ReactDOM.render(React.createElement('h3', null, "Other player left the game. Refresh."), document.getElementById('root'))
	})
	socket.on('won', () => 
	{
		// console.log("WON")
		ReactDOM.render(React.createElement('h3', null, "You won!!"), document.getElementById('root'))
	})
	socket.on('lost', () => 
	{
		ReactDOM.render(React.createElement('h3', null, "You lost. Better luck next time :("), document.getElementById('root'))
	})


	socket.on('start', () => {
		initBoard()
		makeBoard();
		setState()
	})

	socket.on('update', data => 
	{
		Object.assign(board,data)
		// console.log(board)
		makeBoard()
		setState()
	})

	socket.on('yourmove', () => 
	{
		d= React.createElement('p',{},"It's your move ")
	})
	socket.on('notyourmove', () => 
	{
		d= React.createElement('p',{},"Please wait for opponent to make their move ")
	})

	const setState = () => {

		ReactDOM.render(
			React.createElement('p',{},
				React.createElement('h1',{},`CONNECT 4`),
				React.createElement('h3',{},`Select the column number to put the symbol in the corresponding column`),
				React.createElement('p',{},`Your symbol is ${symbol}`),
				React.createElement('input',{type: 'submit', id: 'b1',value: 1,onClick: handleClick1}),
				React.createElement('input',{type: 'submit', id: 'b2',value: 2,onClick: handleClick2}),
				React.createElement('input',{type: 'submit', id: 'b3',value: 3,onClick: handleClick3}),
				React.createElement('input',{type: 'button', id: 'b4',value: 4,onClick: handleClick4}),
				React.createElement('input',{type: 'button', id: 'b5',value: 5,onClick: handleClick5}),
				React.createElement('input',{type: 'button', id: 'b6',value: 6,onClick: handleClick6}),
				React.createElement('input',{type: 'button', id: 'b7',value: 7,onClick: handleClick7}),
				React.createElement('div',{},grid),   					//print the grid
				React.createElement('div',{},d),						//shows the pllayer eho has to make the current move
			// reactNode
			//state.msgList.map(m=>React.createElement('div',null,m))
			),
			document.getElementById('root')
			)
	}




		