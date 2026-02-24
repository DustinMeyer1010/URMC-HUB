<!--
    TODO: Needs to display users in a clean format to be added to the group
-->

<script lang="ts">
	import UserCard from "@components/Cards/User.svelte";
	import Confirm from "@components/User/Confirm.svelte";
	import { AddStateClass } from "./AddState.svelte";
	import { fly } from "svelte/transition";

    let { group } : { group: string } = $props()

    let AddState: AddStateClass = new AddStateClass()


</script>


{#if AddState.Results.length != 0}
        {#each AddState.Results as Result }
            <div class:success={Result.successful} class:error={!Result.successful} in:fly={{y: -100}} out:fly={{y: -100}}>
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



<form action="" onsubmit={AddState.Search}>
    <input type="text" placeholder="Search for user to Add" bind:value={AddState.SearchValue}>
    <button type="submit">Search</button>
</form>


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

    div {
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

