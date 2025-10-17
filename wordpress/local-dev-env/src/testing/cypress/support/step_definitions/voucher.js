import { Given, When, Then } from "@badeball/cypress-cucumber-preprocessor";

Then("I see the input field with id {string} to be {string}", (id, state) => {
  if (state == "empty") cy.get(`input#${id}`).should("be.empty");
  else cy.get(`input#${id}`).invoke("val").should("not.be.empty");
});

When("I set the valid until date to {string} by clicking", (date) => {
  const validUntil = new Date(date == 'tomorrow' ? Date.now() : date);
  validUntil.setDate(validUntil.getDate() + 1);
  cy.get('input[name="valid_until"]').type(validUntil.toISOString().slice(0, 10));
});

When("I set the voucher type to {string}", (type) => {
  cy.get(`input[value="${type}"]`).check();
});


When("I set the voucher code to {string}", (code) => {
  cy.get(`input#code`).type(code)
}); 

When("I set the voucher value to {string}", (value) => {
  cy.get('input[name="value"]').clear().type(value);
});

When("I submit the voucher form", () => {
  cy.get('input[type="submit"]').click();
});

Then("I see the success message that the voucher was created", () => {
  cy.get(".notice-success").should(
    "contain.text",
    "Gutschein erstellt."
  );
});

Then("I see the new voucher in the voucher list", () => {
  cy.get("table").find("tr").should("have.length.greaterThan", 1);
}); 

Then("I see the new voucher with code {string} in the voucher list", (voucher_code) => {
  cy.get("table").find("tr").contains(voucher_code).should("be.visible");
}); 

When("I delete the created voucher", () => {
  cy.get("a").contains("🗑️")
    .click();
});

When("I delete the voucher with code {string}", (voucher_code) => {
 cy.get("tr").contains(voucher_code).parent().find("a").contains("🗑️")
    .click({force: true});
});

Then("I see the success message that the voucher was deleted", () => {
  cy.get(".notice-success").should(
    "contain.text",
    "Gutschein gelöscht."
  );
});