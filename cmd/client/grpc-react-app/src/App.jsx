
import { useState } from 'react'
import './App.css'


import pb from './pb/echo_pb.js'
import client from './pb/echo_grpc_web_pb.js'

function App() {

  const [message, setMessage] = useState("")
  const [reply, setReply] = useState(null)
  
  
  const newEchoService = () => {
    return new client.EchoServiceClient('http://localhost:8199');
  }
  
  const grpc_request = () => {
    
    const echoService = newEchoService()
  
    const request = new pb.EchoRequest();
    
    if (message.trim() != "") {

      request.setMessage(message);
    
      echoService.echo(request, {}, function (err, response) {
        // ...
        setReply(response.array[0])
      })
    }
  }
    return (
    <div className="App">
      <div className='action'>
        <p class="title"> gRPC-web echo demo</p>
        <input onChange={(event) => setMessage(event.target.value)} value={message} /><br />
        <button onClick={grpc_request}>Send the message</button>
      </div>
      <div className='result'>
        {reply &&
          <p className="resultText">
            {reply}
          </p>
        }
      </div>
    </div>
  )
}

export default App
