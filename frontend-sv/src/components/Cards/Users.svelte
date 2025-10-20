<script lang="ts">
	import type { UserSimpleInfo } from "@types/user";


    let {
        users
    } : {
        users: UserSimpleInfo[]
    } = $props();
</script>


{#if users.length != 0}
{#each users as u}
    <ul>
        <li>{u.name}</li>
        <li>{u.username}</li>
        <li>{u.email}</li>
        <li>{u.net_id}</li>
        <li>{u.ou}</li>
    </ul>
{/each}
{:else}
No Users Found
{/if}