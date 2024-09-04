window.onload = () => {
    const target = document.querySelector(".dynamic-content")

    const sampleHeading = "Tip of the day"
    const sampleText = "Brush your teeth. Brush your teeth. Brush your teeth. Brush your teeth. Brush your teeth. Brush your teeth. Brush your teeth. Brush your teeth. Brush your teeth. Brush your teeth. "

    setTimeout(()=> {
        target.innerHTML = ''
        for(i=0; i<=3; i++) {
            target.innerHTML += `<div class="section">
                    <h2>${sampleHeading}</h2>
                    <p>${sampleText}</p>
                </div>
            `
            
        }
    }, 4000)
}