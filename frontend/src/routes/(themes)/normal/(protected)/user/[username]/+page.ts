export const prerender = true;

import type { PageLoad } from './$types';;

export const load: PageLoad = async ({ params }) => {

    return {username: params.username} ;
};
