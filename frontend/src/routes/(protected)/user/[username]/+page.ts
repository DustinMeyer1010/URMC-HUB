export const prerender = true;

import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import type { UserFullInfo } from '@t/user';

export const load: PageLoad = async ({ params, fetch }) => {
    const response: Response = await fetch(`http://localhost:8000/api/user/info/${params.username}`);

    if (!response.ok) {
        throw redirect(301, "/search")
    }

    const data: UserFullInfo = await response.json();

    return data ;
};
