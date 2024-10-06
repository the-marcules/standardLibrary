// classes:


class WeatherApiCall extends XMLHttpRequest{
   constructor(p_url, parameters) { // url: http://www.thisOrThat.com/, parameters: [{name: '', value: ''}, {name: '', value: ''}, ...]
        super(); // needs to be called to initiate parent class
        this.url = p_url;
        this.params = parameters;
        self = this; // workaround for the override of this in a promise construct
    }

    //function to get weather
    async get() {
        return this.getHandler().then(function(response){
            return response;
        });
    }
    // encapsulation needed for the promise then() construct
    async getHandler() {
        this.open("GET", this.url+this.buildParamStr(), true);
        this.send();
        
        //INFO: great example: view-source:https://mdn.github.io/js-examples/promises-test/
        return new Promise(function (resolve,reject){
           self.onreadystatechange = function() {
                if (self.readyState == 4 && self.status == 200) {
                    resolve(self.responseText);
                }
            };
        });
        
    }   

    buildParamStr() {
        let paramStr = "";
        this.params.forEach((par, num, arr) => {
            
            if(num === 0) paramStr += "?"+par.name+"="+par.value
            else paramStr += "&"+par.name+"="+par.value
        });
        return paramStr;
    }
   
    
}

class WeatherRender {
    constructor(data, target){
        this.weatherData = JSON.parse(data);
        this.domTarget = document.getElementById(target);
        this.displayItems = Array("icon","wetter-details","wind-icon");
        this.subDisplayItems = Array("temp", "weatherData-text","wind");
        //console.log("typ:" + this.weatherData);
        this.myCanvases = Array();
        this.windBedingungen = [
            {
            speedVon: 0,
            speedBis: 0.3,
            description: "Windstille",
            color_1: "green",
            color_2: "lime"
        },
        {
            speedVon: 0.3,
            speedBis: 1.6,
            description: "leiser Zug",
            color_1: "green",
            color_2: "lime" 
        },
        {
            speedVon: 1.6,
            speedBis: 3.4,
            description: "leichte Brise",
            color_1: "green",
            color_2: "lime"  
        },
        {
            speedVon: 3.4,
            speedBis: 5.5,
            description: "schwache Brise",
            color_1: "green",
            color_2: "lime"  
        },
        {
            speedVon: 5.5,
            speedBis: 8.0,
            description: "mäßige Brise",
            color_1: "green",
            color_2: "lime"  
        },
        {
            speedVon: 8.0,
            speedBis: 10.8,
            description: "frische Brise",
            color_1: "yellow",
            color_2: "orange"  
        },
        {
            speedVon: 10.8,
            speedBis: 13.9,
            description: "starker Wind",
            color_1: "yellow",
            color_2: "orange" 
        },
        {
            speedVon: 13.9,
            speedBis: 17.2,
            description: "steifer Wind",
            color_1: "yellow",
            color_2: "orange" 
        },
        {
            speedVon: 17.2,
            speedBis: 20.8,
            description: "stürmischer Wind",
            color_1: "orange",
            color_2: "OrangeRed" 
        },
        {
            speedVon: 20.8,
            speedBis: 24.5,
            description: "Sturm",
            color_1: "orange",
            color_2: "OrangeRed"  
        },
        {
            speedVon: 24.5,
            speedBis: 28.5,
            description: "schwerer Sturm",
            color_1: "red",
            color_2: "DarkRed"  
        },
        {
            speedVon: 28.5,
            speedBis: 32.7,
            description: "orkanartiger Sturm",
            color_1: "red",
            color_2: "DarkRed" 
        },
        {
            speedVon: 32.7,
            speedBis: 1000000,
            description: "Orkan",
            color_1: "Purple",
            color_2: "Pink" 
        },

    ];

    }

