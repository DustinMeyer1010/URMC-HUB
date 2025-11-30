import { goto } from "$app/navigation";
import { isLoggedIn } from "$lib/login";
import type { Filters, Results } from "@t/filters";
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
        this.SetURL()
    };

    SetURL() {
        goto(`?search=${this.santizedSearch}&filter=${this.filter.toLowerCase()}`, { replaceState: true, keepFocus: true, noScroll: true })
    }

}