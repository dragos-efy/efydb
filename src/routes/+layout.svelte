<h1>Welcome to EfyDb</h1>
<p>Visit <a href="https://efy.ooo">efy.ooo</a> to read the documentation</p>
<slot />
<script>
    import { onMount } from "svelte";
    onMount(() => {
        import ("../../efy/efy.js");
    })
</script>