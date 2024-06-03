"use strict";
function showDialog() {
    const dialog = document.querySelector("dialog");
    dialog === null || dialog === void 0 ? void 0 : dialog.showModal();
}
function hideDialog() {
    const dialog = document.querySelector("dialog");
    dialog === null || dialog === void 0 ? void 0 : dialog.close();
}
