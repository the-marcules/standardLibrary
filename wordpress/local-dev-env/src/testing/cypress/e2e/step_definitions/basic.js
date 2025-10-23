import { Given, Then, When } from "@badeball/cypress-cucumber-preprocessor";
import "cypress-iframe";

Given("I open the homepage", () => {
  cy.visit("/");
});

Then("I see the top navigation menu", () => {
  cy.get("#menu-haupt")
    .should("be.visible")
    .children()
    .should("have.length", 4);
  cy.get("a.nav-link-container-warenkorb").should("exist");
});

Then("I see the logo image", () => {
  cy.get("div.logo img").should("be.visible");
});

Then("I see {string}", (text) => {
  cy.contains(text);
});

Then("I see the footer navigation menu", () => {
  cy.get("footer #menu-footer")
    .should("be.visible")
    .children()
    .should("have.length", 3);
});

When("I click on the {string} link in the top navigation menu", (linkText) => {
  cy.get("#menu-haupt").contains(linkText).click();
});

Then("I see {string} in the {string} content area", (text, region) => {
  cy.get(region).contains(text);
});

Then("I do not see {string} in the {string} content area", (text, region) => {
  if(text) {
    cy.get(region).contains(text).should("not.exist");
  }
});

Then("I see the product {string}", (productName) => {
  cy.get(".wopro-products-container").contains(productName);
});

When("I click on the cart icon", () => {
  cy.get("a.nav-link-container-warenkorb").should("exist").click();
});

When("I click on the {string} button", (buttonText) => {
  cy.contains("button", buttonText).click();
});

When("I click on the {string} link", (linkText) => {
  cy.contains("a", linkText).click({ force: true });
});

When(
  "I click on the {string} link within the element {string}",
  (linkText, containerSelector) => {
    cy.get(containerSelector).contains("a", linkText).click({ force: true });
  }
);

Then("I see the element {string}", (selector) => {
  cy.get(selector).should("exist").and("be.visible");
});

Then("I see that the element {string} is gone", (selector) => {
  cy.get(selector, { timeout: 6000 }).should("not.be.visible");
});

When("I check the checkbox with id {string}", (id) => {
  cy.get(`input[type=checkbox]${id}`)
    .should("exist")
    .check({ force: true })
    .should("be.checked");
});

Then("I see no error or warning within the element {string}", (selector) => {
  cy.get(selector)
    .should("not.contain.text", "Error")
    .should("not.contain.text", "Warning")
    .should("not.contain.text", "Notice")
    .should("not.contain.text", "Fatal");
});
