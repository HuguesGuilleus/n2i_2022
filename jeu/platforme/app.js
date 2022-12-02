(document => {
	const canvasElement = document.getElementById("jeuCanvas"),
		canvasContext = canvasElement.getContext("2d"),
		canvasWith = canvasElement.width,
		canvasHeight = canvasElement.height,
		VAGINA = new Image(),
		VIH = new Image(),
		virus = [];

	VAGINA.src = "/img/vagina.png";
	VIH.src = "/img/vih.png";

	for (let i = 0, position = 600; i < 10; i++) {
		virus.push(position);
		position += 200 + Math.random() * 150;
	}

	let vaginaTime = 0,
		interval = setInterval(_ => {
			if (vaginaTime) vaginaTime--;

			const height = 50 + 1.5 * (vaginaTime - 10) ** 2;
			for (const i in virus) {
				virus[i] -= 10;
				if (height > 170 && virus[i] < 25) {
					console.log("bad end!");
					clearInterval(interval);
					interval = 0;
					return;
				}
			}
			if (virus[0] < 0) {
				virus.shift();
			}
			if (!virus.length) {
				console.log("happy end!");
				clearInterval(interval);
				interval = 0;
			}
		}, 40);

	window.onkeydown = event => {
		console.log(event.key);
		if (event.key == " ") {
			if (!vaginaTime) {
				vaginaTime = 20;
				event.preventDefault();
				event.stopPropagation();
			}
		}
	};

	canvasContext.fillStyle = "grey";
	function draw() {
		if (interval) requestAnimationFrame(draw);

		// Clean
		canvasContext.fillRect(0, 0, canvasWith, canvasHeight);

		// Draw vagina
		const height = 1.5 * (vaginaTime - 10) ** 2;
		canvasContext.drawImage(VAGINA, 0, height, 50, 50);

		// draw VIH
		for (const v of virus) {
			canvasContext.drawImage(VIH, v, 170, 25, 25);
		}
	}
	requestAnimationFrame(draw);
})(document);
