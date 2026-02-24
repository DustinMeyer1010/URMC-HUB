import { User } from "$lib/types/user";

interface LockoutState {
    LockoutInformation: User.LockoutInfo[]
    RefreshMsg: string
    RefreshTimeout: ReturnType<typeof setTimeout> | null
    Username: string
    RefreshLockout: () => void
}


/**
 * This class supports that state of the lockout component. 
 *
 * Tracks all the information on the page and the functions to make those updates
 */
export class LockoutStateClass implements LockoutState {
    Username: string;
    LockoutInformation: User.LockoutInfo[] = $state([])
    RefreshMsg: string = $state("")
    RefreshTimeout: NodeJS.Timeout | null = $state(null)
    loading: boolean = $state(false)
    isMount: boolean = false

    constructor(username: string) {
        this.Username = username
    }


    /**
     * Retrieves the lockout status for the user page that is loaded
     */
    RefreshLockout = async () => {
        if (this.RefreshTimeout != null) {
            clearTimeout(this.RefreshTimeout)
            this.RefreshMsg = ""
        }

        this.loading = true

        await fetch(`http://localhost:8000/api/user/${this.Username}/lockout`)
        .then(async (res) => {
            this.LockoutInformation = await res.json()


            this.LockoutInformation.sort((a,b) => new Date(b.time).getTime() - new Date(a.time).getTime())
            this.loading = false
            
            // Prevents refreshed from showing when the compoent is mounted
            if (!this.isMount){
                this.isMount = true
                return
            }
            this.RefreshMsg = "Refreshed"
            this.RefreshTimeout = setTimeout(() => {
                this.RefreshMsg = ""
            }, 1000)
        })
    }
}