<script lang="ts">

    // === LIBRARIES ===
    import Prism from "prismjs"
    import 'prismjs/components/prism-json';
    import 'prismjs/components/prism-http';

    // === COMPONENTS ===
    import Code from "$lib/components/docs/Code.svelte";
	import Description from "$lib/components/docs/Description.svelte";
	import Response from "$lib/components/docs/Response.svelte";

    // === CONSTANTS ===
    import { endpoint, information, response, responses } from "./details";


</script>



<main>
    <section class="all">
        <h1 class="title">Verify</h1>
        <Description header={"Description"} information={information}/>
        <Code header={"Endpoint"} code={endpoint} prism={Prism.languages.http} language={"http"} />
        <Response responses={responses}/>
        <Code header={"Response"} code={response} prism={Prism.languages.http} language={"http"} />
    </section>

</main>


<style>

    h1.title {
        text-align: center;
    }

    main {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;

    }

    section.all {
        width: 50%;
        display: flex;
        flex-direction: column;
        gap: 10px;
    }


</style>