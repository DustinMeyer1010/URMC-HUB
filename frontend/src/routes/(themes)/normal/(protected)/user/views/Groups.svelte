<script lang="ts">
	import { browser } from "$app/environment";
	import Group from "$lib/components/cards/Group.svelte";
	import PageLoading from "$lib/components/loading/PageLoading.svelte";
	import { pullNameFromDN } from "$lib/parsers/ou";
	import { onMount } from "svelte";
    import type { Group as g } from "$lib/types/group";
	import { Icons } from "$lib/managers/icons";



    let {
        userDN
    } : {
        userDN: string
    } = $props()

    let userGroupPromise: Promise<g.CardInfo[]> = $state(new Promise(() => {}))
    let searchFilter: string = $state("")
    let allGroups: g.CardInfo[] = $state([])

    let modificationResult: any = $state({})
    let messageShow: boolean = $state(false)


    let filteredGroups: g.CardInfo[] = $derived.by(() => {
        if (searchFilter == "") {
            return allGroups
        }
        console.log(searchFilter)
        return allGroups.filter((group: g.CardInfo) => group.cn.join().toLowerCase().includes(searchFilter.toLocaleLowerCase()))
    })



    async function fetchUserGroups(): Promise<any> {
        const res =  await fetch(`http://localhost:8000/api/user/groups?dn=${encodeURIComponent(userDN)}`)
        if (!res.ok && browser) {
            throw new Error("Failed to load groups")
        }
        return res.json()
    }

    async function removeGroup(groupDN: string) {
        const res = await fetch(`http://localhost:8000/api/user/group?dn=${userDN}`, 
        {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({group: groupDN})
        })

        setTimeout(() => {
            messageShow = false
        }, 3000);

        modificationResult = await res.json()
        messageShow = true
    }

    onMount(() => {
        userGroupPromise = fetchUserGroups()
    })

</script>

<section>

    {#if messageShow}
        {@render ResultMessage()}
    {/if}

    {#await userGroupPromise}
        <PageLoading/>
    {:then groups}
        <div>
            <span hidden>{allGroups = groups}</span>
            <span>Search: </span><input bind:value={searchFilter}>
        </div>
        {#each filteredGroups as group, idx (group.dn)}
            <Group item={group} {idx}>
                <button id="group" aria-label="Remove Group" onclick={() => removeGroup(group.dn.join())}><img src={Icons.TRASH_CAN} alt=""></button>
            </Group>
        {/each}
    {:catch error}
        <h1>{error.message} Redirecting in 3s</h1>
    {/await}
</section>

{#snippet ResultMessage()}
    <div id="message" class:error={!modificationResult.successful}>
        <span>{modificationResult.errorMessage.toUpperCase()}</span>
        <span>|</span>
        <span>{pullNameFromDN(modificationResult.group)}</span>
    </div>
{/snippet}

<style>


    div#message {
        display: flex;
        gap: 1rem;
        width: fit-content;
        padding: 0.5rem 1rem;
        border: 1px solid var(--color-text);
        border-radius: 5px;
    }

    div#message.error {
        border: 1px solid var(--color-error);
    }

    section {
        display: flex;
        flex-direction: column;
        gap: 1rem;
        justify-content: center;
        align-items: center;
    }


    button#group {
        position: absolute;
        top: 8px;
        right: 8px;
        background: none;
        border: none;
    }

    button#group img {
        width: 25px;
        transition: 0.3s ease-in-out;
    }

    button#group:hover img {
        transform: scale(1.1);
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