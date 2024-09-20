window.onload = () => {
    console.log('loaded')
    const firstShitty = document.querySelector('shitty-element')
    const attributeNames = firstShitty.getAttributeNames()
    console.log(attributeNames)
    firstShitty.classList.add('test')
}
