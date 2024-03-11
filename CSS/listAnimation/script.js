
function init() {
  generateItems(7)
}

function generateItems(count) {
    const list = document.getElementById('list')
    if (list) {
        for (i=1; i<count+1; i++) {
            const newItemName = `New Item #${i}`
            const newItem = document.createElement('div')
            newItem.classList.add('listItem') 

            newItem.onclick = (e) => { 
                
                const startHeight = window.getComputedStyle(e.currentTarget).height
                e.currentTarget.style.height = startHeight
                e.currentTarget.classList.add('fadeOut') 
            
            }
            
            const subtitle = (i%2?'subtitle':'subtitle line1<br>subtitle line2')

            newItem.innerHTML = `
                <div class='icon'>I</div>
                <div class='itemHeading'>${newItemName}</div>
                <div class='cta'>X</div>
                <div class='subtitle'>${subtitle}</div>`
            
            list.append(newItem)
        }

    }
}
