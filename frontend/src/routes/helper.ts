import { AllFilters, type Filters } from "@t/filters"


export function getSearchValue(urlParams: URLSearchParams): string {
        let urlSearch = urlParams.get('search') || ""
        let storageSearch = localStorage.getItem("searchValue") || ""
            return (urlSearch !== "" ? decodeURI(urlSearch) : storageSearch)
}

export function getFilter(urlParams: URLSearchParams): Filters {

    let filter: Filters = "USERS"
    let urlFilter = urlParams.get('filter')?.toUpperCase() || ""
    let storagefilter = localStorage.getItem("filter")?.toUpperCase()  || ""

    if (!urlFilter && !storagefilter){
        return filter
    }

    console.log(urlFilter)
    if (AllFilters.includes(urlFilter)) {
        console.log("url")
        return urlFilter as Filters
    }

    if (AllFilters.includes(storagefilter)) {
        return storagefilter as Filters
    }


    return filter


}