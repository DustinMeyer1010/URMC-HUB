export const prerender = false

import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import type { UserFullInfo } from '@t/user';

export const load: PageLoad = async ({ params, fetch }) => {
    const response: Response = await fetch(`http://localhost:8000/user/info/${params.username}`);

    if (!response.ok) {
        throw error(response.status, 'Failed to load user');
    }

    const data: UserFullInfo = await response.json();

    return data ;
};
