const fs = require('fs')
const http = require('http')
const socketio = require('socket.io')
const data = []
const readFile = f => new Promise((resolve,reject) =>
	fs.readFile(f, (e, d) => e? reject(e):resolve(d)))


const server = http.createServer( async (req,resp) =>
	resp.end( await readFile(req.url.substr(1))))

//changes
let players = {}
let boards = {}
let waiting = null
let gameId = 0 
let p1 =""
let p2 =""
let win = false;  		                    //winning flag

let full = false

turn = p1                                   //player1 gets first turn by default
const switchturn = () => {
	if (turn=== p1)
		turn = p2
	else
		turn = p1
}

const io = socketio(server)
io.sockets.on('connection', socket=>
{
	players[socket.id] = gameId
	socket.join(players[socket.id])
	           
	if (!waiting)
	{
		waiting = socket
		console.log("p1")
		console.log(socket.id)
		p1=socket.id
		socket.emit('set', 'X')	  						  //player1 always gets X as icon
		socket.emit('wait', 'Waiting for player 2')
	}
	else
	{
		socket.emit('set', 'O')
		boards[players[socket.id]] = create()
		turn=p1
		p2=socket.id
		console.log("p2")
		console.log(socket.id)
		io.to(players[socket.id]).emit('start')
		
		waiting = null                                    //to allow 2 more players to connect
		gameId++
	}


	socket.on('move', array=>
	{
		if (socket.id == turn)
		{
			Object.assign(myBoard, array)
			switchturn()
			if(turn==p1)
			{
				io.to(p1).emit('yourmove')
				io.to(p2).emit('notyourmove')
			}
			else
			{
				io.to(p2).emit('yourmove')
				io.to(p1).emit('notyourmove')
			}
			// socket.emit('yourmove')
		}
		// io.sockets.emit('update',myBoard)
		io.to(players[socket.id]).emit('update',myBoard)
	})

	socket.on('disconnect', () => {
		if (waiting === socket) waiting = null
		io.sockets.emit('left')
		delete players[socket.id]
	})
	socket.on('checkwin', () => {
		// console.log("checking win"))
		if (checkVertical() || checkHorizontal() || checkDiagonal())
		{
			if (socket.id == p1)
			{
				io.to(p1).emit('won')
				io.to(p2).emit('lost')
				win=false
			}
			else
			{
				io.to(p2).emit('won')
				io.to(p1).emit('lost')
				win=false
			}
		}
		if (gridFull())
			io.sockets.emit('draw')

	})
})
const gridFull= () =>{
	var tally=0
	for (var y = 0; y <= 5; y++) {
        for (var x = 0; x <= 6; x++) {
            currentValue = myBoard[y][x];
            if (currentValue !== 'X' && currentValue !== 'O' ) {
                tally += 1;
            }
            if (tally>0) {
                return false;
            }
            // previousValue = currentValue;
        }

    }
    return true;



	}
