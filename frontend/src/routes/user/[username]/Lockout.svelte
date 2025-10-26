<script lang="ts">
	import type { LockoutInfo } from "@t/user";
	import { onMount } from "svelte";
	import { fly } from "svelte/transition";
    import RefreshIcon from '$lib/assets/Refresh-color-text.png'

    const duration: number = 200;
    const delay: number = 250;

    let loading: boolean = $state(false)

    let {
        username
    } : {
        username: string
    } = $props();

    let data: LockoutInfo[] = $state([])
    let refreshedMsg: string = $state("")
    let refreshedTimeout: ReturnType<typeof setTimeout> | null = $state(null)


    onMount(async () => {
        refreshLockout()
    })

    async function refreshLockout() {
        if (refreshedTimeout != null) {
            clearTimeout(refreshedTimeout)
            refreshedMsg = ""
        }

        loading = true
        const response: Response = await fetch(`http://localhost:8000/lockout/${username}`)
        data = await response.json()

        data.sort((a, b) =>  new Date(b.time).getTime() - new Date(a.time).getTime())
        loading = false
        refreshedMsg = "Refreshed"
        refreshedTimeout = setTimeout(() => {
            refreshedMsg = ""
        }, 1000);
    }

</script>


<section in:fly={{delay: delay, duration: duration, x: -200, y: 200}} out:fly={{duration: duration, x: 200, y: 200}}>
    <span>{refreshedMsg}<button class:loading={loading} onclick={refreshLockout}><img src={RefreshIcon} alt=""></button></span>
        <div class="header">
            <div class="remove">Server</div>
            <div>Attempts</div>
            <div>Last Attempt</div>
        </div>
        {#each data as server}
        <div class="row">
            <div class="remove">{server.name}</div>
            <div>{server.count}</div>
            <div>{server.time}</div>
        </div>
            
        {/each}
</section>


<style>
    section {
        padding: 2rem 1rem;
        font-size: 15px;
        border-radius: 10px;
        flex-wrap: wrap-reverse;
        background: var(--color-surface);
        position: relative;
        min-width: 50%;
    }

    div.header {
        display: grid;
        font-weight: bold;
        text-align: center;
        grid-template-columns: 1fr 1fr 1fr;
        padding: 0.5rem;
    }

    div.row {
        display: grid;
        text-align: center;
        grid-template-columns: 1fr 1fr 1fr;
        padding: 0.5rem;
    }

    span {
        position: absolute;
        display: flex;
        justify-content: center;
        align-items: center;
        top: 0.5rem;
        right: 0.5rem;
        font-size: 12px;
    }
    
    button {

        background: none;
        border: none;
        transform: rotate(45deg);
        margin: 0;

    }

    button img {
        animation: spin 1s linear infinite;
        animation-play-state: paused;
        animation-delay: 0ms;
        width: 20px;
        
    }

    button:hover img {
        animation-play-state: running;

    }

    button.loading img {
        animation-play-state: running;
    }

    @keyframes spin {
        from {transform: rotate(45deg);}
        to {transform: rotate(405deg);}
    }

    @media (max-width: 850px) {
        section {
            width: 100%;
            min-width: 275px;
            padding: 1.5rem 0.5rem;
        }

        div.remove {
            display: none;
        }


        
    div.header,
    div.row {
        grid-template-columns: 0.5fr 1fr;
        padding: 0.5rem;
    }

    }

</style>