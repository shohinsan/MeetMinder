// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		interface Error {
			code?: string
			errorId?: string
		}
		interface Locals {
			user: Partial<Data> | null
			locale?: string
		}

		// interface PageData {}
		// interface Platform {}
	}

	declare interface Document {
		startViewTransition(callback: () => void): void
	}

}

interface NavItem {
	name: string;
	href: string;
}

type Data = {
	type: string
	message: string
	token: string
	userId: number
	meetingId: string
	name: string
	username: string
	email: string
	role: string
	password: string
	confirm_password: string
	image?: string
	bio?: string
}

type Restrict = Partial<Data>


export { NavItem, SidebarItem, Window, Document, Data, Restrict };
