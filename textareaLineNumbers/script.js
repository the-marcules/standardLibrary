document.addEventListener("DOMContentLoaded", () => {
  const textarea = document.querySelector("textarea");
  const lineNumbersEle = document.querySelector(".lineNumbers");

  const textareaStyles = window.getComputedStyle(textarea);
  [
    "fontFamily",
    "fontSize",
    "fontWeight",
    "letterSpacing",
    "lineHeight",
    "padding",
  ].forEach((property) => {
    lineNumbersEle.style[property] = textareaStyles[property];
  });

  const parseValue = (v) =>
    v.endsWith("px") ? parseInt(v.slice(0, -2), 10) : 0;

  const font = `${textareaStyles.fontSize} ${textareaStyles.fontFamily}`;
  const paddingLeft = parseValue(textareaStyles.paddingLeft);
  const paddingRight = parseValue(textareaStyles.paddingRight);

  const canvas = document.createElement("canvas");
  const context = canvas.getContext("2d");
  context.font = font;

  const calculateNumLines = (str) => {
    const textareaWidth =
      textarea.getBoundingClientRect().width - paddingLeft - paddingRight;
    const words = str.split(" ");
    let lineCount = 0;
    let currentLine = "";
    for (let i = 0; i < words.length; i++) {
      const wordWidth = context.measureText(words[i] + " ").width;
      const lineWidth = context.measureText(currentLine).width;

      if (lineWidth + wordWidth > textareaWidth) {
        lineCount++;
        currentLine = words[i] + " ";
      } else {
        currentLine += words[i] + " ";
      }
    }

    if (currentLine.trim() !== "") {
      lineCount++;
    }

    return lineCount;
  };

  const calculateLineNumbers = () => {
    const lines = textarea.value.split("\n");
    const numLines = lines.map((line) => calculateNumLines(line));

    let lineNumbers = [];
    let i = 1;
    while (numLines.length > 0) {
      const numLinesOfSentence = numLines.shift();
      lineNumbers.push(i);
      if (numLinesOfSentence > 1) {
        Array(numLinesOfSentence - 1)
          .fill("")
          .forEach((_) => lineNumbers.push(""));
      }
      i++;
    }

    return lineNumbers;
  };

  const displayLineNumbers = () => {
    const lineNumbers = calculateLineNumbers();
    lineNumbersEle.innerHTML = Array.from(
      {
        length: lineNumbers.length,
      },
      (_, i) => `<div>${lineNumbers[i] || "&nbsp;"}</div>`
    ).join("");
  };

  const validateJson = () => {
    const notificationElement = document.querySelector(".inlineNotification");
    try {
      JSON.parse(textarea.value);
      notificationElement.textContent = "";
    } catch (e) {
      notificationElement.textContent = e.message;

    }
  };

  textarea.addEventListener("input", () => {
    validateJson();
    displayLineNumbers();
  });

  displayLineNumbers();

  const ro = new ResizeObserver(() => {
    const rect = textarea.getBoundingClientRect();
    lineNumbersEle.style.height = `${rect.height}px`;
    displayLineNumbers();
  });
  ro.observe(textarea);

  textarea.addEventListener("scroll", () => {
    lineNumbersEle.scrollTop = textarea.scrollTop;
  });

    textarea.addEventListener("keyup", () => {

    const positionStart = textarea.selectionStart;



    const value = textarea.value;
    const lines = value.split("\n");

    let positionsToSubtract = 0;
    let currentLine = 1
    for(i = 0; i < lines.length; i++) {
      let lineLength = lines[i].length + 1; // +1 for the newline character
      if ((lineLength + positionsToSubtract) > positionStart) {
        break;
      } else {
        positionsToSubtract += lineLength;
        currentLine = i + 2; // +1 to convert to 1-based index
      }
    }

    const correctedPosition = (positionStart - positionsToSubtract) +1 

    const cursorPositionElement = document.querySelector(".cursorPosition");
    cursorPositionElement.textContent = `Line: ${currentLine}, Column: ${correctedPosition}`;

  });

  textarea.addEventListener('keydown', function(e) {
  if (e.key == 'Tab') {
    e.preventDefault();
    var start = this.selectionStart;
    var end = this.selectionEnd;

    // set textarea value to: text before caret + tab + text after caret
    this.value = this.value.substring(0, start) +
      "\t" + this.value.substring(end);

    // put caret at right position again
    this.selectionStart = this.selectionEnd = start + 1;
  }
});


});
