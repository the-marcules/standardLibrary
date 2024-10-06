window.onload = function () {
    
    function randomize() {
        try {

            const allRects = document.getElementsByTagName("rect");
            const allRectsCount = allRects.length;
            for(let item of allRects) {
                item.setAttribute("stroke", "rgb(11,11,11)");
                item.setAttribute("stroke-width", "1");
            }

            
            loop =setInterval(()=> {
                let myRect = document.getElementById("myRectum"+getRandNum(1,allRectsCount));
                myRect.setAttribute("fill","url(#grad"+getRandNum(1,3));
                //console.log("Updated " + myRect.id + " with fill " + myRect.getAttribute("fill"));
            },2000);

        } catch (e) {
            console.log("Error: " + e);
        }
    }

    function MySVGEditor(svgID) {
        const svg = document.getElementById(svgID);
        let mySelection = [];
        let mySelSettings = [];
        let shiftPressed = false;
        let mouseOffsetX = 0;
        let mouseOffsetY = 0;


        function init() {
            document.onkeydown = keyDownHandler;
            document.onkeyup =  keyUpHandler;
        

            /* svg.addEventListener("click", (event) => {
                            console.log("SVG geklickt" + event.target);
                            removeFromSelection();
                        }); */
            elements = svg.children;
            for (let element of elements) {
                if (element != undefined) {
                    //element.addEventListener('click', select);
                    //element.addEventListener('mousedown', mouseDownHandler);
                    //element.addEventListener('mouseup', mouseUpHandler);
                    element.addEventListener('click', select);
                   // 
                }
            }
          
        }

     

        function getMousePos(event) {
            let pt = svg.createSVGPoint();
            pt.x = parseFloat(event.pageX);
            pt.y = parseFloat(event.pageY);
            pt.matrixTransform(svg.getScreenCTM().inverse());

            return {
                x: pt.x,
                y: pt.y
            }
        }

        function mouseDownHandler(event) {
            event.target.addEventListener('mouseleave',mouseUpHandler);
            event.target.addEventListener('mousemove',mouseMoveHandler);
            let pt = getMousePos(event);
            
            mouseOffsetX = pt.x - parseInt(event.target.getAttributeNS(null, "x"));
            mouseOffsetY = pt.y - parseInt(event.target.getAttributeNS(null, "y"));;
            //console.log("Offset x/y ", pt.x, pt.y);
        }

        function mouseUpHandler(event) {
            //console.log("mouseUP");
            event.target.removeEventListener('mousemove', mouseMoveHandler);
            event.target.removeEventListener('mouseleave',mouseUpHandler);
            mouseOffsetX = 0;
            mouseOffsetY = 0;
            
        }

        function mouseMoveHandler(event) {
            //console.log("mouseMOVE");
            const et = event.target;
            let pt = svg.createSVGPoint();
            event.preventDefault();

            pt = getMousePos(event);

            // let correctionX = pt.x-et.getAttribute("width")/2;
            // let correctionY = pt.y-et.getAttribute("height")/2;

            let correctionX = pt.x - mouseOffsetX;
            let correctionY = pt.y - mouseOffsetY;

            //console.log(pt.x, pt.y, " mit korretur: " + correctionX , correctionY, " offset:" , mouseOffsetX, mouseOffsetY);
            //et.setAttributeNS(null,"transform","translate("+parseInt(pt.x - correctionX)+ " "+parseInt(pt.y - correctionY)+")");
            et.setAttribute("x",parseInt(correctionX));
            et.setAttribute("y",parseInt(correctionY));
           
        }

        function keyDownHandler(event) {
           
           if(event.key === "Shift") {
               shiftPressed = true;
           }
        }

        function keyUpHandler(event) {
            //console.log(event.key);
            switch(event.key) {
                case "shift":
                    shiftPressed = false;
                    break;
                case "Escape":
                    removeFromSelection();
                    break;
                default:
                    break;
            }
         }

        function select(event) {
            let element = event.target; // ausgewähltes element.
            console.log("Shift:" + shiftPressed);
            // wenn nicht SHIFT gedrückt wurde werden vorhandene abgewählt.
            if(!shiftPressed) { 
               
                removeFromSelection(undefined);
            }

            if(mySelection.indexOf(element) === -1 ) { // wenn noch nicht selektiert.
                // add element to selection
                mySelection.push(element);
                addTransformationFrame(element);
            } else {
                //console.log("lösche einzelnes, da es schon selektiert ist.");
               removeFromSelection(element);

            }
            
            //updateForm(element);
            
        }

        function addTransformationFrame(element) {
            let targetMeta = {
                "type": "",
                "x": 0,
                "y":0,
                "w": 0,
                "h": 0,
                "r":0
            };
            let knobs = {};
            
            let margin = 5; // distance from Transformationframe to Element
            let knobMeta = {
                "w": 4,
                "h": 4,
                "fill": "rgb(242, 122, 255)",
                "fill-opacity": 0.5
            };

            let frameMeta = {
                "stroke-dasharray": "5,5",
                "d": ""
            }

            let frame;

            let group;
            let groupMeta = {
                id: "tranformationFrameGroup",
                stroke: "rgba(242, 122, 255, 0.3)",
                "stroke-width": 1,
                "fill": "rgb(242, 122, 255)",
                "fill-opacity": 0
            };

            targetMeta.type = element.nodeName;
            targetMeta.x = parseFloat(element.getAttributeNS(null,"x"));
            targetMeta.y = parseFloat(element.getAttributeNS(null,"y"));
            targetMeta.w = parseFloat(element.getAttributeNS(null,"width"));
            targetMeta.h = parseFloat(element.getAttributeNS(null,"height"));
            targetMeta.r = parseFloat(element.getAttributeNS(null,"r"));

            console.dir(targetMeta);

            if(targetMeta.type == "rect") {
                group = document.createElementNS("http://www.w3.org/2000/svg","g");
                group.setAttribute("id",groupMeta.id);
                group.setAttribute("stroke",groupMeta.stroke);
                group.setAttribute("stroke-width", groupMeta["stroke-width"]);
                group.setAttribute("fill", groupMeta["fill"]);
                group.setAttribute("fill-opacity", groupMeta["fill-opacity"]);
                svg.appendChild(group);


                frame = document.createElementNS("http://www.w3.org/2000/svg","path");
                frame.setAttribute("d", "M"+(targetMeta.x - margin) + " " + (targetMeta.y - margin) + " l"+ (targetMeta.w + margin + margin) + " 0 l"+ (0) + " " + (targetMeta.h + margin + margin) + " l" + (-targetMeta.w - (2*margin)) + " 0 Z")
                frame.setAttribute("stroke-dasharray", frameMeta["stroke-dasharray"]);
                group.appendChild(frame);

                knobs["top-left"] = document.createElementNS("http://www.w3.org/2000/svg","rect");
                knobs["top-left"].setAttribute("x",targetMeta.x - margin);
                knobs["top-left"].setAttribute("y",targetMeta.y - margin);
                knobs["top-left"].setAttribute("width",knobMeta.w);
                knobs["top-left"].setAttribute("height",knobMeta.h);
                knobs["top-left"].setAttribute("fill",knobMeta.fill);
                knobs["top-left"].setAttribute("fill-opacity",knobMeta["fill-opacity"]);
                knobs["top-left"].setAttribute("id","knob-top-left");
                group.appendChild(knobs["top-left"]);

                group.appendChild(element);
                group.addEventListener("mousedown",mouseDownHandler);
                group.addEventListener("mousemove",mouseMoveHandler);
            }

           


        }

        function removeFromSelection(element) {
            console.log("removeFromSelection " + element);
            if(mySelection.length > 0) {
                console.log("lösche " + mySelection.length + " objekte");
                // remove element from selection
                let removed;
                if(element === undefined) {
                        removed = mySelection.splice(0, mySelection.length) 
                } else {
                  
                    removed = mySelection.splice(mySelection.indexOf(element),1);
                }

            } else {
                // nothing to delete
                console.log("nichts zu löschen");
            }


        }


        function updateForm(element) {
            const attributes = ["id","name","x","y","fill","fill-opacity", "style"];
            const fieldContainer = document.getElementById("field-container");

            fieldContainer.childNodes.forEach((child) => {
                fieldContainer.removeChild(child);
                console.log("removing Child: "+child);
            });

            attributes.forEach((attribute, i, array) => {
                let input = document.createElement("input");
                input.setAttribute("name", attribute);
                input.setAttribute("id",attribute);
                input.placeholder = attribute;
                input.value = element.getAttribute(attribute);
                input.setAttribute("type","text");
                fieldContainer.insertBefore(input, fieldContainer.firstChild);
                //fieldContainer.appendChild(document.createElement("br"));
                
            });


           /*  const macNameOrigin =  document.getElementById("machineNameOrigin");
            const macName = document.getElementById("machineName");
            const macX = document.getElementById("machineX");
            const macY = document.getElementById("machineY");

            
            macNameOrigin.value = element.getAttribute("id");
            macName.value = element.getAttribute("id");
            macX.value = element.getAttribute("x");
            macY.value = element.getAttribute("y");
 */
        }

        updateSVG = function() {
            const macNameOrigin =  document.getElementById("id");
            const macName = document.getElementById("name");
            const macX = document.getElementById("x");
            const macY = document.getElementById("y");
            const macFill = document.getElementById("fill");
            const macFillOpacity = document.getElementById("fill-opacity");
            const macStyle = document.getElementById("style");
            const machine = document.getElementById(macNameOrigin.value);

            machine.setAttribute("x",macX.value);
            machine.setAttribute("y",macY.value);
            machine.setAttribute("id",macName.value);
            machine.setAttribute("style",macStyle.value);
            machine.setAttribute("fill",macFill.value);
            machine.setAttribute("fill-opacity", macFillOpacity.value);
  
        }

        init();
    }
    
   var editor = MySVGEditor("master");

   
   
}

function updateSVG() {
    editor.updateSVG();
}

function getRandNum(min,max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1) + min); //The maximum is inclusive and the minimum is inclusive
}