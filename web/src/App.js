import './App.css';
import Container from './layout/container';
import Game from './tictactoe/game'
import Footer from './layout/footer';
import Header from './layout/header';

function App() {
  return (
    <div>
        <Header/>
        <Game />
        <Footer/>
    </div>
  );
}

export default App;
