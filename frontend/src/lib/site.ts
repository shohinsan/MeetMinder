import { dev } from '$app/environment'

export const siteConfig = {
	name: 'MeeetMinder',
	url: dev ? 'http://localhost:5173' : 'http://localhost:5173',
	ogImage: 'S3DeployedImageURL',
	description: 'Website description',
	links: {
		twitter: 'yourTwitterBusinessUrl',
		github: 'yourGithubBusinessUrl',
		shadTwitter: 'personalTwitterUrl',
		shadGithub: 'personalGithubUrl'
	},
	keywords: `meeting, reminder, teams`
}

export type SiteConfig = typeof siteConfig
