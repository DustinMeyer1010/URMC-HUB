
<div class="loading">
    <span>Generating Excel File</span>
    <div style="--delay: 0ms"></div>
    <div style="--delay: 100ms"></div>
    <div style="--delay: 200ms"></div>

</div>


<style>
    div.loading {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        display: flex;
        gap: 1rem;
        justify-content: center;
        align-items: center;
    }

    div.loading div {
        background: var(--color-text);
        width: 10px;
        height: 10px;
        border-radius: 10px;
        opacity: 0;
        animation: blink 2s infinite ease;
        animation-delay: var(--delay);
    }

    @keyframes blink {
        0% {
            opacity: 0%;
        }
        50% {
            opacity: 100%;
        }
        75% {
            opacity: 100%;
        }
        100% {
            opacity: 0%;
        }
    }
</style>