/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("App Component", () => {
  beforeEach(() => {
    cy.visit("http://localhost:3000/");
  });

  it("goes to the home page", () => {
    cy.get(".header").should("exist");
    cy.get(".motto").should("exist");
  });

  it("routes to details page", () => {
    cy.get(".link").click();
    cy.url().should("include", "/details");
    cy.get(".inputBox").should("have.value", '')
  });
});
