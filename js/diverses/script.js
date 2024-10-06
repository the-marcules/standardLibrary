var doDrag = function(evt) {
	evt.preventDefault();
	// Box Position und Maße (Die Werte kommen mit "px" am ende, dies wird abgeschnitten)
	targetW = evt.target.style.width.substring(0, evt.target.style.width.length-2);
	targetH = evt.target.style.height.substring(0, evt.target.style.height.length-2);
	targetPosX = evt.target.style.left.substring(0, evt.target.style.left.length-2);
	targetPosY = evt.target.style.top.substring(0, evt.target.style.top.length-2);

	// Berechnung neue Koordinaten für CSS
	// posLeft = evt.clientX - (targetW/2);
	// posTop = evt.clientY - (targetH/2);
	posLeft = evt.clientX - (evt.clientX - targetPosX);
	posTop = evt.clientY - (evt.clientY - targetPosY);
	
	console.log(posLeft + " / " + posTop)
	
	//Setzen der neuen Position + Effekte
	evt.target.style.left = posLeft+"px";
	evt.target.style.top = posTop+"px";
	evt.target.innerText = "Drop mich!";
	evt.target.style.boxShadow = "2px 2px 5px grey";
	
}

function setup() {
	var myP = document.getElementById("demo");
	myP.style.backgroundColor = "#6699CC";
	myP.style.border = "1px solid #333";
	myP.style.display = "block";
	myP.style.width = "100px";
	myP.style.height = "100px";
	myP.style.position = "absolute";
	myP.style.top = "20px";
	myP.style.left = "10px";
	myP.style.padding = "3px";

	marcules.drag(myP);
}


function enableDrag(evt) {
	evt.preventDefault();
	evt.target.addEventListener("mousemove",doDrag);
	//console.log("MouseDOWN: Drag aktiviert.");
	
}

function disableDrag(evt) {
	evt.target.removeEventListener("mousemove",doDrag);
	// console.log("MouseUP: Drag DEaktiviert.");
	evt.target.innerText = "Drag mich!";
	evt.target.style.boxShadow = "none";
}

var marcules = {
	drag: function (obj) {
		
		obj.addEventListener("mousedown",enableDrag);
		obj.addEventListener("mouseup",disableDrag);
	}
}