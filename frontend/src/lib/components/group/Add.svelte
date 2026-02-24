<!--
    TODO: Needs to display users in a clean format to be added to the group
-->

<script lang="ts">
	import UserCard from "../cards/User.svelte";
	import Search from "../search/Search.svelte";
	import Confirm from "../user/Confirm.svelte";
	import { AddStateClass } from "./states/AddState.svelte";
	import { fly } from "svelte/transition";

    let { group } : { group: string } = $props()

    let AddState: AddStateClass = new AddStateClass()


</script>


{#if AddState.Results.length != 0}
        {#each AddState.Results as Result }
            <div id="message" class:success={Result.successful} class:error={!Result.successful} in:fly={{y: -100}} out:fly={{y: -100}}>
                <span>{Result.group}</span>
                <span>{Result.message}</span>
            </div>
        {/each}

{/if}

<section>
{#each AddState.Users as user, idx }
    <UserCard item={user} {idx}>
        {#if !user.ou.toLowerCase().includes("disabled") && !user.ou.toLocaleLowerCase().includes("offboarded") && user.username}
            <Confirm name={group} username={user.username} title={"Add Group?"} value={"Add"} action={AddState.AddToGroup}/>
        {/if}
    </UserCard>
{/each}
</section>


<div id="search">
    <Search bind:searchValue={AddState.SearchValue} bind:loading={AddState.Loading} search={AddState.Search}></Search>
</div>

<style>
    section {
        display: flex;
        flex-direction: column;
        gap: 10px;
        width: 90%;
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

