<script lang="ts">
	import type { GroupSimpleInfo } from "@types/group";


    let {
        groups
    } : {
        groups: GroupSimpleInfo[]
    } = $props()

</script>

{#if groups.length != 0}
{#each groups as g}
    <ul>
        <li>{g.name}</li>
        <li>{g.description}</li>
        <li>{g.information}</li>
        <li>{g.ou}</li>
    </ul>
    
{/each}
{:else}
No Groups Found
{/if}