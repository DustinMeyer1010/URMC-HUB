import { goto } from "$app/navigation";
import { page } from "$app/state";
import { isLoggedIn } from "$lib/login";
import { Filter, type Filters, type Results } from "@t/filters";
import type { AllResults } from "@t/resutls";

const emptyResults: AllResults = {
    users: [],
    computers: [],
    groups: [],
    printers: [],
    drives: [],
}

interface SearchStateInterface {
    data: AllResults 
    currentFilterItems: Results
    filter: Filters
    loading: boolean
    searchValue: string
    santizedSearch: string
    Search: () => Promise<void>
    SwitchFilter: (newFilter: Filters) => void

}



export class SearchStateClass implements SearchStateInterface {
    data: AllResults = $state(emptyResults)


    constructor() {
        $inspect(this.currentFilterItems)
    }

    currentFilterItems: Results = $state([])
    filter: Filters = $state("USERS")
    loading: boolean = $state(true)
    searchValue: string = $state("")
    santizedSearch: string = $derived.by(() => {
        return encodeURIComponent(this.searchValue.trim())
    });
    


    Search = async () => {

        this.loading = true
        if (this.searchValue == "") {
            this.loading = false
            return
        }

        await fetch(`http://localhost:8000/api/search/all/${this.santizedSearch}`)
        .then(async (res) => {
            if (res.status != 200) {
                this.data = emptyResults
                return
            }
            this.data = await res.json()
        }).catch(() => isLoggedIn.Logout());

        this.SwitchFilter(this.filter)
        this.loading = false
    }

    SwitchFilter = (newFilter: Filters) => {
        this.filter = newFilter;
        window.scrollTo(0,0)
        this.SetURLParams()
    };

    GetURLParams = () => {
        let urlParams = page.url.searchParams
        this.searchValue = urlParams.get("search") ?? ""
        let filter = urlParams.get("filter")?.toUpperCase() ?? "USERS"
        this.filter = Filter.isValid(filter) ? filter as Filters : "USERS"

        if (this.filter == "USERS") {
            this.filter = localStorage.getItem("filter")?.toUpperCase() as Filters ?? "USERS"
        }

        if (this.searchValue == "") {
            this.searchValue = localStorage.getItem("search") ?? ""
        }


    }

    SetURLParams() {
        localStorage.setItem("filter", this.filter)
        localStorage.setItem("search", this.searchValue)
        goto(`?search=${this.santizedSearch}&filter=${this.filter.toLowerCase()}`, { replaceState: true, keepFocus: true, noScroll: true })

    }

}