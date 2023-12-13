import adapter from '@sveltejs/adapter-auto';
import preprocess from 'svelte-preprocess';

const config = {
	preprocess: preprocess({
		postcss: true,
	}),
	kit: {
		adapter: adapter(),
		alias: {
			$components: 'src/lib/components',
			$customTypes: 'src/app.d.ts',
			$helpers: 'src/helpers',
			$assets: 'src/assets',
			$fonts: 'src/lib/fonts',
			$styles: 'src/styles',
			$blog: 'src/posts',
		},
	},
};

export default config;
