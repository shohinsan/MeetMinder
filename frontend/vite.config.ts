import {defineConfig, loadEnv} from 'vite';
import {sveltekit} from '@sveltejs/kit/vite';

export default defineConfig(({ mode }: { mode: string }) => {
	const env = loadEnv(mode, process.cwd());
	Object.assign(process.env, env);
	return {
		plugins: [sveltekit()],
		define: {
			'process.env': process.env,
		},
		server: {
			warmup: {
				clientFiles: ['./src/lib/site.ts'],
			},
		},
	};
});

