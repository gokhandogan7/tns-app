/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("App Component", () => {
  before(() => {
    cy.visit("/");
  });

  it("AC1:Ensures to go to the main page and checks the header, motto and the link of the detail page, then clicks on the link and checks the url of the detail page.", () => {
    cy.get('[data-cy="home-page"]');
    cy.get('[data-cy="header"]').should("exist");
    cy.get('[data-cy="motto"]').should("exist");
    cy.get('[data-cy="link"]').should("exist");
    cy.get('[data-cy="link"]').click();
    cy.url().should("include", "/details");
  });

  it("AC2: Ensures to the existance of detail page and whether there is data in it.", () => {
    cy.intercept("GET", "http://localhost:8081/articles", {
      fixture: "articles.json",
    }).as("getArticles");
    cy.get('[data-cy="detail-page"]').should("exist");
  });

  it("AC3: Ensures access to the Search box in the Details page and allows you to see if it found an article by doing a search in the search box. Then it checks the title, description and content of the article.", () => {
    cy.request("/details").should("exist", ".inputBox");
    cy.get('[data-cy="search-box"]').should("have.value", "");
    cy.get('[data-cy="search-box"]').type("4");
    cy.get('[data-cy="title"]').findByText("title 4").should("exist");
    cy.get('[data-cy="description"]')
      .findByText("description 4")
      .should("exist");
    cy.get('[data-cy="content"]').findByText("content 4").should("exist");
  });
});
