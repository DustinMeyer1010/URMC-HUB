<script lang="ts">
	import Group from '$lib/components/cards/Group.svelte';
	import { onMount } from 'svelte';

    let driveInfo: any = $state(undefined)
    let filterSearchValue: string = $state("")

    let {
        data,
    } : {
        data: {drive_name: string}
    } = $props()

    let groupFiltered: any = $derived.by(() => {
        if(driveInfo == undefined) {
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
        const params = new URLSearchParams();
        params.set("drive", "\\\\GGE-Depts.FLH.local\\departments$\\Nurse Recruiter");


        await fetch(`http://localhost:8000/api/drive?drive=${encodeURIComponent(data.drive_name)}`)
        .then(async (res) => {
            driveInfo = await res.json()
            console.log(driveInfo)
        })
    })
</script>


<main>
    <h1>{data.drive_name}</h1>
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