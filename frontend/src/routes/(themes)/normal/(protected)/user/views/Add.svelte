<script lang="ts">
	import Group from "$lib/components/cards/Group.svelte";
	import { pullNameFromDN } from "$lib/parsers/ou";
	import { onMount } from "svelte";

    let { userDN } : {userDN: string} = $props()
    let search: string = $state("")
    let groups: Promise<any> = $state(new Promise(() => {}))
    let commonGroups: string[] = ["ISDG_CTX_eRecord", "ISDG_CTX_eRecord2", "ISDU_CitrixAccessGateway", "ISDU_VPN_GP_FullAccess", "URMC_PACS_CLINICAL", "ISDL_PyxisUser_C", "ISDG_Imprivata_Users", "ISDG_CustomApp_OutboundEx_NewReport"]
    let commonGroupsValues: any[] = $state([])
    let modificationResult: any = $state({})
    let messageShow: boolean = $state(false)

    async function searchGroups(e: SubmitEvent) {
        e.preventDefault()

        if (search.length == 0) {
            fetchCommonGroups()
            groups = new Promise(() => {})
            return
        }
        const res = await fetch(`http://localhost:8000/api/search/groups?value=${encodeURIComponent(search)}&attributes=cn,samaccountname,information,description,dn`)

        groups = await res.json()
    }

    async function addGroup(groupDN: string) {
        const res = await fetch(`http://localhost:8000/api/user/group?dn=${userDN}`, 
        {
            method: "POST",
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

    async function fetchCommonGroups() {
        for (let i =0; i < commonGroups.length; i++) {
            const res = await fetch(`http://localhost:8000/api/search/groups?value=${encodeURIComponent(commonGroups[i])}&attributes=cn,samaccountname,information,description,dn`)
            if (res.ok) {
                let jsonData = await res.json()
                commonGroupsValues.push(jsonData[0])
            }
        }
        console.log(commonGroupsValues)
    }

    onMount(() => fetchCommonGroups())



</script>

<section>

    {#if messageShow}
        {@render ResultMessage()}
    {/if}
    <form onsubmit={searchGroups} action="submit">
        <span>Search: </span> <input type="text" bind:value={search}>
        <button type="submit">Search</button>
    </form>



    <div id="groups" >
        {#await groups}
            {#each commonGroupsValues as group, idx (group.dn.join())}
                <Group item={group} {idx}>
                    <button onclick={() => addGroup(group.dn.join())}>Add</button>
                </Group>
            {/each}
        {:then groupsList} 
            {#each groupsList as group, idx (group.dn.join())}
                <Group item={group} {idx}>
                    <button onclick={() => addGroup(group.dn.join())}>Add</button>
                </Group>
            {/each}
        {/await}

    </div>

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
        padding: 0.5rem 1rem;
        border: 1px solid var(--color-text);
        border-radius: 5px;
    }

    div#message.error {
        border: 1px solid var(--color-error);
    }

    section {
        position: relative;
        display: flex;
        flex-direction: column;
        align-items: center;
        padding-top: 1rem;
        gap: 1rem;
    }
    
    form {
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 1rem;
        width: 100%;
        position: sticky;
        padding: 10px 0px;
        top: 0;
        background: var(--color-bg);
        z-index: 50;
    }

    form button {
        background: var(--color-surface);
        color: var(--color-text);
        border-radius: 5px;
        padding: 0.5rem 1rem;
        border: none;

    }

    form button:hover {
        background: var(--color-surface-hover);
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

    div#groups {
        display: flex;
        flex-direction: column;
        gap: 1rem;
        width: 100%;
    }

    div#groups button {
        position: absolute;
        top: 10px;
        right: 10px;
    }

</style>