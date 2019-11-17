import React from 'react';
import 'semantic-ui-css/semantic.min.css'
// import 'semantic-ui-css/semantic.js';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import Login from './components/Authentication/Login';

import Main from './components/Home/Main';

function App() {
  return (
    <div className="App">
      <Router>

        <Switch>
          <Route path='/login'>
            <Login />
          </Route>
        </Switch>

        <Switch>
          <Route path='/home'>
            <Main />
          </Route>
        </Switch>

      </Router>
    </div>
  );
}

export default App;
