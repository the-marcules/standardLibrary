class LightEditor {
  constructor(htmlSelector) {
    this.selector = htmlSelector;
    this.init();
  }

  init() {
    const editorContainer = document.querySelector(this.selector);
    if (!editorContainer) {
      console.error(
        `Could not initialize Editor. Element with selector ${this.selector} not found.`
      );
      return;
    }
    this.editorContainer = editorContainer;
    this.editorContainer.classList.add("editorContainer");
    this.editorContainer.innerHTML = ""; // Clear any existing content

    const lineNumbers = document.createElement("div");
    lineNumbers.className = "lineNumbers";

    const editor = document.createElement("div");
    editor.className = "editor";
    editor.contentEditable = true;
    editor.innerText = "{}";

    this.editorContainer.appendChild(lineNumbers);
    this.editorContainer.appendChild(editor);

    this.lineNumbersElement = lineNumbers;
    this.editor = editor;

    this.addListeners();

    console.log(`LightEditor initialized with selector: '${this.selector}'`);
  }

  addListeners() {
    this.editor.addEventListener("input", () => {
      this.updateLineNumbers();
    });

    this.editor.addEventListener("keydown", (event) => {
      if (event.key === "Tab") {
        event.preventDefault();
        const selection = document.getSelection();
        const startNode = selection.anchorNode;
        console.log('selection:', selection);
        const cursorPosition = selection.anchorOffset;
        const editorContent = this.editor.innerText;
        console.log(`Tab pressed at position ${cursorPosition}`);
        this.editor.innerText =
          editorContent.substring(0, cursorPosition) +
          "\t" +
          editorContent.substring(cursorPosition);
        // Move the cursor to the right after inserting the tab
       
      }
    });

    this.editor.addEventListener("scroll", () => {
      this.lineNumbersElement.scrollTop = this.editor.scrollTop;
    });
  }

  updateLineNumbers() {
    const lines = this.getLines();
    console.log(lines);
    this.lineNumbersElement.innerHTML = "";
    for (let i = 1; i <= lines.length; i++) {
      const lineNumber = document.createElement("div");
      lineNumber.className = "lineNumber";
      lineNumber.innerText = i;
      lineNumber.style.fontFamily = getComputedStyle(this.editor).fontFamily;
      lineNumber.style.fontSize = getComputedStyle(this.editor).fontSize;

      this.lineNumbersElement.appendChild(lineNumber);
    }
  }

  getLines() {
    return this.editor.innerText.trimEnd().split("\n");
  }
}
