(document => {
	const HIDDEN = "brightness(0)",
		VISIBLE = "brightness(1)",
		jeuMemoryName = document.getElementById("jeuMemoryName"),
		jeuMemoryCarte = document.getElementById("jeuMemoryCarte"),
		drapeaux = [
			["asexuel", "asexuel"],
			["bisexuel", "bisexuel"],
			["gay", "gay"],
			["lesbien_feministe", "lesbien et feministe"],
			["lesbien", "lesbien"],
			["non_binaire", "non binaire"],
			["progress", "LGBT en progrès"],
			["transgenre", "transgenre"]
		]
			.map(x => [x, x])
			.flat()
			.map(([key, name]) => ({
				k: key,
				n: "Drapeau " + name,
				// The container of the image
				d: document.createElement("div"),
				// The image
				i: new Image(),
				// resolved
				r: false
			}))
			.sort(_ => Math.random() < 0.9);

	let other = null,
		jammed = false;

	for (const drapeau of drapeaux) {
		jeuMemoryCarte.append(drapeau.d);
		drapeau.d.append(drapeau.i);
		drapeau.i.src = "drapeau-" + drapeau.k + ".png";
		drapeau.i.style.filter = HIDDEN;
		drapeau.d.onclick = _ => {
			if (jammed || drapeau.r || other == drapeau) return;
			jeuMemoryName.innerText = drapeau.n;
			if (other) {
				if (other.k == drapeau.k) {
					jeuMemoryName.innerText = "";
					other.r = drapeau.r = true;
					other = null;
					if (drapeaux.every(d => d.r)) {
						jeuMemoryName.innerText = "Vous avez gagné!";
					}
				} else {
					jammed = true;
					setTimeout(_ => {
						jammed = false;
						other.i.style.filter = drapeau.i.style.filter = HIDDEN;
						other = null;
					}, 500);
				}
			} else {
				other = drapeau;
			}

			drapeau.i.style.filter = VISIBLE;
		};
	}
})(document);
