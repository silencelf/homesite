import './App.css';
import Container from './layout/Home';
import Footer from './layout/footer';
import Header from './layout/header';
import Game from './tictactoe/game';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";
import Home from './layout/Home';


function App() {
  return (
    <Router>
      <div>
        <Header />
        <div className="container">
          <Switch>
            <Route path="/game">
              <Game />
            </Route>
            <Route path="/">
              <Home />
            </Route>
          </Switch>
        </div>
        <Footer />
      </div>
    </Router>
  );
}

export default App;
