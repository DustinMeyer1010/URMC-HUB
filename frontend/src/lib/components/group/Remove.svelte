<!-- TODO: Add message for results when someone is removed from the group -->

<script lang="ts">
	import UserCard from '@components/Cards/User.svelte';
    import { RemoveStateClass } from './RemoveState.svelte';
    import Remove from "@components/User/Remove.svelte";

    let { group } : { group: string } = $props()

    let RemoveState: RemoveStateClass = new RemoveStateClass()
    $inspect(RemoveState.Results)
</script>


<section>
{#each RemoveState.Users as user, idx }
    <UserCard item={user} {idx}>
        {#if !user.ou.toLowerCase().includes("disabled") && !user.ou.toLocaleLowerCase().includes("offboarded") && user.username}
            <Remove username={user.username} group={group} removeGroup={RemoveState.RemoveFromGroup}/>
        {/if}
    </UserCard>
{/each}
</section>

{@render SearchForm()}

<!-- REFACTOR: Fix the way this looks -->
{#snippet SearchForm()}
    <form action="" onsubmit={RemoveState.Search}>
        <input type="text" placeholder="Search for user to Remove" bind:value={RemoveState.SearchValue}>
        <button type="submit">Search</button>
    </form>
{/snippet}


<style>
    section {
        display: flex;
        flex-direction: column;
        gap: 10px;
    }

    input {
        padding: 5px 10px;
        font-size: 12px;
        color: var(--color-text);
        border: 3px solid var(--color-text);
        background: var(--color-bg);
        border-radius: 10px;
        height: 100%;
        max-width: 200px;
        box-sizing: border-box;

    }

    input:focus {
        outline:none
    }


    form {
        display: flex;
        flex-direction: row;
        flex-wrap: nowrap;
        justify-content: center;
        align-items: center;
        gap: 10px;
        height: 40px;
        position: fixed;
        bottom: 30px;
        left: 50%;
        width: 50%;
        transform: translateX(-50%);
    }

    button {
        height: 100%;
        box-sizing: border-box;
        color: var(--color-text);
        border: 3px solid var(--color-text);
        background: var(--color-bg);
        padding: 5px 15px;
        border-radius: 10px;
    }
</style>
