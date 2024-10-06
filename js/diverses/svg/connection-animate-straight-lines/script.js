/* 

    connect svg elements automatically


*/

window.onload = () => {
    const svg = document.getElementById("master");
    const genXCircles = 12; // set how many circles should be created
    let circles = [];
    let lineSettings = {
        stroke: "rgba(4, 127, 180, 0.3)",
        "stroke-width": 1,
        class: "connection"
    }
    let currrentLines =[]; // all connnection lines will be stored in an array. for later reference(removal). this is faster as scanning the dom for the lines 

    const baseCircleSettings = {
        stroke: "rgba(4, 127, 180,1)",
        "stroke-width": 1,
        "fill-opacity": 1,
        fill: "rgba(4, 127, 180)",      
        r: 5
    }

    const conOrder  = [];
  
    const circlesSettings = [];
    const colorList = ["rgba(4, 127, 180,0.5)", "rgba(255, 51, 51,0.5)","rgba(255, 228, 107,0.5)","rgba(25, 250, 104,0.5)","rgba(0, 3, 194,0.5)"];

    function init() {
        for (i=0; i<genXCircles;i++) {
            let x  = document.createElementNS("http://www.w3.org/2000/svg", "circle");
            for (const [key, value] of Object.entries(baseCircleSettings)) {
                x.setAttribute(key, value);
            }
            x.setAttribute("id","c"+i);
          
            x.setAttribute("cx",getRandomMinMax(0, parseInt(svg.getAttribute("width"))-50));
            x.setAttribute("cy",getRandomMinMax(0, parseInt(svg.getAttribute("height"))-50));
            x.setAttribute("r",getRandomMinMax(2, 7));

            svg.appendChild(x);
            circles.push(x);

            if(i<genXCircles-1) conOrder.push(["c"+i, ["c"+(i+1),"c"+getRandomMinMax(0,genXCircles-1)]]);
            else conOrder.push(["c"+i, ["c0","c"+getRandomMinMax(0,genXCircles-1)]]);
        }
      
        //circles = document.getElementsByTagNameNS("http://www.w3.org/2000/svg","circle");
        
        for (circle of circles) {
            circlesSettings[circle.getAttribute("id")] = {
                vx: getRandomMinMax(1, 2)*getRandomDirection(), //speed vektor x
                vy: getRandomMinMax(1, 2)*getRandomDirection(), // speed vektor y
                r: parseFloat(circle.getAttribute("r"))
            };
        }
       
    }

    // connect circles with each other in described order
    function connectCirclesHandler() {
        //console.log("conorder", conOrder);

        conOrder.forEach((item) => {
            //console.log("conOrder item", item);
            // item[0] = the elements id to start with, item[1] list of element ids to connect to.
            let aElement = document.getElementById(item[0]);
            let aData = {
                x: parseFloat(aElement.getAttribute("cx")),
                y: parseFloat(aElement.getAttribute("cy")),
                r: parseFloat(aElement.getAttribute("r"))
            }
          

            item[1].forEach((targetId) => {
                //console.log("TargetID: ", targetId);
                let tElement = document.getElementById(targetId);
                let tData = {
                        x: parseFloat(tElement.getAttribute("cx")),
                        y: parseFloat(tElement.getAttribute("cy")),
                        r: parseFloat(tElement.getAttribute("r"))

                }

                let line = document.createElementNS("http://www.w3.org/2000/svg", "path");
                let d = calculateLine(aData,tData); // concatenated "d" string
                line.setAttribute("d",d);
                for (const [key, value] of Object.entries(lineSettings)) {
                    line.setAttribute(key, value);
                };
                svg.appendChild(line);
                currrentLines.push(line);
            });
        });

    }

    // calculate the line between 2 points a and b
    function calculateLine(a, b) {
        //console.log("received data: ", a,b)
        let line="M"; //corresponding to the 'd' in <path> 
        line+= a.x + " " + a.y; // add start coords
        line+= " l" + (b.x - a.x) + " " + (b.y - a.y); // add end coords
        
        //line+= (b.x - a.x)*(a.r/100) + " " + (b.y - a.y)*(a.r/100); // add start coords from stroke
        //line+= " l" + (b.x - a.x) + " " + (b.y - a.y); // add end coords to stroke

        return line;
    }

    function removeConnectionLines() {
       //const lines = document.getElementsByClassName("connection");
        for(line of currrentLines) {
            //line.setAttribute("stroke","rgba(0,0,0,0)");
            //line.setAttribute("stroke-width",0);
            line.remove();
        }

    }

    // move circles and deal with collosion
    function moveAround(){
        removeConnectionLines();
        for (circle of circles) {
            const id = circle.getAttribute("id");
            let pos = svg.createSVGPoint();
            pos.x = parseFloat(circle.getAttribute("cx"));
            pos.y = parseFloat(circle.getAttribute("cy"));
            if(pos.x + circlesSettings[id].r > svg.getAttribute("width") || pos.x - circlesSettings[id].r < 0 ) circlesSettings[id].vx = -(circlesSettings[id].vx);
            if(pos.y + circlesSettings[id].r > svg.getAttribute("height") || pos.y - circlesSettings[id].r < 0 ) circlesSettings[id].vy = -(circlesSettings[id].vy);
            let newPos = {
                x: pos.x + circlesSettings[id].vx,
                y: pos.y + circlesSettings[id].vy
            }
            //console.log(pos, newPos);
            circle.setAttribute("cx", parseInt(newPos.x));
            circle.setAttribute("cy", parseInt(newPos.y));
        }
      
        connectCirclesHandler();
    }

    function changeColor() {
        console.log("change");
        color = colorList[getRandomMinMax(0, colorList.length-1)];
        lineSettings.stroke = color;
        for ( circle of circles ) {
           
            circle.setAttribute("fill",color);
            circle.setAttribute("stroke",color);
        }

        for ( line of currrentLines) {
            line.setAttribute("stroke", color)
        }
    }

    function getRandomDirection() {
        
        const min = 1;
        const max = 2;
        const rand =  Math.floor(Math.random() * (max - min +1)) + min; 
        if(rand == 1)return 1;
        else if(rand == 2) return -1;
    }

    function getRandomMinMax(mi, ma) {
        
        const min = Math.ceil(mi);
        const max = Math.floor(ma);
        return rand =  Math.floor(Math.random() * (max - min +1)) + min; 
        
    } 


    init();
    
    flow = setInterval(moveAround,17);
    //colorize = setInterval(changeColor,2000);

}

function stop() {
    clearInterval(flow);
    clearInterval(colorize);
    console.log("stopped: ", flow);
}