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

		cy.wait('@jobs');
		cy.get('span.font-medium').should('exist');
	});
});
