import { Search } from '$lib/types/search';
import type { PageLoad } from './$types';

export const load: PageLoad = ({ url }) => {
    const searchValue = url.searchParams.get("search") ?? "";
    let filter = url.searchParams.get("filter") ?? "USERS"

    if (!Search.isValidFilter(filter)) {
        filter = "USERS"
    }

    filter = filter.toUpperCase() as Search.Filter;


    return {
        searchValue,
        filter
    };
};