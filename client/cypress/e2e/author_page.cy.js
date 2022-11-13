/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("App Component", () => {
   
  
    it("AC1:Ensures to go to the fullarticles page and clicks author button, then checks if the url of page, author page, header, author container, name and email exist.", () => {

      cy.visit("/fullarticles");
      cy.get('[data-cy="author-button"]').click();
      cy.request("/authors")
      cy.url().should("include", "/authors"); 
      cy.get('[data-cy="author-page"]').should("exist");
      cy.get('[data-cy="author-header"]').should("exist");
      cy.get('[data-cy="author-container"]').should("exist");
      cy.get('[data-cy="name"]').should("exist");
      cy.get('[data-cy="email"]').should("exist");
    });
  
    it("AC2: Ensures to the existance of author page and whether there is data in it.", () => {
      cy.intercept("GET", "http://localhost:8081/authors", {
        fixture: "authors.json",
      }).as("getAuthors");
    });
  
    it("AC3: Ensures access to the authors in the AuthorPage and allows you to see if it found an author by doing a search in the search box. Then it checks the name and email of the author. And then checks the go back button and click on the button to go back full articles page.", () => {
        cy.get('[data-cy="article-search-box"]').type("mack");
        cy.get('[data-cy="name"]').findByText("mack").should("exist");
        cy.get('[data-cy="email"]').findByText("mack@mack.com").should("exist");
        cy.get('[data-cy="A-goback-button"]').should("exist");
        cy.get('[data-cy="A-goback-button"]').click();
        cy.url().should("include", "/fullarticles"); 
    });
  });
  