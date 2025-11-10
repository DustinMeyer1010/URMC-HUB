import { AllFilters, type Groups } from "@t/filters"


export function getSearchValue(urlParams: URLSearchParams): string {
        let urlSearch = urlParams.get('search') || ""
        let storageSearch = localStorage.getItem("searchValue") || ""
            return (urlSearch !== "" ? decodeURI(urlSearch) : storageSearch)
}

export function getFilter(urlParams: URLSearchParams): Groups {

    let filter: Groups = "USERS"
    let urlFilter = urlParams.get('filter')?.toUpperCase() || ""
    let storagefilter = localStorage.getItem("filter")?.toUpperCase()  || ""

    if (!urlFilter && !storagefilter){
        return filter
    }

    console.log(urlFilter)
    if (AllFilters.includes(urlFilter)) {
        console.log("url")
        return urlFilter as Groups
    }

    if (AllFilters.includes(storagefilter)) {
        return storagefilter as Groups
    }


    return filter


}