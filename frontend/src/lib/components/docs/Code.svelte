
<script lang="ts">
        import Prism from 'prismjs';


    let {
        header,
        code,
        prism,
        language
    } : {
        header: string
        code: string
        prism: Prism.Grammar
        language: string
    }  = $props()

</script>



<section>
    <header>
        <h1>{header}</h1>
    </header>
    <pre class="code-block">
{@html Prism.highlight(code, prism, language)}
    </pre>
</section>


<style>

    section {
        background: var(--color-surface);
        border-radius: 6px;
        overflow: hidden;
    }

    h1 {
        width: 100%;
        font-size: 15px;
        margin: 0;
        padding: 5px;

        box-sizing: border-box;
        background: var(--color-surface-lighter);
    }

    pre {
        padding-left: 20px;
    }

</style>