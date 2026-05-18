import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = ({ url }) => {
    const userDN = decodeURIComponent(url.searchParams.get("dn")?? "");
    let section = url.searchParams.get("section") ?? "GROUPS";

    if (!userDN) {
        throw redirect(302, '/normal/search');
    }


    if (!["GROUPS", "ADD", "LOCKOUT", "DRIVE"].includes(section)) {
        section = "GROUPS"
    }

    return {
        userDN,
        section
    };
};