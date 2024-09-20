class ShittyElement extends HTMLElement {
    #ComponentName = 'ShittyComponent'
    constructor() {
        super()
        const shadow = this.attachShadow({ mode: 'open' })
        const wrapper = document.createElement('div')
        wrapper.setAttribute('id', 'wrapper')
        
        const title = document.createElement('h2')
        title.setAttribute('id','title')

        const description = document.createElement('div')
        description.setAttribute('id', 'description')

        wrapper.appendChild(title)
        wrapper.appendChild(description)

        // Lade die externe CSS-Datei
        const linkElement = document.createElement('link');
        linkElement.setAttribute('rel', 'stylesheet');
        linkElement.setAttribute('href', `./${this.#ComponentName}/style.css`); // Der Pfad zu deiner CSS-Datei


        shadow.appendChild(linkElement);
        shadow.appendChild(wrapper);

    }
    
    connectedCallback() {
        this.updateContent()
    }

    static get observedAttributes() {
        return ['title', 'description'];
    }

    attributeChangedCallback(name, oldValue, newValue) {
        if (oldValue !== newValue) {
            console.log(`Attribut ${name} geändert von ${oldValue} zu ${newValue}`);
            this.updateContent();
        }
    }


    updateContent() {
        this.shadowRoot.querySelector('h2#title').textContent = `💩 ${this.getAttribute('title')}`
        this.shadowRoot.querySelector('div#description').textContent = `${this.getAttribute('description')}`
    }
}

customElements.define('shitty-element', ShittyElement)