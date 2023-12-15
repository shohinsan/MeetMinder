import { fail , redirect} from '@sveltejs/kit'
import type { Actions } from './$types'
import { isValidEmail, isValidPassword } from '$lib/utils'

// import { TOKEN_SECRET } from '$env/static/private'
// import {SignJWT} from "jose";
// import {isValidPassword} from "$lib/utils";

const API_ENDPOINT = "https://localhost:8443/api/auth";

interface Validation {
	nameMissing?: string
	usernameMissing?: string
	emailMissing?: string
	passwordMissing?: string
	confirmPasswordMissing?: string
	passwordMismatch?: string
	invalidEmail?: string
	weakPassword?: string
}

export const actions = {
    default: async ({request, cookies}) => {
		const data = await request.formData()
		// Form Data
		const step = data.get('step') as string
		const name = data.get('name') as string
		const username = data.get('username') as string
		const email = data.get('email') as string
		const password = data.get('password') as string
		const confirm_password = data.get('confirm_password') as string

		const collection = {
			name,
			username,
			email,
			password,
			confirm_password
		}

        // Validation
        if (!name || !username || !email || !password || !confirm_password) {
			let errors: Validation = {}
			if (!name) {
				errors.nameMissing = 'Name is required'
			}
			if (!username) {
				errors.usernameMissing = 'Username is required'
			}
			if (!email) {
				errors.emailMissing = 'Email is required'
			}
			if (!password) {
				errors.passwordMissing = 'Password is required'
			}
			if (confirm_password) {
				errors.confirmPasswordMissing = 'Confirm Password is required'
			}
			if (isValidEmail(email)) {
				errors.invalidEmail = 'Invalid email'
			}
			if (isValidPassword(password)) {
				errors.weakPassword = 'Password must be at least 8 characters long'
			}

			return fail(400, {
				required: !name || !username || !email || !password || !confirm_password,
				errors: errors
			})
		}

        // API Call
        const res = await fetch(`${API_ENDPOINT}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({collection}),
        });

        if (res.ok) {
            const { access_token } = await res.json()
            cookies.set('access_token', access_token, { path: '/' })
            redirect(303, '/');
        }

        return fail(400, { msg: "Something went wrong" })
    },
} satisfies Actions;



