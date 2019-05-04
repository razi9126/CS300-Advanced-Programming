const url = 'http://api.openweathermap.org/data/2.5/forecast?id=524901&APPID=9e02c65d2effce27701fb6cff0753395'
let temp = 0
let city = ''

const formSubmit = async ev => {
	ev.preventDefault()
	const resp = await fetch(`${url}&q=${city}`)
	temp = (await resp.json()).main.temp
	redraw()
}

const redraw = () =>
ReactDOM.render(
	React.createElement('form',{onSubmit: formSubmit},
		React.createElement('input', {
				type:'text',
				onChange: ev => city = ev.target.value

			}),
		React.createElement('input', {type: 'submit'}),
		React.createElement('h2', null, `T: ${temp}`)),
	document.getElementById('root'))

redraw()




		