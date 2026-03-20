<script lang="ts">
	import type { Snippet } from "svelte";


    let {
        header,
        body,
        idx
    } : {
        header: Snippet,
        body: Snippet,
        idx: number,
    } = $props()

</script>


<div style={`--delay: ${Math.min(idx * 50, 1000)}ms`}>
    <header>
        {@render header()}
    </header>

    {@render body()}
</div>


<style>

    div {
        display: flex;
        flex-direction: column;
        height: fit-content;
        width: 100%;
        background: var(--color-surface);
        border-radius: 10px;
        position: relative;
        opacity: 0;
        transform: translateY(100px);
        animation: slideIn 0.3s ease-out forwards;
        animation-delay: var(--delay);
    }

    header {
        display: flex;
        gap: 0.5rem;
        align-items: center;
        width: 100%;
        background: var(--color-surface-hover);
        padding: 5px;
        padding-left: 10px;
        min-height: 40px;

    }

    
    @keyframes slideIn {
        from {
            opacity: 0;
            transform: translateY(100px);
        }
        to {
            opacity: 1;
            transform: translateY(0);
        }
    }



</style>