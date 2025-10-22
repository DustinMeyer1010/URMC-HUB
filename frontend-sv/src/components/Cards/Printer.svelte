<script lang="ts">
    import type { PrinterSimpleInfo } from "@t/printer";

    let {
        printer,
        idx
    } : {
        printer: PrinterSimpleInfo
        idx: number
    } = $props()

</script>


    <ul style="--delay: {Math.min(idx * 50, 2000)}ms">
        <li>{printer.server}</li>
        <li>{printer.queue}</li>
        <li>{printer.model}</li>
        <li>{printer.ip}</li>
        <li>{printer.print_processor}</li>
        <li>{printer.location}</li>
        <li>{printer.notes}</li>
    </ul>




<style>
    ul {
        border-radius: 10px;
        padding: 1.5rem;
        box-sizing: border-box;
        background: var(--background-surface);
        color: var(--text);
        opacity: 0;
        transform: translateY(20px);
        animation: slideIn 0.2s ease-out forwards;
        animation-delay: var(--delay);
        margin: 0;
        list-style: none;

    }

    a {
        position: absolute;
        top: 0.3rem;
        left: 0.35rem;
    }

    li.title {
        margin-top: 0.5rem;
        font-weight: bold;
        font-size: 18px;
        margin-bottom: 1rem;
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


