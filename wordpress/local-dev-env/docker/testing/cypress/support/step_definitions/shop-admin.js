import { Given, When, Then } from "@badeball/cypress-cucumber-preprocessor";

Given("I log in to the wordpress admin area", () => {
  cy.visit("/wp-admin");
  cy.get("input#user_login").type("admin");
  cy.get("input#user_pass").type("admin");
  cy.get("input#wp-submit").click();
  cy.url().should("include", "/wp-admin/");
});

Given(
  "I am on the {string} admin page with title {string}",
  (page, pageTitle) => {
    const pageSlug = page.toLowerCase();
    cy.visit(`wp-admin/edit.php?post_type=${pageSlug}`);
    cy.get("h1.wp-heading-inline").should("contain", pageTitle);
  }
);

Then("I see {string} in the products list", (text) => {
  cy.get("table.wp-list-table").contains(text);
});

When(
  "I create a new product with the title {string}, price {string}, stock {string} and details {string}",
  (name, price, stock, details) => {
    cy.get("input#title").type(name);
    cy.get("input[name='price']").clear().type(price);
    cy.get("input[name*='stock[']").clear().type(stock);
    cy.get("iframe#details_ifr").then(($iframe) => {
      const $body = $iframe.contents().find("body");
      cy.wrap($body).clear().type(details);
    });
    cy.get("input[type='submit']#publish").click({ force: true });
  }
);

When("I select all products except {string}", (productName) => {
  cy.get("tbody#the-list")
    .find("tr")
    .each(($el) => {
      const regex = /Testprodukt \d/;
      if (regex.test($el.text())) {
        cy.log("found match: " + $el.text());
        cy.wrap($el).find("input[type='checkbox']").check();
      }
    });
});

When("I select {string} in the bulk actions dropdown", (action) => {
  cy.get("select#bulk-action-selector-top").select(action);
  cy.get("input#doaction").click();
});

Then("I see one product in the products list", () => {
  cy.get("tbody#the-list").children().should("have.length", 1);
});
