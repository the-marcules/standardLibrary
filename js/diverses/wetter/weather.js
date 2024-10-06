(function loadXMLDoc() {
  var xhttp = new XMLHttpRequest();
  xhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) {
      document.getElementById("demo").innerHTML =
      console.log(this.responseText);
    }
  };
  xhttp.open("GET", "https://api.openweathermap.org/data/2.5/weather?q=Ingolstadt,de&appid=950fb8be66fdae5f7c9eb58cba59d8b8&units=metric&lang=de", true);
  xhttp.send();
})