var keysPressed = {};

document.addEventListener('keydown', (event) => {
    keysPressed[event.key] = true;
});

document.addEventListener('keyup', (event) => {
    delete keysPressed[event.key];
});

function showError(error) {
	const errordiv = document.getElementById("errordiv");
	errordiv.innerHTML = error;
	console.log("Error:", error);
}

function clearCanvas() {
	const canvas = document.getElementById("drawspace");
	const ctx = canvas.getContext("2d");
	ctx.clearRect(0, 0, canvas.width, canvas.height);
}

function drawRect(x, y, w, h, c) { //.. color
	const canvas = document.getElementById("drawspace");
	const ctx = canvas.getContext("2d");
	ctx.strokeStyle = c;
	ctx.strokeRect(x, y, w, h);
}

function drawFillRect(x, y, w, h, c) { //.. color
	const canvas = document.getElementById("drawspace");
	const ctx = canvas.getContext("2d");
	ctx.fillStyle = c;
	ctx.fillRect(x, y, w, h);
}

function drawLine(x1, y1, x2, y2, c) { //.. color
	const canvas = document.getElementById("drawspace");
	const ctx = canvas.getContext("2d");
	ctx.beginPath();
	ctx.moveTo(x1, y1);
	ctx.lineTo(x2, y2);
	ctx.strokeStyle = c;
	ctx.stroke();
}

function drawCircle(x, y, r, c) { //x, y, radius, color
	const canvas = document.getElementById("drawspace");
	const ctx = canvas.getContext("2d");
	ctx.beginPath();
	ctx.strokeStyle = c;
	ctx.arc(x, y, r, 0, 2 * Math.PI);
	ctx.stroke();
}

function drawPartialCircle(x, y, r, s, e, c) { //x, y, radius, start, end, color
	const canvas = document.getElementById("drawspace");
	const ctx = canvas.getContext("2d");
	ctx.beginPath();
	ctx.fillStyle = c;
	ctx.arc(x, y, r, s * Math.PI, e * Math.PI);
	ctx.stroke();
}

function drawText(x, y, s, t, c) { // x, y, size, text, color
	const canvas = document.getElementById("drawspace");
	const ctx = canvas.getContext("2d");
	ctx.font = s + "px Arial"
	ctx.fillStyle = c;
	ctx.fillText(t, x, y)
}