import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Books from './Books';
import Members from './Members';

const App = () => {
  return (
    <Router>
      <Switch>
        <Route path="/books" component={Books} />
        <Route path="/members" component={Members} />
      </Switch>
    </Router>
  );
};

export default App;