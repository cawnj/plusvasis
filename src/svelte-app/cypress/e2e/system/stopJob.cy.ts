/* eslint-disable cypress/unsafe-to-chain-command */
/* eslint-disable cypress/no-unnecessary-waiting */

describe('stopJob spec', () => {
	it('User stops a current job', () => {
		cy.intercept(
			'DELETE',
			'https://api.plusvasis.xyz/job/lK5VKiQIR6UpBXkZ35Ccft5Uyqg1-cypress-test'
		).as('jobStopped');
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

		cy.contains('Stop').click();

		cy.wait('@jobStopped');
		cy.window().should('not.have.prop', 'beforeReload');
	});
});
