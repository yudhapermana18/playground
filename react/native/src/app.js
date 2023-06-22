import { useState } from "react";
import Board from "./board";

export default function Game() {
	const [history, setHistory] = useState([Array(9).fill(null)]);
	const [currentSquare, setCurrentSquare] = useState(0);
	const xIsNext = currentSquare % 2 === 0;

	function onPlay(nextSquare) {
		const historyNext = [...history.slice(0, currentSquare + 1), nextSquare];
		setHistory(historyNext);
		setCurrentSquare(historyNext.length - 1);
	}

	function onHistoryMove(historyMove) {
		setCurrentSquare(historyMove);
	}

	const historyMoves = history.map((_, move) => {
		return (
			<li key={move}>
				<button onClick={() => onHistoryMove(move)}>
					{move > 0 ? "Go to move #" + move : "Go to game start"}
				</button>
			</li>
		);
	});

	return (
		<>
			<div className="game">
				<div className="game-board">
					<Board
						squares={history[currentSquare]}
						xIsNext={xIsNext}
						handlePlay={onPlay}
					/>
				</div>
				<div className="game-info">
					<ol>{historyMoves}</ol>;
				</div>
			</div>
		</>
	);
}
