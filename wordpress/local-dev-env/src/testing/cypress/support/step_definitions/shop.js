import { Given, When, Then } from "@badeball/cypress-cucumber-preprocessor";

Then("I see {string} in the shipping cost summary", (text) => {
  cy.get("table.wopro-sum-table")
    .find("tr")
    .eq(2)
    .find("td")
    .eq(1)
    .contains(text);
});

Then("I see {string} in the cart summary", (text) => {
  cy.get("div.wopro-cart-content").contains(text);
});

Then("I see {string} in the cart icon", (number) => {
  cy.get("a.nav-link-container-warenkorb")
    .should("exist")
    .contains(number)
    .should("be.visible");
});

Then("I see the product details for {string}", (productName) => {
  cy.contains("button", "zurück zur Übersicht");
  cy.get("article").contains(productName);
  cy.get("div.wopro-info")
    .should("be.visible")
    .should("contain", "Details")
    .should("contain", "Bestellen")
    .should("contain", "Preis")
    .should("contain", "Farbe")
    .should("contain", "Menge")
    .should("contain", "Aktueller Lagerbestand:")
    .and("contain", "in den Warenkorb");
  cy.get("input#prod-amount").should("exist").should("have.value", "1");
});
