/* eslint-disable no-undef */
/// <reference types="Cypress"/>
describe("App Component", () => {
  
    it("AC1:Ensures to go to the fullarticles page and clicks content button, then checks if the url of page, content page, header, content container, text and image exist.", () => {

      cy.visit("/fullarticles");
      cy.get('[data-cy="content-button"]').click();
      cy.request("/contents")
      cy.url().should("include", "/contents"); 
      cy.get('[data-cy="content-page"]').should("exist");
      cy.get('[data-cy="content-header"]').should("exist");
      cy.get('[data-cy="content-container"]').should("exist");
      cy.get('[data-cy="text"]').should("exist");
      cy.get('[data-cy="image"]').should("exist");
    });
  
    it("AC2: Ensures to the existance of content page and whether there is data in it.", () => {
      cy.intercept("GET", "http://localhost:8081/contents", {
        fixture: "contents.json",
      }).as("getContents");
    });
  
    it("AC3:  Ensures access to the Search box in the Content page and allows you to see if it found an content by doing a search in the search box. Then it checks the text and image of the content. And then checks the go back button and clicks on the button to go back full articles page.", () => {
        cy.get('[data-cy="article-search-box"]').type("nice");
        cy.get('[data-cy="text"]').findByText("nice view sea side").should("exist");
        cy.get('[data-cy="image"]').should("exist");
        cy.get('[data-cy="C-goback-button"]').should("exist");
        cy.get('[data-cy="C-goback-button"]').click();
        cy.url().should("include", "/fullarticles"); 
    });
  });
  