/* eslint-disable cypress/unsafe-to-chain-command */
/* eslint-disable cypress/no-unnecessary-waiting */

describe('startJob spec', () => {
	it('User starts a job that was previously stopped', () => {
		cy.intercept(
			'GET',
			'https://api.plusvasis.xyz/job/lK5VKiQIR6UpBXkZ35Ccft5Uyqg1-cypress-test/start'
		).as('jobStarted');
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

		cy.window().then((w) => (w.beforeReload = true));
		cy.window().should('have.prop', 'beforeReload', true);

		cy.contains('Start').click();

		cy.wait('@jobStarted');
		cy.window().should('not.have.prop', 'beforeReload');
	});
});
