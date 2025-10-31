<script lang="ts">




const addGroup = async () => {
    const res = await fetch("http://localhost:8000/api/user/group/add", {
        method: "POST",
        mode: "cors",
        body: JSON.stringify({
            users: ["dmeyer20"],
            groups: ["test"]
        })
    })

    const data = await res.json()

    console.log(data)
}


</script>

<button onclick={addGroup}>Click me</button>


<style>

</style>