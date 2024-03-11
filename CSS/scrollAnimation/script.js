let timeout 

window.addEventListener('scroll', ()=> {
    const cube = document.getElementById('cube')
    cube.classList.remove('rotation')
    if (timeout) {
        clearTimeout(timeout)
    }
    const setRotation = ` rotateX(${window.scrollY/10}deg) rotateY(${window.scrollY/10}deg)`
    console.log("setRotation", setRotation)
    cube.style.transform = setRotation

    timeout =  setTimeout(()=> {
        cube.classList.add('rotation')
    }, 1500)

})