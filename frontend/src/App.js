import React from 'react';
import './App.css';
import {Route} from 'react-router';
import {withRouter, Switch} from "react-router-dom";
import LandingPage from "./pages/LandingPage";
import Header from "./components/common/Header";
import MainPage from "./pages/MainPage";
import UniversityPage from "./pages/UniversityPage";
import SubjectPage from "./pages/SubjectPage";
import MaterialPage from "./pages/MaterialPage";
import AddMaterialPage from "./pages/AddMaterialPage";

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
        <Route path="/subject/:id" render={(props) =>
          <SubjectPage id={props.match.params.id}/>
        }/>
        <Route path="/material/:id" render={(props) =>
          <MaterialPage id={props.match.params.id}/>
        }/>
        <Route path="/add_material/:id" render={(props) =>
          <AddMaterialPage id={props.match.params.id}/>
        }/>
      </Switch>
    </div>
  );
}

export default withRouter(App);
