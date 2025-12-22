<!--
    TODO: Needs to display users in a clean format to be added to the group
-->

<script lang="ts">
	import User from "@components/Cards/User.svelte";
	import Confirm from "@components/User/Confirm.svelte";
	import type { UserSimpleInfo } from "@t/user";


    let { group } : { group: string } = $props()


    let searchValue: string = $state("")
    let users: UserSimpleInfo[] = $state([])

    const onsubmit = async (e: SubmitEvent) => {
        e.preventDefault()
        await fetch(`http://localhost:8000/api/search/users/${searchValue}`)
        .then(async (res) => users = await res.json())

    }

    

    const test = (username: string, name: string) => {
        return
    }



</script>


<section>
{#each users as user, idx }
    <User item={user} {idx}>
        {#if !user.ou.toLowerCase().includes("disabled") && !user.ou.toLocaleLowerCase().includes("offboarded") && user.username}
            <Confirm name={group} username={user.username} title={"Add Group?"} value={"Add"} action={test}/>
        {/if}
    </User>
{/each}
</section>



<form action="" onsubmit={onsubmit}>
    <input type="text" placeholder="Search for user to Add" bind:value={searchValue}>
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

