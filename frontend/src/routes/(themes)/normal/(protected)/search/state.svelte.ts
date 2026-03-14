import { goto } from "$app/navigation";
import { page } from "$app/state";
import { isLoggedIn } from "$lib/stores/login";
import { type Results } from "$lib/types/filters";
import { Search } from "$lib/types/search";

const emptyResults: Search.Results = {
    users: [],
    computers: [],
    groups: [],
    printers: [],
    drives: [],
}




export class SearchStateClass {
    data: Search.Results = $state(emptyResults)


    constructor() {
        $inspect(this.currentFilterItems)
    }

    currentFilterItems: Results = $state([])
    filter: Search.Filter = $state("USERS")
    loading: boolean = $state(true)
    searchValue: string = $state("")
    santizedSearch: string = $derived(encodeURIComponent(this.searchValue.trim()));
    


    Search = async () => {

        
        if (this.searchValue.length == 0) {
            return
        }
        this.loading = true

        await fetch(`http://localhost:8000/api/search?value=${this.santizedSearch}`)
        .then(async (res) => {
            this.data = await res.json()
        }).catch(() => isLoggedIn.Logout());
        this.SwitchFilter(this.filter)
        this.loading = false
    }

    SwitchFilter = (newFilter: Search.Filter) => {
        this.filter = newFilter;
        window.scrollTo(0,0)
        this.SetURLParams()
    };

    GetURLParams = () => {
        let urlParams = page.url.searchParams
        this.searchValue = urlParams.get("search") ?? ""

        let filter = urlParams.get("filter")?.toUpperCase() ?? "USERS"
        this.filter = Search.isValidFilter(filter) ? filter as Search.Filter : "USERS"

        if (this.filter == "USERS") {
            filter = localStorage.getItem("filter")?.toUpperCase() ?? "USERS"
            this.filter = Search.isValidFilter(filter) ? filter as Search.Filter : "USERS"
        }

        if (this.searchValue == "") {
            this.searchValue = localStorage.getItem("search") ?? ""
        }
        this.searchValue = decodeURIComponent(this.searchValue)



    }

    SetURLParams() {
        localStorage.setItem("filter", this.filter)
        localStorage.setItem("search", encodeURIComponent(this.searchValue) )
        goto(`?search=${this.santizedSearch}&filter=${this.filter.toLowerCase()}`, { replaceState: true, keepFocus: true, noScroll: true })

    }

}