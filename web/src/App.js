import './App.css';

function App() {
  return (
    <div>
      <header>
      <nav className="navbar navbar-expand-sm navbar-toggleable-sm navbar-light bg-white border-bottom box-shadow mb-3">
          <div className="container">
              <a className="navbar-brand" asp-area="" asp-controller="Home" asp-action="Index">signalr</a>
              <button className="navbar-toggler" type="button" data-toggle="collapse" data-target=".navbar-collapse" aria-controls="navbarSupportedContent"
                      aria-expanded="false" aria-label="Toggle navigation">
                  <span className="navbar-toggler-icon"></span>
              </button>
              <div className="navbar-collapse collapse d-sm-inline-flex flex-sm-row-reverse">
                  <ul className="navbar-nav flex-grow-1">
                      <li className="nav-item">
                          <a className="nav-link text-dark" asp-area="" asp-controller="Home" asp-action="Index">Home</a>
                      </li>
                      <li className="nav-item">
                          <a className="nav-link text-dark" asp-area="" asp-controller="Home" asp-action="Privacy">Privacy</a>
                      </li>
                  </ul>
              </div>
          </div>
        </nav>
      </header>
      <div className="container">
          <main role="main" className="pb-3">
              <div className="text-center">
                  <h1 className="display-4">Welcome</h1>
              </div>
              <div className="container">
                  <div className="row">&nbsp;</div>
                  <div className="row">
                      <div className="col-2">User</div>
                      <div className="col-4"><input type="text" id="userInput" /></div>
                  </div>
                  <br/>
                  <div className="row">
                      <div className="col-2">Message</div>
                      <div className="col-4"><input type="text" id="messageInput" /></div>
                  </div>
                  <div className="row">&nbsp;</div>
                  <div className="row">
                      <div className="col-6">
                          <input type="button" id="sendButton" value="Send Message" />
                      </div>
                  </div>
              </div>
              <div className="row">
                  <div className="col-12">
                      <hr />
                  </div>
              </div>
              <div className="row">
                  <div className="col-6">
                      <ul id="messagesList"></ul>
                  </div>
              </div>
      </main>
      </div>

      <footer className="border-top footer text-muted">
          <div className="container">
              &copy; 2021 HomeSite - <a>Privacy</a>
          </div>
      </footer>
    </div>
  );
}

export default App;
