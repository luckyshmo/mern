import React from 'react'
import 'materialize-css'
import {useRoutes} from './routes.js'
import {useAuth} from './hooks/auth.hook'
import {BrowserRouter as Router} from 'react-router-dom'
import { AuthContext } from './context/AuthContext.js'
import { Navbar } from './components/Navbar.js'
// import {Router} from 'express';

function App() {
  const {token, login, logout, userId} = useAuth()
  const isAuthenticated = !!token //fucking js magic
  const routes = useRoutes(isAuthenticated)
  return (
    <AuthContext.Provider value={{ //context for whole app
      token, login, logout, userId
    }}>
      <Router>
        { isAuthenticated && <Navbar/> /*fucking magical js again!! */ } 
        <div className="container">
          {routes}
        </div>
      </Router>
    </AuthContext.Provider>
  );
}

export default App
