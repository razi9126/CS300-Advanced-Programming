const socket = io()
const state = {}

socket.on('yourmsg',data=> setState({msgList: [...state.msgList,data]}))

const formSubmit = event => {
	event.preventDefault()
	socket.emit('mymsg',state.message)
	console.log("sending msg from " + socket.id + ": "+state.message);
	setState({message:''})
}

const setState = updates => {
	Object.assign(state, updates)
	ReactDOM.render(
		React.createElement('form',{onSubmit: formSubmit},
			React.createElement('input', {
				value: state.message,
				onChange: event => setState({message: event.target.value}),
				type: 'text'
			}),
			React.createElement('input',{type:'submit',value: 'send'}),
			state.msgList.map(m=>React.createElement('div',null,m))),
		document.getElementById('root'))
}


setState({message:'',msgList:[]})



		