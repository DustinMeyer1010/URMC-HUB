import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = ({ url }) => {
    const computerDN = url.searchParams.get("dn");

    if (!computerDN) {
        throw redirect(302, '/normal/search');
    }

    return {
        computerDN
    };
};