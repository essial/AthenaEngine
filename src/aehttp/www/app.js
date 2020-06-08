setTimeout(load, 250)

function load() {

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
}
