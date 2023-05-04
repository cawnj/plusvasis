/* eslint-disable cypress/unsafe-to-chain-command */
/* eslint-disable cypress/no-unnecessary-waiting */

Cypress.on('window:before:load', (win) => {
	cy.stub(win.console, 'log').as('consoleLog');
});

describe('updateJob spec', () => {
	it('User updates a current job', () => {
		cy.intercept(
			'POST',
			'https://api.plusvasis.xyz/job/lK5VKiQIR6UpBXkZ35Ccft5Uyqg1-cypress-test'
		).as('jobUpdated');
		cy.intercept('GET', '/__data.json?x-sveltekit-invalidated=1_1').as('jobs');
		cy.visit('https://app.plusvasis.xyz/login');

		cy.wait(1500);

		cy.get('input[name="email"]').type('example@email.com');
		cy.get('input[name="password"]').type('th1s1sJustATest');
		cy.get('button[type="submit"]').click();

		cy.wait(1500);

		cy.wait('@jobs');
		cy.get('div[data-testid="job-list"]').should('exist');

		cy.get('a[href*="/container/lK5VKiQIR6UpBXkZ35Ccft5Uyqg1-cypress-test"]').click();

		cy.wait(1500);

		cy.contains('Settings').click();

		cy.get('input[name="dockerImage"]', { timeout: 10000 })
			.should('be.visible')
			.clear()
			.type('nginx', { delay: 100 })
			.should('have.value', 'nginx');

		cy.get('input[name="port"]', { timeout: 10000 })
			.should('be.visible')
			.clear()
			.type('80', { delay: 100 })
			.should('have.value', '80');

		cy.get('button[type="submit"]').click();

		cy.wait('@jobUpdated');
		cy.get('@consoleLog').should('be.calledWith', 'Container Updated');
	});
});
