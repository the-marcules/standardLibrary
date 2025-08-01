"use client";
import React, { useEffect } from "react";
import style from "./documentation.module.css";

const DocumentationPage = () => {
  useEffect(() => {
    const targetElementId = window.location.hash;
    if (targetElementId) {
      console.log("targetElementId:", targetElementId);
      const targetElement = document.querySelector(targetElementId);
      if (targetElement) {
        const header = document.querySelector("header");
        const headerHeight = header ? header.getBoundingClientRect().height : 0;
        console.log("Header height:", headerHeight);

        // "use client" sorgt nur dafür, dass der Code im Browser ausgeführt wird,
        // aber offsetTop ist nur für HTMLElement verfügbar, nicht für ein generisches Element.
        // TypeScript weiß nicht, ob targetElement ein HTMLElement ist.
        // Daher musst du es explizit casten:
        const element = targetElement as HTMLElement;
        console.log("Element offsetTop:", element.offsetTop);
        window.scrollTo({
          top: element.offsetTop + headerHeight,
          behavior: "smooth",
        });
      }
    }
  }, []);

  return (
    <div className={style.documentationContainer}>
      <h1>Documentation</h1>
      <div>
        Welcome to the documentation! Here you will find guides, API references,
        and examples to help you get started and make the most of this project.
      </div>
      <section id="introduction">
        <h2>Introduction</h2>
        <div>
          This is the introduction section. Here you will learn why this
          documentation exists and how it can help you become a documentation
          wizard. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
          Bananas are not required.
        </div>
      </section>
      <section id="usage">
        <h2>Usage</h2>
        <div>
          To use this project, simply press the big red button and hope for the
          best. If nothing happens, try turning your computer off and on again.
          Remember, usage may vary depending on the phase of the moon.
        </div>
      </section>
      <section id="examples">
        <h2>Examples</h2>
        <div>
          Here are some examples:
          <ul>
            <li>Example 1: Do something amazing with just one click.</li>
            <li>
              Example 2: Watch as nothing happens, but in a very impressive way.
            </li>
            <li>
              Example 3: Combine two features and get a third, completely
              unrelated result.
            </li>
          </ul>
        </div>
      </section>
      <section id="api">
        <h2>API</h2>
        <div>
          The API is so simple, even a rubber duck could use it. Just call{" "}
          <code>doMagic()</code> and wait for the fireworks. If you encounter
          any errors, blame the duck.
        </div>
      </section>
    </div>
  );
};

export default DocumentationPage;
