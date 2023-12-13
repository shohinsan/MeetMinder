import {fail, type Actions, redirect} from '@sveltejs/kit';
// import { TOKEN_SECRET } from '$env/static/private'
// import {SignJWT} from "jose";
// import {isValidPassword} from "$lib/utils";

const API_ENDPOINT = "https://localhost:8443/api/auth";

interface Validation {
    emailMissing?: string;
    passwordMissing?: string;
    passwordMismatch?: string;
}

export const actions = {
    default: async ({request, cookies}) => {
        const data = await request.formData();

        // Form Data
        const email = data.get("username") as string;
        const password = data.get("password") as string;
        const confirm_password = data.get("confirm_password") as string;

        // Validation
        if (!email || !password || password !== confirm_password) {
            const errors: Validation = {};
            if (!email) {
                errors.emailMissing = "Email is required";
            }
            if (!password) {
                errors.passwordMissing = "Password is required";
            }
            if (password !== confirm_password) {
                errors.passwordMismatch = "Passwords do not match";
            }
            return fail(400, {
                required: !email || !password,
                errors: errors
            });
        }

        // API Call
        const res = await fetch(`${API_ENDPOINT}/login`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            credentials: 'include',
            body: JSON.stringify({email, password}),
        });

        // const tokenSecret = new TextEncoder().encode(process.env.TOKEN_SECRET);
        // const token = await new SignJWT({
        //     userId: res.user.userId,
        //     username: res.user.username,
        //     role: res.user.role,
        // })
        //     .setProtectedHeader({alg: 'HS256', typ: 'JWE'})
        //     .setIssuedAt()
        //     .setExpirationTime('1w')
        //     .setJti('1a2b-3c4d-5e6f-7g8h')
        //     .setSubject('auth')
        //     .sign(tokenSecret);
        // cookies.set('access-token', token, {
        //     path: '/',
        //     httpOnly: true,
        //     sameSite: 'strict',
        //     maxAge: 24 * 60 * 60,
        //     secure: true,
        // });

        if (res.ok) {
            const { access_token } = await res.json()
            cookies.set('access_token', access_token, { path: '/' })
            throw redirect(303, '/');
        }
        return fail(400, { msg: "Something went wrong" })
    },
} satisfies Actions;
