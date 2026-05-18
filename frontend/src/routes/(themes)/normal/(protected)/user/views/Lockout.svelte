<script lang="ts">
	import { browser } from "$app/environment";
	import { goto } from "$app/navigation";
	import { onMount } from "svelte";
    import refreshIcon from "$lib/assets/Icons/dark/refresh.png"



    let { userDN } : {userDN: string} = $props()

    let lockoutStatus: Promise<any> = $state(new Promise(() => {}))


    async function fetchLockoutStatus(): Promise<any> {
        const res = await fetch(`http://localhost:8000/api/user/lockout?dn=${encodeURIComponent(userDN)}`);
        if (!res.ok && browser) {
            goto("/normal/search")
            throw new Error(`Failed to find User with ${userDN}`);
        }
        return res.json()
    }


    onMount(() => {
        lockoutStatus = fetchLockoutStatus()
    })

    async function refresh() {
        lockoutStatus = fetchLockoutStatus()
    }

</script>

<section>
    <button onclick={refresh}>Refresh</button>
    {#await lockoutStatus}
        <img src={refreshIcon} alt="">
    {:then statuses}

        <div>
            <span id="count">Count</span>
            <span>Sever</span>
            <span>Time</span>
        </div>
        {#each statuses as status}
            <div>
                <span id="count">{status.count}</span>
                <span>{status.name}</span>
                <span>{status.time}</span>
            </div>
            
        {/each}
    {/await}

</section>


<style>

    section {
        position: relative;
        display: flex;
        flex-direction: column;
        gap: 1rem;
        height: 90%;
        width: 100%;
    }

    section div {
        display: flex;
        justify-content: space-evenly;
        flex-grow: 1;
        align-items: center;
        text-align: center;
    }

    
    section div span#count {
        flex-basis: 10%;
    }

    span {
        text-align: center;
        width: 100%;
        font-size: 18px;
    }

    button {
        background: var(--color-surface);
        align-self: center;
        color: var(--color-text);
        padding: 10px;
        border: none;
        border-radius: 8px;
        width: 50%;
    }

    button:hover {
        background: var(--color-surface-hover)
    }

    img {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);

        width: 50px;
        animation: spin 1s linear infinite;
    }

    @keyframes spin {
        from {
            transform: rotate(0deg);
        }
        to {
            transform: rotate(360deg);
        }
    }

    @media (max-width: 700px) {
        section {
            width: 100%;
        }
    }

</style>