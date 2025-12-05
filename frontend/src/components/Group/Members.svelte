<script lang="ts">
	import { onMount } from "svelte";

    let {
        group
    } : {
        group: string
    } = $props()

    let members: string[] = $state([])

    onMount(async () => {
        await fetch(`http://localhost:8000/api/group/${group}/members`).then( async (res) => {
            members = await res.json()
        })
    })
</script>

{#each members as member}
    {member}
{/each}