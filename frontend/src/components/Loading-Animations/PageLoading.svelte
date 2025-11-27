<script lang="ts">

</script>





<div>
    <ul>
        <li style="--delay: 0.1s; --left: 0px"></li>
        <li style="--delay: 0.3s; --left: 40px"></li>
        <li style="--delay: 0.5s; --left: 80px"></li>
    </ul>
</div>


<style>
    div {
        position: absolute;
        top: 50%;
        left: calc(50% - 40px);
        transform: translate(-50%, -50%);
    }


    ul {
        position: relative;
    }

    li {
        position: absolute;
        width: 10px;
        height: 10px;
        border-radius: 10px;
        list-style: none;
        background: var(--color-text);
        animation: 2s bounce infinite;
        animation-delay: var(--delay);
        top: 0px;
        left: var(--left);
    }


    @keyframes bounce {
        0% {
            transform: translateY(0px);
            width: 10px;
            height: 10px;
        }
        50% {
            transform: translateY(20px);
            width: 25px;
            height: 10px;
        }
        100% {
            transform: translate(0px);
            width: 10px;
            height: 10px;
        }
    }
</style>