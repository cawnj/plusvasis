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
	}),
	rest.get(`${hostname}/job/job123`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	}),
	rest.get(`${hostname}/job/nomadClient`, (req, res, ctx) => {
		return res(
			ctx.status(200),
			ctx.json({
				Name: 'test',
				TaskGroups: [
					{
						Tasks: [
							{
								Config: {
									Image: 'test'
								},
								Resources: {
									CPU: '100',
									MemoryMB: '300'
								}
							}
						]
					}
				],
				Meta: {
					shell: 'sh',
					volumes: [],
					env: 'test',
					port: '8080',
					expose: 'true'
				}
			})
		);
	}),
	rest.post(`${hostname}/job/nomadClient`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	}),
	rest.delete(`${hostname}/job/nomadClient`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	}),
	rest.post(`${hostname}/job/nomadClient/restart`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	}),
	rest.get(`${hostname}/job/nomadClient/start`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	}),
	rest.get(`${hostname}/job/nomadClient/alloc`, (req, res, ctx) => {
		return res(ctx.status(200), ctx.json({}));
	})
];
