import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = ({ url }) => {
    const userDN = url.searchParams.get("dn");

    if (!userDN) {
        throw redirect(302, '/normal/search');
    }

    return {
        userDN
    };
};