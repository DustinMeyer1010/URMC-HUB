<script lang="ts">
    // === LIBRARIES ===
    import Prism from "prismjs"
	import 'prismjs/components/prism-json';
    import 'prismjs/components/prism-http';
	import 'prismjs/themes/prism-tomorrow.css';

    // === COMPONENTS ===
	import Code from "@components/Docs/Code.svelte";
	import Object from "@components/Docs/Object.svelte";
	import Description from "@components/Docs/Description.svelte";
	import Response from "@components/Docs/Response.svelte";
    
    // === CONSTANTS ===
    import { endpoint, information, objects, response, request, responses } from "./details";

</script>

<main>
    <section class="all">
        <h1 class="title">Login</h1>
        <Description header={"Description"} information={information}/>
        <Code header={"Endpoint"} code={endpoint} prism={Prism.languages.http} language={"http"} />
        <Object objects={objects} />
        <Code header={"Request Body"} code={request} prism={Prism.languages.json} language={"json"} />
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