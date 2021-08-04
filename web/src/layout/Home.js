function Home(props) {
	return (
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
					<ul id="messages"></ul>
				</div>
			</div>
		</main>
	);
}

export default Home;