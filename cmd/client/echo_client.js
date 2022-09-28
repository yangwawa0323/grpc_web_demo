const {EchoRequest, EchoResponse} = require('./js/echo_pb');
const {EchoServiceClient} = require('./js/echo_grpc_web_pb');

var echoService = new EchoServiceClient('http://localhost:8199');

var request = new EchoRequest();
request.setMessage('Hello World!');

echoService.echo(request, {}, function(err, response) {
  // ...
	const responseDiv = document.querySelector("#response")
	responseDiv.innerHTML = response
});