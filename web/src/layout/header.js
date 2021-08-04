import {Link} from 'react-router-dom';

function Header() {
	return (
		<header>
			<nav className="navbar navbar-expand-sm navbar-toggleable-sm navbar-light bg-white border-bottom box-shadow mb-3">
				<div className="container">
					<a className="navbar-brand">Sigma</a>
					<button className="navbar-toggler" type="button" data-toggle="collapse" data-target=".navbar-collapse" aria-controls="navbarSupportedContent"
							aria-expanded="false" aria-label="Toggle navigation">
						<span className="navbar-toggler-icon"></span>
					</button>
					<div className="navbar-collapse collapse d-sm-inline-flex flex-sm-row-reverse">
						<ul className="navbar-nav flex-grow-1">
							<li className="nav-item">
								<Link to="/" className="nav-link text-dark">Home</Link>
							</li>
							<li className="nav-item">
								<Link to="/game" className="nav-link text-dark">Game</Link>
							</li>
						</ul>
					</div>
				</div>
			</nav>
		</header>
	);
}

export default Header;