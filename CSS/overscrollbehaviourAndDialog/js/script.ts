function showDialog():void {
    const dialog = document.querySelector("dialog")

    dialog?.showModal()
    
}

function hideDialog():void {
    const dialog = document.querySelector("dialog")

    dialog?.close()
    
}