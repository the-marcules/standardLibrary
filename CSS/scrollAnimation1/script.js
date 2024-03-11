
window.addEventListener('scroll', ()=> {
    const logo = document.getElementById('logo')
    const setRotation = ` rotate(-${window.scrollY}deg)`
    console.log("setRotation", setRotation)
    logo.style.transform = setRotation

})