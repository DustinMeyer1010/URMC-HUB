<script lang="ts">

</script>




<div>
    {#each {length: 7} as _,x}
        <span style={`--delay: ${250*x}ms`}></span>
    {/each}
</div>



<style>

    div {
        padding: 1rem;
        display: flex;
        flex-direction: column;
        gap: 1rem;
    }

    span {
        background-color: var(--color-surface);
        width: 100%;
        height: 100px;
        border-radius: 10px;
        animation: 3s flash infinite;
        animation-delay: var(--delay);
    }

    
    @keyframes flash {
        from {
            background-color: var(--color-bg-flash);
        }
        to {
            background-color: var(--color-surface);
        }
    }

</style>