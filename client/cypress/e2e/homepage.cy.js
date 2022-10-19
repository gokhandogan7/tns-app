/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("App Component", () => {
  before(() => {
    cy.visit("/");
  });

  it("goes to the home page", () => {
    cy.get(".header").should("exist");
    cy.get(".motto").should("exist");
    cy.get(".link").should("exist");
    cy.get(".link").click();
    cy.url().should("include", "/details");
    cy.request("/details").should("exist", ".inputBox");
    cy.get(".inputBox").should("have.value", "");
    cy.get(".inputBox").type("4");
    cy.get(".title").findByText("title 4").should("exist")
    cy.get(".description").findByText("description 4").should("exist")
    cy.get(".content").findByText("content 4").should("exist")
    
  });
});
