import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ url: { searchParams } }) => {
	return {
		items: 'Schedules',
	}
}
