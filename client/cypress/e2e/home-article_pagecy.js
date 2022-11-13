/* eslint-disable no-undef */
/// <reference types="Cypress"/>


describe("App Component", () => {

  it("AC1:Ensures to go to the main page and checks the header, motto and the link of the fullarticle page, then clicks on the link and checks the url of the fullarticle page.", () => {
    cy.visit("/");
    cy.get('[data-cy="home-page"]');
    cy.get('[data-cy="header"]').should("be.visible");
    cy.get('[data-cy="motto"]').should("exist");
    cy.get('[data-cy="link"]').should("exist");
    cy.get('[data-cy="link"]').click();
    cy.url().should("include", "/fullarticles");
  });

  it("AC2: Ensures to the existance of fullarticle page and whether there is data in it.", () => {
    cy.intercept("GET", "http://localhost:8081/fullarticles", {
      fixture: "articles.json",
    }).as("getArticles");
    cy.get('[data-cy="detail-page"]').should("exist");
  });

  it("AC3: Ensures access to the Search box and limit box in the Articles page and allows you to see if it found an article by doing a search in the search box. Then it checks the title, description,content, short_text, author name, author email and created date of the article.", () => {
    cy.request("/fullarticles").should("exist", ".inputBox");
    cy.get('[data-cy="header"]').should("exist");
    cy.get('[data-cy="limit-container"]').should("exist");
    cy.get('[data-cy="limit-searchbox"]').should("exist");
    cy.get('[data-cy="article-search-box"]').should("have.value", "");
    cy.get('[data-cy="article-search-box"]').type("jack");
    cy.get('[data-cy="title"]').findByText("Jack` holiday").should("exist");
    cy.get('[data-cy="description"]')
      .findByText("jack is creating new article thhis is description")
      .should("exist");
    cy.get('[data-cy="content"]').findByText("nice view sea side").should("exist");
    cy.get('[data-cy="short-text"]').findByText("this is jack`s 1.contents highlight").should("exist");
    cy.get('[data-cy="user"]').findByText("jack").should("exist");
    cy.get('[data-cy="email"]').findByText("jack@jack.com").should("exist");
    cy.get('[data-cy="date"]').findByText("2022-11-09 11:32:13").should("exist");
    cy.get('[data-cy="article-search-box"]').clear()
    //cy.get('[data-cy="article"]').eq(0);
    
  });


});
