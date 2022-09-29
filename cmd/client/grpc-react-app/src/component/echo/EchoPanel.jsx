
import { useState } from "react"


import pb from "../../../../js/echo/v1/echo_pb.js"
import client from "../../../../js/echo/v1/echo_grpc_web_pb.js"
import { TextField, Box, Button, Typography } from "@mui/material"

function EchoPanel() {

  const [message, setMessage] = useState("")
  const [reply, setReply] = useState(null)


  const newEchoService = () => {
    return new client.EchoServiceClient("http://localhost:8199");
  }

  const grpc_request = () => {

    const echoService = newEchoService()

    const request = new pb.EchoRequest();

    if (message.trim() != "") {

      request.setMessage(message);

      echoService.echo(request, {}, function (err, response) {
        // ...
        if (err === null) {
          setReply(response.array[0])
        } else {
          setReply(err.message)
        }
      })
    }
  }
  return (
    <Box sx={{}}>
      <Box sx={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        flexDirection: "column"

      }}>
        <Box sx={{ padding: 6 }}>
          <Typography variant="body2">
            Echo gRPC is simplest gRPC request and response demostration.
          </Typography>
        </Box>
        <Box>

          <TextField
            sx={{

            }}
            size="small"
            variant="outlined"
            label="message"
            onChange={(event) => setMessage(event.target.value)}
            value={message} />
          <Button sx={{
            marginLeft: 3
          }}
            variant="contained"
            size="small"
            onClick={grpc_request}>
            Send
          </Button>
        </Box>
      </Box>
      <Box sx={{ marginTop: 3, display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
        {reply &&
          <Box sx={{
            opacity: 0.7
          }}>
            {reply}
          </Box>
        }
      </Box>
    </Box>
  )
}

export default EchoPanel
