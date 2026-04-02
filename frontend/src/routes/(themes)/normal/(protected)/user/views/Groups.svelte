<script lang="ts">
	import { browser } from "$app/environment";
	import { goto } from "$app/navigation";
	import PageLoading from "$lib/components/loading/PageLoading.svelte";
	import { seperateMemberOfGroups } from "$lib/parsers/members";
	import { onMount } from "svelte";


    let {
        userDN
    } : {
        userDN: string
    } = $props()

    let userGroupPromise: Promise<any> = $state(new Promise(() => {}))



    async function fetchUserGroups(): Promise<any> {
        if (userDN == "") goto("/search")
        const res =  await fetch(`http://localhost:8000/api/user?dn=${encodeURIComponent(userDN)}&attributes=memberof`)
        if (!res.ok && browser) {
            setTimeout(() => {
                goto("/search")
            }, 3000);
            throw new Error(`Failed to find User with ${userDN}`)
        }
        return res.json()
    }

    onMount(() => {
        userGroupPromise = fetchUserGroups()
    })

</script>

<section>
    {#await userGroupPromise}
        <PageLoading/>
    {:then user} 
        {#each seperateMemberOfGroups(user.memberof) as group}
            <p>
            {group}
            </p>
        {/each}
    {:catch error}
        <h1>{error.message} Redirecting in 3s</h1>
    {/await}
</section>

<style>
</style>