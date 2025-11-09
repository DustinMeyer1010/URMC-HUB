<script lang="ts">
    import { copyToClip, type CopyState } from '$lib/helper/copy.svelte';
    import goToIcon from '$lib/assets/double-left-arrow-primary.png';
    import copyAllIcon from "$lib/assets/copy-color-text.png"
	import type { DriveSimpleInfo } from "@t/drive";
	import { blur } from 'svelte/transition';

    let copyState: CopyState = $state({
        copied: "",
        timeout: null
    })


    let searchValue: string = $state("")

    let {
        drive,
        idx
    } : {
        drive: DriveSimpleInfo
        idx: number
    } = $props()

    let groups: string[] = $derived.by(() => {
        if (!searchValue) {
            return drive.groups
        }
        return drive.groups.filter((group) => group.toUpperCase().includes(searchValue.toUpperCase()))
    })


    let santizedDrive: string = $derived.by(() => {
        return encodeURIComponent(drive.drive)
    })

    $inspect(copyState)
    const allCopyText: string = `Name: ${drive.drive}\nLocal_Path: ${drive.local_path}\nGroups:\n${drive.groups.join("\n")}`;

</script>

<ul style="--delay: {Math.min(idx * 50, 2000)}ms">
    {#if copyState.copied != allCopyText}
        <button class="copy-all" title="Copy All" onclick={() => copyToClip(allCopyText, copyState)}><img src={copyAllIcon} alt="Copy All"></button>
    {:else}
        <span class="copied-all">ALL COPIED</span>
    {/if}
    <a href={`/drive/${santizedDrive}`}> <img src={goToIcon} alt=""></a>
    <li class="name">
        <button onclick={() => copyToClip(drive.drive, copyState)}>
            {copyState.copied === drive.drive ? "Copied" : drive.drive}
        </button>
    </li>
    <li>
        <button onclick={() => copyToClip(drive.local_path, copyState)}>
            LOCAL_PATH: {copyState.copied === drive.local_path ? "Copied" : drive.local_path}
        </button>
    </li>
    <li>
        GROUPS: 
        {#if drive.groups.length >= 10}
            <input placeholder="Search For Group" bind:value={searchValue}/> 
        {/if}
        <div >
            {#each groups as group,i }
                <button class:active={copyState.copied === group} onclick={() => copyToClip(group, copyState)} out:blur={{ duration: 200}} class="group-button" style="--delay: {Math.min(i * 20, 2000)}ms">{copyState.copied === group ? `Copied` : group}</button>
            {/each}
        </div>
    </li>
</ul>

<style >

    input {
        background: var(--color-bg-opacity-30);
        border: none;
        color: var(--color-text);
        padding: 0.2rem;
        border-radius: 5px;
    }


    button.active {
        color: var(--color-primary-focus);

    }

    input:focus {
        outline: none;
    }

    div {
        display: flex;
        gap: 0.5rem;
        transition: 0.5s ease;
        margin-top: 0.5rem;
        padding: 0.5rem;
        flex-wrap: wrap;
        max-height: 150px;
        overflow-y: auto;
        overflow-x: hidden;
        scrollbar-width: thin; 
        scrollbar-color: rgba(255,255,255,0.3) transparent;
    }

    img {
        width: 25px;
    }



    button.group-button {
        padding: 0.5rem 0.6rem;
        background: var(--color-bg-opacity-50);
        border-radius: 50px;
        animation: 0.3s fillin forwards;
        animation-delay: var(--delay);
        opacity: 0;
        max-width: 49%;
        flex: 1 0 49%;
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
        padding-bottom: 2rem;
        padding-left: 1.5rem;
        padding-right: 3rem;
        box-sizing: border-box;
        background: var(--color-surface);
        color: var(--color-text);
        opacity: 0;
        overflow-x: hidden;
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

    @media (max-width: 400px) {
        ul {
            padding-right: 1rem;

        }
    }

    @media (max-width: 800px) {

        div {
            overflow-y: scroll;
        }
    }

    @keyframes fillin {
        from {
            opacity: 0;
            transform: translateY(100%);
        }
        to {
            opacity: 1;
            transform: translate(0,0);
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

