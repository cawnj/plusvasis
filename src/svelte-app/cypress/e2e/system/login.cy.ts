/* eslint-disable cypress/no-unnecessary-waiting */

describe('login spec', () => {
	it('logs into testing account', () => {
		cy.intercept('GET', 'https://api.plusvasis.xyz/jobs').as('jobs');
		cy.visit('https://app.plusvasis.xyz');

		// check if user is logged in
		if (cy.get('a[href*="/logout"]')) {
			cy.get('a[href*="/logout"]').click();
		}

		cy.get('input[name="email"]').type('example@email.com');
		cy.get('input[name="password"]').type('th1s1sJustATest');
		cy.get('button[type="submit"]').click();

		// Wait for a few seconds (e.g., 3 seconds) before proceeding
		cy.wait(3000);

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
