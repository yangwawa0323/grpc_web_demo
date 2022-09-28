import { EchoRequest, EchoResponse } from './js/echo_pb';
import { EchoServiceClient } from './js/echo_grpc_web_pb';

var echoService = new EchoServiceClient('http://localhost:8199');

var request = new EchoRequest();
request.setMessage('Hello World!');

echoService.echo(request, {}, function(err, response) {
  // ...
	const responseDiv = document.querySelector("#response")
	responseDiv.innerHTML = response
});