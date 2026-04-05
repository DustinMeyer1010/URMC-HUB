<script lang="ts">
	import { browser } from "$app/environment";
	import { onMount } from "svelte";


    let {userDN} : {userDN: string} = $props()

    let userDrives: Promise<any> = $state(new Promise(() => {}))


    async function fetchDrives(): Promise<any> {
        const res = await fetch(`http://localhost:8000/api/user/drives?dn=${encodeURIComponent(userDN)}`)
        if (!res.ok && browser) {
            throw new Error("Failed to load drives")
        }
        return res.json()
    }

    onMount(() => {
        userDrives = fetchDrives()
    })



</script>

{#await userDrives}
    <div>Fetching Drives</div>
{:then drives}
    {#each drives as drive}
        <div>
            <span>{drive.drive}</span>
            {#each drive.groups as group}
                <span>{group}</span>
            {/each}
        </div>
        
    {/each}

{:catch error}
    {console.log(error)}
    
{/await}


<style>

</style>