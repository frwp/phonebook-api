/// <reference types="cypress" />

const initialItems = [
  {
    name: 'Cypress',
    phone_number: '123456789',
  },
  {
    name: 'Presscy',
    phone_number: '987654321',
  },
];

const getItems = () => cy.request('/').its('body');

const add = (item) => {
  cy.request('POST', '/', item);
};

const deleteItem = (item) => {
  cy.request('DELETE', `/${item.id}`);
};

const deleteAll = () => {
  getItems().each(deleteItem);
};

const reset = () => {
  deleteAll();
  initialItems.forEach(add);
};

describe('Contacts GET API', () => {
  beforeEach(() => {
    reset();
    cy.request('/').as('root');
  });

  it('Validate the header', () => {
    cy.get('@root')
      .its('headers')
      .its('content-type')
      .should('include', 'application/json; charset=utf-8');
  });

  it('Validate status code', () => {
    cy.get('@root').its('status').should('equal', 200);
  });

  it('Validate response', () => {
    cy.get('@root').its('body').should('be.an', 'array');
  });

  it('Each item should have name', () => {
    getItems().each((value) => {
      expect(value).to.have.all.keys('id', 'name', 'phone_number');
    });
  });

  it('Initial value should contains 2 objects', () => {
    getItems().should('have.length', 2);
  });
});

describe('Contacts POST API', () => {
  it('Test adding new contact', () => {
    cy.request('POST', '/', { name: 'Cypress3', phone_number: '987654321' })
      .its('headers')
      .its('content-type')
      .should('include', 'application/json; charset=utf-8');
  });

  it('Test adding dupplicate contact', () => {
    cy.request({
      method: 'POST',
      url: '/',
      body: {
        name: 'Cypress',
        phone_number: '987654321',
      },
      failOnStatusCode: false,
    })
      .its('status')
      .should('equal', 409);
  });

  it('Test adding incomplete contact 1', () => {
    cy.request({
      method: 'POST',
      url: '/',
      body: {
        name: 'zupper',
      },
      failOnStatusCode: false,
    })
      .its('status')
      .should('equal', 400);
  });

  it('Test adding incomplete contact 1', () => {
    cy.request({
      method: 'POST',
      url: '/',
      body: {
        phone_number: '789456123',
      },
      failOnStatusCode: false,
    })
      .its('status')
      .should('equal', 400);
  });
});

const getLastIndex = () => getItems();

describe('Contact by Id API', () => {
  beforeEach(() => {
    reset();
  });

  it('Test GET item response', () => {
    getItems().each((item) => {
      cy.request(`/${item.id}`).its('status').should('equal', 200);
    });
  });

  it('Test GET item response', () => {
    getItems().each((item) => {
      cy.request(`/${item.id}`)
        .its('body')
        .then((value) =>
          expect(value).to.have.all.keys('id', 'name', 'phone_number')
        );
    });
  });

  it('Test GET invalid id', () => {
    cy.request({ url: `/${999}`, failOnStatusCode: false })
      .its('status')
      .should('equal', 404);
  });

});
