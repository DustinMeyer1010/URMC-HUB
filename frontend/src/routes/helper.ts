export function getSearchValue(urlParams: URLSearchParams): string {
        let urlSearch = urlParams.get('search') || ""
        let storageSearch = localStorage.getItem("searchValue") || ""
            return (urlSearch !== "" ? decodeURI(urlSearch) : storageSearch)
}