const checkDiagonal= () =>{
	var x = null,
	y = null,
	xtemp = null,
	ytemp = null,
	currentValue = null,
	previousValue = 0,
	tally = 0;

    // Test for down-right diagonals across the top.
    for (x = 0; x <= 6; x++) {
    	xtemp = x;
    	ytemp = 0;

    	while (xtemp <= 6 && ytemp <= 5) {
    		currentValue =myBoard[ytemp][xtemp];
    		if (currentValue === previousValue && currentValue !=='- ') {
    			tally += 1;
    		} else {
                // Reset the tally if you find a gap.
                tally = 0;
            }
            if (tally === 4 - 1) {

            	return true;
            }
            previousValue = currentValue;

            // Shift down-right one diagonal index.
            xtemp++;
            ytemp++;
        }
        // Reset the tally and previous value when changing diagonals.
        tally = 0;
        previousValue = 0;
    }

    // Test for down-left diagonals across the top.
    for (x = 0; x <= 6; x++) {
    	xtemp = x;
    	ytemp = 0;

    	while (0 <= xtemp && ytemp <= 5) {
    		currentValue = myBoard[ytemp][xtemp];
    		if (currentValue === previousValue && currentValue !== '- ') {
    			tally += 1;
    		} else {
                // Reset the tally if you find a gap.
                tally = 0;
            }
            if (tally === 3) {
            	return true;
            }
            previousValue = currentValue;

            // Shift down-left one diagonal index.
            xtemp--;
            ytemp++;
        }
        // Reset the tally and previous value when changing diagonals.
        tally = 0;
        previousValue = 0;
    }

    // Test for down-right diagonals down the left side.
    for (y = 0; y <= 5; y++) {
    	xtemp = 0;
    	ytemp = y;

    	while (xtemp <= 6 && ytemp <= 5) {
    		currentValue = myBoard[ytemp][xtemp];
    		if (currentValue === previousValue && currentValue !== '- ') {
    			tally += 1;
    		} else {
                // Reset the tally if you find a gap.
                tally = 0;
            }
            if (tally === 3) {
            	return true;
            }
            previousValue = currentValue;

            // Shift down-right one diagonal index.
            xtemp++;
            ytemp++;
        }
        // Reset the tally and previous value when changing diagonals.
        tally = 0;
        previousValue = 0;
    }

    // Test for down-left diagonals down the right side.
    for (y = 0; y <= 5; y++) {
    	xtemp = 6;
    	ytemp = y;

    	while (0 <= xtemp && ytemp <= 5) {
    		currentValue = myBoard[ytemp][xtemp];
    		if (currentValue === previousValue && currentValue !== '- ') {
    			tally += 1;
    		} else {
                // Reset the tally if you find a gap.
                tally = 0;
            }
            if (tally === 3) {
            	return true;
            }
            previousValue = currentValue;

            // Shift down-left one diagonal index.
            xtemp--;
            ytemp++;
        }
        // Reset the tally and previous value when changing diagonals.
        tally = 0;
        previousValue = 0;
    }

    // No diagonal wins found. Return false.
    return false;
}
const checkVertical = () => {
	var currentValue = null,
        previousValue = 0,
        tally = 0;

    // Scan each column in series, tallying the length of each series. If a
    // series ever reaches four, return true for a win.
    for (var x = 0; x <= 6; x++) {
        for (var y = 0; y <= 5; y++) {
            currentValue = myBoard[y][x];
            if (currentValue === previousValue && currentValue !== '- ') {
                tally += 1;
            } else {
                // Reset the tally if you find a gap.
                tally = 0;
            }
            if (tally === 3) {
                return true;
            }
            previousValue = currentValue;
        }

        // After each column, reset the tally and previous value.
        tally = 0;
        previousValue = 0;
    }

    // No vertical win was found.
    return false;
}

const checkHorizontal = () => {
	var currentValue = null,
        previousValue = 0,
        tally = 0;

    // Scan each row in series, tallying the length of each series. If a series
    // ever reaches four, return true for a win.
    for (var y = 0; y <= 5; y++) {
        for (var x = 0; x <= 6; x++) {
            currentValue = myBoard[y][x];
            if (currentValue === previousValue && currentValue !== '- ') {
                tally += 1;
            } else {
                // Reset the tally if you find a gap.
                tally = 0;
            }
            if (tally === 3) {
                return true;
            }
            previousValue = currentValue;
        }

        // After each row, reset the tally and previous value.
        tally = 0;
        previousValue = 0;
    }

    // No horizontal win was found.
    return false;
}

myBoard =[]
const create = () => {
	for (let i=0; i<6; i++)
	{
		
		myBoard[i]=[]

		for (let j=0; j<7; j++)
		{
			
			myBoard[i][j]='- '
		}
	}

}

server.listen(8000, () => 
	{
		console.log('Started')
		create()
	})