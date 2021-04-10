import React, { useContext, useEffect, useState } from 'react'
import { useHttp } from '../hooks/http.hook'
import { useMessage } from '../hooks/message.hook'
import { AuthContext } from '../context/AuthContext'

export const AuthPage = () => {
    const auth = useContext(AuthContext)
    const message = useMessage()
    const {loading, error, request, clearError} = useHttp()
    const [form, setForm] = useState({
        email: '', password: ''
    })

    //error controller for user. works with "message", but can work with err[]
    useEffect( () => {
        message(error)
        clearError()
      },
      [error, message, clearError]
    )

    const changeHandler = event => {
        setForm({ ...form, [event.target.name]: event.target.value })
    }

    const registerHandler = async () => {
        try {
            const data = await request('api/auth/sign-up', 'POST', {...form})
            message(data.message)
        } catch (e) { }
    }

    const loginHandler = async () => {
      try {
          const data = await request('api/auth/sign-in', 'POST', {...form})
          auth.login(data.token, data.userID)
      } catch (e) { }
    }

    return (
        <div className="row">
          <div className="col s6 offset-s3">
            <h1>Shorten</h1>
            <div className="card blue darken-1">
              <div className="card-content white-text">
                <span className="card-title">Authorization</span>
                <div>
    
                  <div className="input-field">
                    <input
                      id="email"
                      type="text"
                      name="email"
                      className="yellow-input"
                    //   value={form.email}
                      onChange={changeHandler}
                    />
                    <label htmlFor="email">Email</label>
                  </div>
    
                  <div className="input-field">
                    <input
                      id="password"
                      type="password"
                      name="password"
                      className="yellow-input"
                    //   value={form.password}
                      onChange={changeHandler}
                    />
                    <label htmlFor="email">Password</label>
                  </div>
    
                </div>
              </div>
              <div className="card-action">
                <button
                  className="btn yellow darken-4"
                  style={{marginRight: 10}}
                  disabled={loading}
                  onClick={loginHandler}
                >
                  Enter
                </button>
                <button
                  className="btn grey lighten-1 black-text"
                  onClick={registerHandler}
                  disabled={loading}
                >
                  Register
                </button>
              </div>
            </div>
          </div>
        </div>
      )
}