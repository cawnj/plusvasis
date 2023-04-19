import { rest } from 'msw';
import { hostname } from '../stores/environmentStore';

// Define handlers that catch the corresponding requests and return the mock data.
// Will add handler later
export const handlers = [
	rest.get(`${hostname}/jobs`, (req, res, ctx) => {
		return res(
			ctx.status(200),
			ctx.json([
				{
					ID: 'job123',
					Name: 'job123'
				}
			])
		);
	}),
	rest.get(`https://nomad.local.cawnj.dev/v1/client/fs/logs/alloc-id-123`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	}),
	rest.get(`${hostname}/job/job123/alloc`, (req, res, ctx) => {
		return res(
			ctx.status(200),
			ctx.json({
				ID: 'alloc-id-123',
				TaskStates: {
					job123: {
						State: 'running'
					}
				}
			})
		);
	}),
	rest.post(`${hostname}/jobs`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	}),
	rest.get(`${hostname}/job/job123/logs`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	}),
	rest.get(`${hostname}/job/job123/exec`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	})
];
