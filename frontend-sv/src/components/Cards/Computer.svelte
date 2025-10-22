<script lang="ts">
    import type { ComputerSimpleInfo } from "@t/computer";
    import { copyToClip } from '$lib/helper/copy';
    import outIcon from '$lib/assets/right-arrow-orange.png';

    let copied: string = $state("");
    let copiedMessage: string = $derived.by(() => {
        return copied + " has been copied"
    }); 
    let timeout: ReturnType<typeof setTimeout> | null = $state(null);


    let {
        computer,
        idx
    } : {
        computer: ComputerSimpleInfo
        idx: number
    } = $props()

    function copy(text: string) {
        if (timeout != null) {
            clearTimeout(timeout)
        }
        copyToClip(text)
        copied = text

        timeout = setTimeout(() => {
            copied = ""
            timeout = null
        }, 2000)
    }

</script>

<ul class:disabled={computer.ou.toLowerCase().includes("disabled")} style="--delay: {Math.min(idx * 50, 2000)}ms">
    <a href={`/computer/${computer.name}`}> <img src={outIcon} alt=""></a>
    <li class="title">
        <button 
        type="button" 
        class="copy-button" 
        onclick={() => copy(computer.name)}>
            {computer.name}
        </button>
    </li>
    <li>                
        <button 
        type="button" 
        class="copy-button" 
        onclick={() => copy(computer.ou)}>
            OU: {computer.ou}
        </button>
    </li>
    <li>                
        <button 
        type="button" 
        class="copy-button" 
        onclick={() => copy(computer.operating_system)}>
            OS: {computer.operating_system}
        </button>
    </li>
    {#if copied != ""}
        {copiedMessage}
    {/if}
</ul>




<style>
    ul.disabled {
        color: red
    }

    img {
        width: 20px;
    }

    ul {
        display: flex;
        flex-direction: column;
        gap: 0.5rem;
        border-radius: 10px;
        padding: 1.5rem;
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

    li.title {
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
