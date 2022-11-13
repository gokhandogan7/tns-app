/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("App Component", () => {
   
    it("AC1:Ensures to go to the fullarticles page and clicks highlight button, then checks if the url of page, highlight page, highlight container, text and date exist.", () => {
      cy.visit("/fullarticles");
      cy.get('[data-cy="highlight-button"]').click();
      cy.request("/highlights")
      cy.url().should("include", "/highlights"); 
      cy.get('[data-cy="highlight-page"]').should("exist");
      cy.get('[data-cy="highlight-header"]').should("exist");
      cy.get('[data-cy="highlight-container"]').should("exist");
      cy.get('[data-cy="short-text"]').should("exist");
      cy.get('[data-cy="date"]').should("exist");
    });
  
    it("AC2: Ensures to the existance of highlights page and whether there is data in it.", () => {
      cy.intercept("GET", "http://localhost:8081/highlights", {
        fixture: "highlights.json",
      }).as("getHighlights");
    });
  
    it("AC3: Ensures access to the Search box in the Highlight page and allows you to see if it found an highlight by doing a search in the search box. Then it checks the short text and date of the highlight. And then checks the go back button and clicks on the button to go back full articles page.", () => {
        cy.get('[data-cy="article-search-box"]').type("jack`s 1");
        cy.get('[data-cy="short-text"]').findByText("this is jack`s 1.contents highlight").should("exist");
        cy.get('[data-cy="date"]').findByText("2022-11-09 11:32:13").should("exist");
        cy.get('[data-cy="H-goback-button"]').should("exist");
        cy.get('[data-cy="H-goback-button"]').click();
        cy.url().should("include", "/fullarticles"); 
    });
  });
  