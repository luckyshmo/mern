import React from 'react'
import {Switch, Route, Redirect} from 'react-router-dom'
import {WordsPage}  from './pages/WordsPage'
import {MainPage} from './pages/MainPage'
import {DetailPage} from './pages/DetailPage'
import {AuthPage} from './pages/AuthPage'

export const useRoutes = isAuth => {
    if (isAuth) {
        return (
            <Switch>
                <Route path="/words" exact>
                    <WordsPage />
                </Route>
                <Route path="/main" exact>
                    <MainPage />
                </Route>
                <Route path="/detail/:id">
                    <DetailPage />
                </Route>
                <Redirect to="/main" />
            </Switch>
        )
    }

    return (
        <Switch>
            <Route path="/" exact>
                <AuthPage />
            </Route>
            <Redirect to="/" />
        </Switch>
    )
}

export default useRoutes