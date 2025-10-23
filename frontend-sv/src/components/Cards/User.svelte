<script lang="ts">
	import type { UserSimpleInfo } from "@t/user";
    import { copyToClip, type CopyState } from '$lib/helper/copy.svelte';
    import outIcon from '$lib/assets/right-arrow-orange.png';

    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })

    let {
        user,
        idx
    } : {
        user: UserSimpleInfo
        idx: number
    } = $props();

</script>


<ul class:disabled={user.ou.toLowerCase().includes("disabled")} style="--delay: {Math.min(idx * 50, 2000)}ms">
    <a href={`/user/${user.username}`}> <img src={outIcon} alt=""></a>

    {#each Object.entries(user) as key}
        {#if key[1] != ""}
            <li class={key[0]}> 
                <button
                type="button"
                onclick={() => copyToClip(key[1], copyState)}>

                    {#if key[0] != "name"}<b>{key[0].toUpperCase()}:</b>{/if}
                    {#if copyState.copied == key[1]} 
                        Copied
                    {:else}
                        {key[1]}
                    {/if}

                </button>
            </li>
        {/if}
    {/each}
</ul>


<style>
ul.disabled {
        color: var(--disabled)
    }

    img {
        width: 20px;
    }

    
    button {
        background: none;
        border: none;
        padding: 0;
        font: inherit;
        color: inherit;
        cursor: pointer;
    }

    button:focus,
    button:active{
        outline: none;
    }


    ul {
        display: flex;
        flex-direction: column;
        word-break: break-all;
        gap: 0.5rem;
        border-radius: 10px;
        padding: 1.5rem;
        padding-right: 4rem;
        box-sizing: border-box;
        background: var(--background-surface);
        color: var(--text);
        opacity: 0;
        overflow-x: hidden;
        transform: translateY(20px);
        animation: slideIn 0.2s ease-out forwards;
        animation-delay: var(--delay);
        margin: 0;
        list-style: none;
    }

    a {
        position: absolute;
        display: flex;
        justify-content: center;
        align-items: center;
        top: 0;
        right: 0;
        height: 100%;
        width: 50px;
    }

    a:hover {
        background-color: var(--color-primary-hover-opacity-20);
    }

    li.name {
        font-weight: bold;
        font-size: 18px;
        margin-bottom: 1rem;
    }

    li {
        text-align: left;
    }

    button {
        background: none;
        border: none;
        padding: 0;
        font: inherit;
        color: inherit;
        cursor: pointer;
    }

    button:active,
    button:focus { 
        border:none;
        outline: none;
    }


    
    @keyframes slideIn {
        from {
            opacity: 0;
            transform: translateY(20px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }
</style>