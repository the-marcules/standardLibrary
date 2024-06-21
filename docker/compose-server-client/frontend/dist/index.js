fetch('http://localhost:8080')
    .then(result => result.text())
    .then(data => console.log(data))
    .catch(err => console.error(err));
//# sourceMappingURL=index.js.map