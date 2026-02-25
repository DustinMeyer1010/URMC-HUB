<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import Group from '$lib/components/cards/Group.svelte';
	import { onMount } from 'svelte';

    let driveInfo: any = $state(undefined)
    let filterSearchValue: string = $state("")

    let drive: string = page.url.searchParams.get("name") ?? ""

    let groupFiltered: any = $derived.by(() => {
        if (driveInfo == undefined) {
            return 
        }

        if (filterSearchValue == "") {
            return driveInfo.groups
        }

        return driveInfo.groups.filter((group: any) => {
            if (group.name.toLowerCase().includes(filterSearchValue.toLocaleLowerCase())) {
                return group
            }
        })
    })


    onMount(async () => {

        await fetch(`http://localhost:8000/api/drive?drive=${encodeURIComponent(drive)}`)
        .then(async (res) => {
            if (res.status == 200) {
                driveInfo = await res?.json()
            }

        })

        if (driveInfo.drive == "") {
            goto("/search")
        }
    })
</script>


<main>
    <h1>{drive}</h1>
    <input type="text" bind:value={filterSearchValue}>

    {#if (driveInfo != undefined)}
        <div>
            {#each groupFiltered as group, idx (idx)} 
                <Group item={group} {idx}></Group>
            {/each}
        </div>
    {/if}
</main>


<style>
    main {
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    div {
        width: 90%;
        display: flex;
        flex-direction: column;
        gap: 1rem; 
    }
</style>