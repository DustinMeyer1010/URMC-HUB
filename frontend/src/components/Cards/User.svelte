<script lang="ts">
	import type { UserSimpleInfo } from "@t/user";
    import { copyToClip, type CopyState } from '$lib/helper/copy.svelte';
    import outIcon from '$lib/assets/double-left-arrow-primary.png';
    import copyAllIcon from "$lib/assets/copy-color-text.png"
    import disabledIcon from '$lib/assets/disabled-color-disabled.png'

    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })

    let {
        item,
        idx,
    } : {
        item: UserSimpleInfo
        idx: number
    } = $props();

    const allCopyText: string = `Name: ${item.name}\nUsername: ${item.username}\nNet_ID: ${item.net_id}\nURID: ${item.urid}\nEmail: ${item.email}\nOU: ${item.ou}\n`;
    const disabled = item.ou.toLowerCase().includes("disabled") || item.ou.toLocaleLowerCase().includes("offboarded")
</script>


<ul class:disabled={disabled} style="--delay: {Math.min(idx * 50, 2000)}ms">
    {#if disabled}
        <span class="disabled"><img src={disabledIcon} alt="">Disabled Account</span>
    {/if}
    {#if item.username != ""}
    <a href={`/user/${item.username}`}> <img src={outIcon} alt=""></a>
    {/if}
    {#if copyState.copied != allCopyText}
        <button class="copy-all" title="Copy All" onclick={() => copyToClip(allCopyText, copyState)}><img src={copyAllIcon} alt="Copy All"></button>
    {:else}
        <span class="copied-all">ALL COPIED</span>
    {/if}
    {#each Object.entries(item) as key}
        {#if key[1] != ""}
            <li class={key[0]}>     
                <button
                type="button"
                class="value"
                onclick={() => copyToClip(key[1], copyState)}>
                        {@html key[0] != "name" ? `<b>${key[0].toUpperCase()}:</b>` : ""}
                        {copyState.copied == key[1] ? "Copied" : key[1]}
                </button>
            </li>
        {/if}
    {/each}
</ul>


<style >

    span.disabled {
        display: flex;
        justify-content: center;
        align-items: center;
        position: absolute;
        color: var(--color-ad-disabled);
        right: 1rem;
        bottom: 0.5rem;
        font-weight: bold;
        font-size: 12px;
    }

    span.disabled img {
        width: 20px;
        margin-right: 2px;
    }


    img {
        width: 25px;
    }

    
    button {
        background: none;
        border: none;
        padding: 0;
        font: inherit;
        color: var(--color-text);
        cursor: pointer;
    }

    button:focus,
    button:active{
        outline: none;
    }

    button.value {
        max-width: 80%;
    }

    span.copied-all {
        font-size: 10px;
        position: absolute;
        top: 0.4rem;
        left: 0.4rem;
        opacity: 0.3;
        z-index: -1;
        color: var(--color-text)
    }

    button.copy-all {
        position: absolute;
        top: 0.2rem;
        left: 0.2rem;
        opacity: 0.3;
        z-index: -1;
    }

    button.copy-all img {
        width: 20px;
    }

    button.copy-all:hover img {
        transform: scale(1.1);
    }


    ul {
        display: flex;
        flex-direction: column;
        word-break: break-all;
        position: relative;
        gap: 0.5rem;
        border-radius: 10px;
        padding: 1.5rem;
        padding-left: 1.5rem;
        padding-right: 3rem;
        box-sizing: border-box;
        background: var(--color-surface);
        color: var(--color-text);
        opacity: 0;
        transform: translateY(20px);
        animation: slideIn 0.2s ease-out forwards;
        animation-delay: var(--delay);
        margin: 0;
        list-style: none;
    }

    a {
        opacity: 0.8;
        position: absolute;
        display: flex;
        justify-content: center;
        align-items: center;
        top: 0.3rem;
        right: -0.5rem;
        width: 50px;
    }

    a img {
        transition: var(--transition-fast);
        transform: rotate(130deg);
    }

    a:hover img {
        transform: rotate(130deg) translate(-3px, -1px);
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

    @media (max-width: 1000px) {
        ul.disabled {
            padding-bottom: 3rem;
        }

        span.disabled {
            left: 50%;
            right: 0;
            transform: translateX(-50%);
            text-wrap: nowrap;
        }
        button.value {
            max-width: 100%;
        }
    } 

    @media (max-width: 400px) {

        ul {
            padding-right: 1rem;
        }

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