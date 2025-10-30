<script lang="ts">
	import { copyToClip, type CopyState } from "$lib/helper/copy.svelte";
	import type { GroupSimpleInfo } from "@t/group";

    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })

    let {
        groups
    } : {
        groups: GroupSimpleInfo[]
    } = $props();

</script>


<section>
    <input/>
    {#each groups as group, idx }
        <ul style="--delay: {Math.min(idx * 50, 2000)}ms">
            <button onclick={() => copyToClip(group.name, copyState)} class="name"><li>{group.name === copyState.copied ? "Copied" : group.name}</li></button>
            {#if group.description != ""}
                <button onclick={() => copyToClip(group.description, copyState)}>
                    <b>Description: </b> 
                    <li>{group.description === copyState.copied ? "Copied" : group.description}</li>
                </button>
            {/if}
            {#if group.information != ""}
                <button onclick={() => copyToClip(group.information, copyState)}>
                    <b>Information: </b> 
                    <li>{group.information === copyState.copied ? "Copied" : group.information}</li>
                </button>
            {/if}
            {#if group.ou != ""}
                <button onclick={() => copyToClip(group.ou, copyState)}>
                    <b>OU: </b> 
                    <li>{group.ou === copyState.copied ? "Copied" : group.ou}</li>
                </button>
            {/if}
        </ul>
    {/each}

</section>


<style>

    section{ 
        position: relative;
        padding-top: 50px;
    }

    input {
        position: fixed;
        top: 230px;
        left: 50%;
        width: 100px;
        transform: translateX(-50%);
        z-index: 1;
    }

    ul {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        list-style: none;
        animation: slideIn 0.5s forwards var(--delay);
        opacity: 0;
        padding: 1rem 1rem;
        border-radius: 10px;
        background: var(--background-surface);
    }

    button.name {
        font-weight: bold;
        font-size: 18px;
    }

    button {
        background: transparent;
        display: flex;
        flex-direction: row;
        border: none;
        gap: 1rem;
        padding-left: 1rem;
        font-size: 15px;
        color: var(--color-text)
    }
    @keyframes slideIn {
        from {
            opacity: 0;
            transform: translateY(100%);
        }
        to {
            opacity: 1;
            transform: translateY(0px);
        }
    }

</style>