import { useState } from "react";

export default function Game() {
	const [history, setHistory] = useState([Array(9).fill(null)]);
	const currentSquare = history[history.length - 1];

	let squareRows = [];
	let boardRows = [];

	for (let i = 0; i < currentSquare.length; i++) {
		squareRows.push(currentSquare[i]);
		if ((i + 1) % 3 === 0) {
			const squares = squareRows.map((square, index) => {
				return (
					<div className="square" key={"sqaure" + index}>
						{square}
					</div>
				);
			});
			const boardRow = <div className="board-row">{squares}</div>;
			squareRows = [];
			boardRows.push(boardRow);
		}
	}

	return (
		<>
			<div className="game">
				<div className="game-board">
					<div className="status">Next Player: X</div>
					{boardRows.map((boardRow) => {
						{
							boardRow;
						}
					})}

					<h1>Halow</h1>
				</div>
				<div className="game-info">
					<ol>
						<li>
							<button onClick={() => console.log("button clicked")}>
								Go to game start
							</button>
						</li>
					</ol>
				</div>
			</div>
		</>
	);
}
