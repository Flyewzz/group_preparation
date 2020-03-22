import React from 'react';
import './App.css';
import {Route} from 'react-router';
import {withRouter, Switch} from "react-router-dom";
import LandingPage from "./pages/LandingPage";
import Header from "./components/common/Header";
import MainPage from "./pages/MainPage";
import UniversityPage from "./pages/UniversityPage";

function App() {
  return (
    <div>
      <Header/>
      <Switch>
        <Route path="/welcome">
          <LandingPage/>
        </Route>
        <Route exact path="/">
          <MainPage/>
        </Route>
        <Route path="/university/:id" render={(props) =>
          <UniversityPage id={props.match.params.id}/>
        }/>
      </Switch>
    </div>
  );
}

export default withRouter(App);
