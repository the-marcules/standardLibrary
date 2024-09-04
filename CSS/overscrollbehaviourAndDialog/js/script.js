"use strict";
function showDialog() {
    const dialog = document.querySelector("dialog");
    if (dialog) {
        dialog.showModal();
        document.addEventListener('click', clickedOutsideHandler)
    }
    const body = document.querySelector("body")
    body.style.overflow = 'hidden';

}
function hideDialog() {
    const dialog = document.querySelector("dialog");
    dialog === null || dialog === void 0 ? void 0 : dialog.close();
    const body = document.querySelector("body")
    body.style.overflow = 'auto';

}


function clickedOutsideHandler (event) {
    console.table(event)
    document.addEventListener('click', clickedOutsideHandler)
}