    // render currrent weather
    currentWeather() {

        const cwHeading = document.createElement("h2");
        cwHeading.innerHTML = "Current Weather";
        this.domTarget.appendChild(cwHeading);

        const newId = "currentWeather";
        const newDiv = document.createElement("div");
        newDiv.setAttribute("id",newId);

        this.domTarget.appendChild(newDiv);

        this.displayItems.forEach((curr, index, arr)=>{
            let element = document.createElement("div");
            element.setAttribute("id",curr);
            newDiv.appendChild(element);
        });

        const subTarget = document.getElementById("wetter-details");
        this.subDisplayItems.forEach((item) => {
            let element = document.createElement("p");
            element.setAttribute("id",item);
            subTarget.appendChild(element);

        });
        document.getElementById("icon").innerHTML = "<img src='http://openweathermap.org/img/wn/"+this.weatherData.current.weather[0].icon+"@2x.png'>";
                    
        document.getElementById("temp").innerHTML = "<b style='font-size: 1.4em'>"+this.formatTempOutput(this.weatherData.current.temp) +"</b> - gefühlt " + this.formatTempOutput(this.weatherData.current.feels_like) ; 
        document.getElementById("weatherData-text").innerHTML = "<b>"+this.weatherData.current.weather[0].description + "</b> - Sonnenaufgang " + this.formatTime(this.weatherData.current.sunrise) + " Sonnenuntergang " + this.formatTime(this.weatherData.current.sunset) ;
        
        const wind = this.matchWindBedingungen(this.weatherData.current.wind_speed);
        document.getElementById("wind").innerHTML = wind.description+" mit <b>"+ this.weatherData.current.wind_speed + "m/s</b> ";
        document.getElementById("wind-icon").innerHTML ="<canvas id='canvasWind' width='90' height='90'></canvas>";
        
        this.windRichtung({...wind,"cid":"canvasWind", "deg":this.weatherData.current.wind_deg}); // Windrose für das aktuelle Wetter 

    }
    //  build all forecast boxes
    forecastWeather() {
        const fcHeading = document.createElement("h2");
        fcHeading.innerHTML = "Forecast Weather";
        fcHeading.style.clear = "left";
        this.domTarget.appendChild(fcHeading);

        const newElement = document.createElement("div");
        newElement.setAttribute("id","forecast");
        newElement.setAttribute("class","grid-container");
        this.domTarget.appendChild(newElement);

        this.weatherData.daily.forEach((curr, index, arr)=>{
            let col = index +1 ;
            let newWrapper = document.createElement("div");
            newWrapper.setAttribute("id","fc_"+col);
            if(col < arr.length) newWrapper.setAttribute("class","column_"+col+" grid-item");
            else newWrapper.setAttribute("class","column_"+col+" grid-item");

            document.getElementById("forecast").appendChild(newWrapper);
            document.getElementById("fc_"+col).innerHTML += "<div class='fc-date'><b>"+this.formatDate(curr.dt) +"</b></div>"; //Datum
            document.getElementById("fc_"+col).innerHTML += "<div><img src='http://openweathermap.org/img/wn/"+curr.weather[0].icon+"@2x.png'></div>";
            document.getElementById("fc_"+col).innerHTML += "<div><b style='font-size: 1.5em; color: #333'>"+this.formatTempOutput(curr.temp.day) +"</b> / "+this.formatTempOutput(curr.temp.min) +"</div>";
            document.getElementById("fc_"+col).innerHTML += "<div>Gefühlt: "+this.formatTempOutput(curr.feels_like.day) +"</div>";
            document.getElementById("fc_"+col).innerHTML += "<div><b>"+curr.weather[0].description + "</b></div>";
            const wind = this.matchWindBedingungen(curr.wind_speed);
            document.getElementById("fc_"+col).innerHTML += "<div>"+wind.description+" mit <b>"+curr.wind_speed + "m/s</b> </div>";
            document.getElementById("fc_"+col).innerHTML += "<div id='canvasContainer_"+col+"' class='canvas-container'><canvas id='canvas_"+col+"' width='80' height='80'></canvas></div>";
            this.myCanvases.push({...wind, "cid":"canvas_"+col, "deg": curr.wind_deg}); // {"cid":"canvas_"+col, "deg": curr.wind_deg, "speed": curr.wind_speed}
        });
       
        this.myCanvases.forEach(this.windRichtung);
    }
    
    formatTempOutput(temp) {
        //console.log("formatTempOutput:"+temp)
        return temp.toString().split(".")[0] + "°C";
    }

    formatTime(timestamp){  
        const date = new Date(timestamp * 1000);
        let hours = date.getHours();
        let minutes = "0" + date.getMinutes();
        
        return hours + ':' + minutes.substr(-2) + " Uhr"; 
    }
    
    matchWindBedingungen(speed) {
        let hit 
        this.windBedingungen.forEach((current, index, arr)=>{
            
            if(speed>= current.speedVon && speed < current.speedBis) {
                
                hit = {
                    speed,
                    ...current
                };
            }
            
        });

        return hit;
    }

