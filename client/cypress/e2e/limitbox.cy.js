/* eslint-disable no-undef */
/* eslint-disable react/no-unknown-property */
/* eslint-disable react/react-in-jsx-scope */

describe("App Component", () => {
  

  it("AC1:It provides filtering of articles by entering the limit value. It only takes the numerical value. When text is entered, no error occurs with the regex feature.", () => {
    cy.visit("/");
    cy.get('[data-cy="link"]').click();
    cy.get('[data-cy="limit-searchbox"]').should("have.value", "");
    cy.get('[data-cy="article"]').should('have.length', 4)
    cy.get('[data-cy="limit-searchbox"]').type("1");
    cy.get('[data-cy="limit-button"]').click();
    cy.get('[data-cy="article"]').should('have.length', 1)
    cy.get('[data-cy="limit-searchbox"]').clear()
    cy.get('[data-cy="limit-searchbox"]').type("0");
    cy.get('[data-cy="limit-button"]').click();
    cy.get('[data-cy="article"]').should('have.length', 0)
    cy.get('[data-cy="limit-searchbox"]').clear()
    cy.get('[data-cy="limit-searchbox"]').type("jhjhj");
    cy.get('[data-cy="limit-button"]').click();
    cy.get('[data-cy="article"]').should('have.length', 4)
    
  });
});
