import { useEffect, useRef, useState } from "react";

export default function App() {
	const [count, setCount] = useState(0);
	const countRef = useRef(1);

	function handleClickDecrement() {
		setCount((currentCount) => {
			return currentCount - 1;
		});
	}

	function handleClickIncrement() {
		setCount((currentCount) => {
			return currentCount + 1;
		});
	}

	function handleRefDecrement() {
		console.log(countRef.current);

		countRef.current = countRef.current + 1;
	}

	console.log("render");

	useEffect(() => {
		console.log("side effect");
		handleRefDecrement();
	});

	return (
		<>
			<h2>Use State</h2>
			<button onClick={handleClickDecrement}>-</button>
			<span>{count}</span>
			<button onClick={handleClickIncrement}>+</button>
			<br />
			<br />
			<h2>Use Ref</h2>
			<button onClick={handleRefDecrement}>-</button>
			<span>{countRef.current}</span>
			<button onClick={handleClickIncrement}>+</button>
		</>
	);
}
