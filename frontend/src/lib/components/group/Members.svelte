<!-- 
 Refactor: Fix how long page takes to geneerate
    ? Should I send a request per page of 1500 users
    ? Agent can press a button that will load all the members
    ? Generate the first 1500 and then press a button to generate more (Handles groups with smaller than 1500 members to always show members)
 TODO: Add a button the will generate a excel file with all the memebers in it with
    - Username
    - Name
    - Email
    - Department   
-->

<script lang="ts">
	import User from "@components/Cards/User.svelte";
	import { MembersStateClass } from "./MembersState.svelte";
	import { onMount } from "svelte";
	import MembersLoading from "@components/Loading-Animations/MembersLoading.svelte";

    let {
        group
    } : {
        group: string
    } = $props()

    let MembersState: MembersStateClass = new MembersStateClass(group)

    onMount(async () => {
        await MembersState.GetMembers()
    })


</script>

<section>
    {#if MembersState.Loading}
        <MembersLoading/>
    {:else}
        <div class="member-filtering">
            <input type="text" bind:value={MembersState.Filter} placeholder="Search for member">
            <div class="paging-buttons">
                <button onclick={MembersState.PrevPage}>-</button>
                <button onclick={MembersState.NextPage}>+</button>
                <span>Page: {Math.ceil(MembersState.Page / 10)} / {Math.ceil(MembersState.MembersLength / 10)}</span>
                <span>Total Members: {MembersState.MembersLength}</span>
            </div>
        </div>
        <div class="users">
            {#each MembersState.PagedMembers as member, idx}
                <User item={member} idx={idx}/>
            {/each}
        </div>
    {/if}
</section>


<style>
    section {
        display: flex;
        flex-direction: column;
        justify-content: center;
        gap: 10px;
        width: 80%;
        padding-bottom: 100px;
        padding-top: 30px;
    }

    input {
        font-size: 15px;
        padding: 10px;
        background: var(--color-surface);
        border: none;
        color: var(--text-color);
        border-radius: 5px;
        align-self: center;
        flex-grow: 1;
    }

    div.users {
        display: flex;
        flex-direction: column;
        overflow-y: scroll;
        overflow-x: hidden;
        flex-grow: 1;
        gap: 10px;
    }

    input:focus {
        outline: none;
    }

    div.paging-buttons {
        display: flex;
        align-items: center;
        justify-content: center;
        height: 100%;
        gap: 10px;

    }

    div.member-filtering {
        padding: 20px 50px;
        box-sizing: border-box;
        background: var(--color-bg);
        gap: 10px;
        position: fixed;
        width: 100%;
        left: 50%;
        bottom: 0px;
        transform: translateX(-50%);
        z-index: 9;
        justify-content: space-between;
        display: flex;
        align-items: center;
        flex-direction: row;
    }

    button {
        padding: 5px;
        width: 200px;
        background: var(--color-surface);
        border: none;
        font-size: 20px;
        border-radius: 5px;
        color: var(--text-color);
    }

    button:hover {
        background: var(--color-surface-hover);
    }


</style>