    formatDate(timestamp){
        const weekday = Array("Sonntag","Montag","Dienstag","Mittwoch","Donnerstag","Freitag","Samstag");  
        const date = new Date(timestamp * 1000);
        let year = date.getFullYear();
        let month = date.getMonth()+1;
        let day = date.getDate();
        let dayOfWeek = date.getDay();
        
        return weekday[dayOfWeek].slice(0,2) + ", " + day + "." + month + "." + year; 
    }


    windRichtung(params) {
        const canvasId = params.cid;
        const angle = params.deg;
        const speed = params.speed;
        const color_1 = params.color_1;
        const color_2 = params.color_2;
        const description = params.description;

        //const wind = matchWindBedingungen(speed);
        
        if(!canvasId || !angle) return false;

        let canvas = document.getElementById(canvasId);
        
        if(!canvas) {
            
            return false;
        }
        try {
            let ctx = canvas.getContext("2d");
                    const padding = 0.75;
                    const rotationAngle = (180+angle)*Math.PI/180; // Berechung des Winkels (invertiert auf Grund der Pfeilausrichtung) und Konvertierung in RAD
                    const radius = 0.99 * canvas.width/2 ;
                    const directions = Array("N","NO","O","SO","S","SW","W","NW");
                    
                    // Kreishintergrund
                    ctx.beginPath();
                    ctx.fillStyle = '#FFF';
                    ctx.globalAlpha = 0.3;
                    ctx.strokeStyle = '#EFEFEF';
                    ctx.arc(canvas.width/2, canvas.height/2, radius*0.75 , 0, 2 * Math.PI);
                    ctx.fill();
                    ctx.stroke();
                    ctx.closePath();
                    ctx.globalAlpha = 1;
                    //Himmelsrichtungen schreiben
                    
                    ctx.strokeStyle = '';
                    
                    ctx.textBaseline="middle"
                    ctx.textAlign="center";
                    ctx.translate(canvas.width/2, canvas.height/2);
                    directions.forEach((current, num, arr)=>{
                        let ang = num * Math.PI / 4;
                        ctx.rotate(ang);
                        // pfeil in richtung
                        ctx.beginPath();
                        ctx.fillStyle = '#EFEFEF';
                        ctx.moveTo(canvas.width/12 , 0);
                        ctx.lineTo(0, ( (num+1)%2? -canvas.height/2*padding:-canvas.height*(1-padding) ));
                        ctx.lineTo(-canvas.width/12 , 0);
                        ctx.closePath();
                        ctx.fill();

                        //schreiben
                        ctx.translate(0, -radius*0.9);
                        ctx.rotate(-ang);
                        ctx.fillStyle = '#333';
                        ctx.font = ((num+1)%2?"bold italic":" italic")+" "+radius*0.20 + "px 'calibri'";
                        ctx.fillText(current, 0, 0);
                        ctx.rotate(ang);
                        ctx.translate(0, radius*0.9);
                        ctx.rotate(-ang);
                        });
                    ctx.translate(-canvas.width/2, -canvas.height/2);
                    



                    
                    //Windrichtung
                    //rotation
                    ctx.translate(canvas.width/2,canvas.height/2);
                    ctx.rotate(rotationAngle);
                    ctx.translate(-canvas.width/2,-canvas.height/2);
                    
                    ctx.beginPath();
                    ctx.moveTo(canvas.width/2 + canvas.width/6, canvas.height* padding);
                    ctx.lineTo(canvas.width/2, canvas.height*(1-padding));
                    ctx.lineTo(canvas.width/2 - canvas.width/6, canvas.height*padding);
                    ctx.lineTo(canvas.width/2,canvas.height*(padding-0.15));
                    ctx.closePath();

                    //stroke
                    ctx.lineWidth = 2;
                    ctx.strokeStyle = '#111111';
                    ctx.stroke();

                    // farbe
                    let grd = ctx.createLinearGradient(0, 0, canvas.width, canvas.height);
                    grd.addColorStop(0, color_1);
                    grd.addColorStop(1, color_2);

                    ctx.fillStyle = grd;
                    ctx.fill();
                
                    ctx.rotate(-rotationAngle);

        } catch (e) {
            console.log(JSON.stringify(e));
        }
    }




}


