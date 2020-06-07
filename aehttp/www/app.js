const term = new Terminal({
	rows: 30,
	fontFamily: 'Oxygen Mono',
	fontSize: 14,
	renderType: 'DOM'
})

var buffer = '';
const serverAddress = 'ws://localhost:3000/wsconnect'

term.open(document.getElementById('terminal'))
term.writeln('Connecting...')
try {
	webSocket = new WebSocket(serverAddress)
} catch (e) {
	term.writeln(e.message)
}
webSocket.onopen = function () {
	term.writeln("Connection established")
}

webSocket.onmessage = function (e) {
	term.writeln("Server: " + e.data)
}
/*
term.onKey(({key, domEvent}) => {
	if (key === '\r') {
		// Newline sent
		if (buffer.length > 0) {
			webSocket.send(buffer)
		}
		buffer = ''
		term.write('\r\n')
		return
	}
	if (domEvent.key === 'Backspace') {
		// Remove a char
		if (buffer.length > 0) {
			buffer = buffer.slice(0, -1)
		}
		term.write('\b')
		return
	}

	buffer += key
	term.write(key)
})
*/

$("#termInput").on('keypress', function(e) {
	if (e.which === 13) {
		const v = $(this).val()
		if (v.length === 0) {
			return
		}
		webSocket.send(v)
		$(this).val('')
	}
})
