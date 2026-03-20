
import { writable } from "svelte/store";


async function createLoginState() {

    var stored: boolean = true
    await fetch("http://localhost:8000/api/verify", {mode: "cors"})
    .then((res) => {
        if (res.status == 200) {
            stored = true
        } else {
            stored = false
        }
    }).catch(() => {
        stored = false
    })

    const { subscribe, set } = writable(stored)

    return {
        subscribe,
        Login: () => {
            set(true)
        },
        Logout: () => {
            set(false)
        },
        CheckStatus: async (): Promise<boolean> => {
            var verified = false
            await fetch("http://localhost:8000/api/verify", {mode: "cors"})
            .then((res) => {
                if (res.status == 200) {
                    set(true)
                    verified = true
                    
                } else {
                    set(false)
                }
            }).catch(() => {
                set(false)
            })
            return verified
        }
    }
}

export const isLoggedIn = await createLoginState()

