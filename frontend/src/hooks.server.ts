import type { Handle } from '@sveltejs/kit'
import { sequence } from '@sveltejs/kit/hooks'
// import { jwtVerify } from 'jose'
// import { redirect } from '@sveltejs/kit'
// import type { Restrict } from '$customTypes'
// import { handleLoginRedirect } from '$lib/utils'


// const auth: Handle = async ({ event, resolve }) => {
// 	const token: string | undefined = event.cookies.get('access_token')
// 	let data: Restrict | null = null
// 	if (token) {
// 		const secret: Uint8Array = new TextEncoder().encode(process.env.TOKEN_SECRET)
// 		const decoded = await jwtVerify(token, secret)
// 		const { userId, role } = decoded.payload
// 		data = {
// 			userId: userId as number,
// 			role: role as string
// 		}
// 		if (!event.locals.user) {
// 			throw redirect(302, handleLoginRedirect(event))
// 		}
// 	}
// 	event.locals.user = data
// 	return resolve(event)
// }

const i18n: Handle = async ({ event, resolve }) => {
	const locale = 'en'
	event.locals.locale = locale
	return resolve(event, {
		transformPageChunk: ({ html }) => {
			return html.replace('%lang%', locale)
		}
	})
}


export const handle: Handle = sequence(i18n)
