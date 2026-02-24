<!-- TODO: Add message for results when someone is removed from the group -->

<script lang="ts">
	import UserCard from '../cards/User.svelte';
    import { RemoveStateClass } from './states/RemoveState.svelte';
    import Remove from "../user/Remove.svelte"
	import Search from '../search/Search.svelte';
	import { fly } from 'svelte/transition';

    let { group } : { group: string } = $props()

    let RemoveState: RemoveStateClass = new RemoveStateClass()
    $inspect(RemoveState.Results)
</script>


{#if RemoveState.Results.length != 0}
    {#each RemoveState.Results as Result }
        <div id="message" class:success={Result.successful} class:error={!Result.successful} in:fly={{y: -100}} out:fly={{y: -100}}>
            <span>{Result.group}</span>
            <span>{Result.message}</span>
        </div>
    {/each}
{/if}

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
<div id="search">
    <Search bind:searchValue={RemoveState.SearchValue} bind:loading={RemoveState.Loading} search={RemoveState.Search}></Search>
</div>
{/snippet}


<style>
    section {
        display: flex;
        flex-direction: column;
        gap: 10px;
    }

    div#search {
        position: fixed;
        bottom: 50px;
        left: 50%;
        transform: translateX(-50%);
    }

    
    div#message {
        position: absolute;
        display: flex;
        flex-direction: column;
        width: 50%;
        left: 50%;
        top: 30px;
        transform: translateX(-50%);
        padding: 20px;
        z-index: 100;
        border-radius: 10px;
        background: var(--color-bg-opacity-30);
        backdrop-filter: blur(20px);
    }

    div.success {
        border: 2px solid green;
    }

    div.error {
        border: 2px solid red;
    }


</style>
