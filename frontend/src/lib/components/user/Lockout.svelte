<script lang="ts">
    import RefreshIcon from '$lib/assets/Refresh-color-text.png'
	import { onMount } from 'svelte';
	import  { LockoutStateClass } from "./states/LockoutState.svelte";

    let {
        username
    } : {
        username: string
    } = $props();

    let LockoutState: LockoutStateClass = new LockoutStateClass(username)

    onMount(() => {
        LockoutState.RefreshLockout()
    })

</script>

    <section>
        <span>{LockoutState.RefreshMsg}<button class:loading={LockoutState.loading} onclick={LockoutState.RefreshLockout}><img src={RefreshIcon} alt=""></button></span>
            <div class="header">
                <div class="remove">Server</div>
                <div>Attempts</div>
                <div>Last Attempt</div>
            </div>
            {#each LockoutState.LockoutInformation as server}
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
        animation: slideIn 0.5s forwards;
        transform: translateY(100%);
        opacity: 0;
        animation-delay: 50ms;
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

    @keyframes slideIn {
        from {
            transform: translateY(100%);
            opacity: 0;
        }
        to {
            transform: translateY(0);
            opacity: 1;
        }
    }

</style>