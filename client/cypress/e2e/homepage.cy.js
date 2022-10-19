/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("App Component", () => {
  before(() => {
    cy.visit("/");
  });

  it("AC1:Ensures to the home page and elements", () => {
    cy.get('[data-cy="home-page"]');
    cy.get(".header").should("exist");
    cy.get(".motto").should("exist");
    cy.get(".link").should("exist");
    cy.get(".link").click();
  });

  it("AC2: Ensures to the details page and articles", () => {
    cy.url().should("include", "/details");
    cy.get('[data-cy="detail-page"]');
    cy.intercept("GET", "http://localhost:8081/articles", {
      fixture: "articles.json",
    }).as("getArticles");
    cy.visit("/details");
    cy.wait('@getArticles');
  });

  it("AC3: Ensures to the details page and articles", () => {
    cy.request("/details").should("exist", ".inputBox");
    cy.get('[data-cy="search-box"]').should("have.value", "");
    cy.get('[data-cy="search-box"]').type("4");
    cy.get('[data-cy="title"]').findByText("title 4").should("exist");
    cy.get('[data-cy="description"]').findByText("description 4").should("exist");
    cy.get('[data-cy="content"]').findByText("content 4").should("exist");
  });
});
