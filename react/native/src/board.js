import Square from "./square";

export default function Board({ squares, xIsNext, handlePlay }) {
	const winner = calculateWinner();
	const status = winner
		? "Winner: " + winner
		: "Next Player: " + (xIsNext ? "X" : "O");

	let squareRows = [];
	let boardRows = [];
	for (let i = 0; i < squares.length; i++) {
		squareRows.push(
			<Square
				value={squares[i]}
				key={"square" + i}
				onSquareClick={() => onSquareClick(i)}
			/>
		);
		if ((i + 1) % 3 === 0) {
			boardRows.push(
				<div className="board-row" key={"board" + i}>
					{squareRows}
				</div>
			);
			squareRows = [];
		}
	}

	function onSquareClick(squareIndex) {
		if (squares[squareIndex] || calculateWinner()) return;

		const squareNext = squares.slice();
		squareNext[squareIndex] = xIsNext ? "X" : "O";

		handlePlay(squareNext);
	}

	function calculateWinner() {
		const lineWinners = [
			[0, 1, 2],
			[0, 4, 8],
			[0, 3, 6],
			[1, 4, 7],
			[2, 5, 8],
			[2, 4, 6],
			[3, 4, 5],
			[6, 7, 8],
		];

		for (let i = 0; i < lineWinners.length; i++) {
			const lineWinner = lineWinners[i];

			if (
				squares[lineWinner[0]] &&
				squares[lineWinner[0]] === squares[lineWinner[1]] &&
				squares[lineWinner[0]] === squares[lineWinner[2]]
			) {
				return squares[lineWinner[0]];
			}
		}

		return null;
	}

	return (
		<>
			<div className="status">{status}</div>
			{boardRows}
		</>
	);
}
