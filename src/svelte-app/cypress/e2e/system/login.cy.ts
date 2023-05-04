/* eslint-disable cypress/no-unnecessary-waiting */

describe('login spec', () => {
	it('logs into testing account', () => {
		cy.intercept('GET', '/__data.json?x-sveltekit-invalidated=1_1').as('jobs');
		cy.visit('https://app.plusvasis.xyz/login');

		cy.wait(1500);

		cy.get('input[name="email"]').type('example@email.com');
		cy.get('input[name="password"]').type('th1s1sJustATest');
		cy.get('button[type="submit"]').click();

		// Wait for a few seconds (e.g., 3 seconds) before proceeding
		cy.wait(1500);

		cy.wait('@jobs');
		cy.get('body').then(($body) => {
			if ($body.find('span.font-medium').length > 0) {
				cy.get('span.font-medium').should('exist');
			} else {
				cy.get('div[data-testid="job-list"]').should('exist');
			}
		});
	});
});
