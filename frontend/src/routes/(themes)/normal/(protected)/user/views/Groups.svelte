<script lang="ts">
	import { browser } from "$app/environment";
	import Group from "$lib/components/cards/Group.svelte";
	import PageLoading from "$lib/components/loading/PageLoading.svelte";
	import { onMount } from "svelte";



    let {
        userDN
    } : {
        userDN: string
    } = $props()

    let userGroupPromise: Promise<any> = $state(new Promise(() => {}))
    let searchFilter: string = $state("")
    let allGroups: any = $state([])
    let filteredGroups: any = $derived.by(() => {
        if (searchFilter == "") {
            return allGroups
        }
        console.log(searchFilter)
        return allGroups.filter((group: any) => group.cn.join().toLowerCase().includes(searchFilter))
    })



    async function fetchUserGroups(): Promise<any> {
        const res =  await fetch(`http://localhost:8000/api/user/groups?dn=${encodeURIComponent(userDN)}`)
        if (!res.ok && browser) {
            throw new Error("Failed to load groups")
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
    {:then groups}
        <div>
            <span hidden>{allGroups = groups}</span>
            <span>Search: </span><input bind:value={searchFilter}>
        </div>
        {#each filteredGroups as group, idx (group.dn)}
            <Group item={group} {idx}>
            </Group>
        {/each}
    {:catch error}
        <h1>{error.message} Redirecting in 3s</h1>
    {/await}
</section>

<style>

    section {
        display: flex;
        flex-direction: column;
        gap: 1rem;
        justify-content: center;
        align-items: center;
    }

    input {
        background: var(--color-surface);
        border: none;
        height: 30px;
        width: 30%;
        border-radius: 8px;
        color: var(--color-text);
        padding: 10px;
        font-size: 15px;
    }

    div {
        position: sticky;
        z-index: 10;
        top: 0;
        display: flex;
        gap: 1rem;
        justify-content: center;
        align-items: center;
        width: 100%;
        padding: 1em;
        background: var(--color-bg);
        backdrop-filter: blur(20px);
    }

    input:focus {
        outline: 1px solid var(--color-text);
    }
</style>