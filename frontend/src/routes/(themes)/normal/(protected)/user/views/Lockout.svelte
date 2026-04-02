<script lang="ts">
	import { browser } from "$app/environment";
	import { goto } from "$app/navigation";
	import { onMount } from "svelte";



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

</script>

<section>

    {#await lockoutStatus}
        <div> awaiting for information</div>
    {:then statuses}
        {#each statuses as status}
            <div>
                <span>{status.count}</span>
                <span>{status.name}</span>
                <span>{status.time}</span>
            </div>
            
        {/each}
    {/await}

</section>


<style>

    section div {
        display: flex;
        gap: 20px;
    }

    span {
        font-size: 20px;
    }

</style>