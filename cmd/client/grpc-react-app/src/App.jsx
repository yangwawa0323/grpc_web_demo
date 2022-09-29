import React from 'react'
import {Routes, Route} from 'react-router-dom'
import Home from './component/home/Home'
import EchoPanel from './component/echo/EchoPanel'
import UserPanel from './component/user/UserPanel'

const App = () => {
	return (
		<Routes>
			<Route path="/" element={<Home />} >
				<Route path="echo" element={<EchoPanel />}></Route>
				<Route path="user" element={<UserPanel />}></Route>
			</Route>
	  </Routes>
  )
}

export default App