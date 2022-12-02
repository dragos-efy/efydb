<section>
    <h3>Create Theme</h3>
    <input type="text" bind:value={title} placeholder="Title">
    <input type="text" bind:value={description} placeholder="Description">
    <div id="buttons">
        <div id="files">
            <input type="file" id="screenshot" accept="image/*" bind:this={screenshot}>
            <button><label for="screenshot">Screenshot</label></button>
            <input type="file" id="config" accept="application/json" bind:this={config}>
            <button><label for="config">Config</label></button>
            <input type="file" id="themeConfig" accept="application/json" bind:this={themeConfig}>
            <button><label for="themeConfig">Theme config</label></button>
        </div>
        <button on:click={create}>Create</button>
    </div>
</section>

<script lang="ts">
	import { fetchFormJson }from "$lib/fetchjs";

    let title: string;
    let description: string;
    let screenshot: any;
    let config: any;
    let themeConfig: any;

    const create = async () => {
        if (!title || !description || !screenshot.files[0] || !config.files[0]) {
            alert("Please provide all the info!");
            return;
        }
        
        let formData = new FormData();
        formData.append("data", JSON.stringify({
            title: title,
            description: description,
        }));
        formData.append("screenshot", screenshot.files[0]);
        formData.append("config", config.files[0]);
        if (themeConfig.files.length > 0) formData.append("themeConfig", themeConfig.files[0]);

        await fetchFormJson("/themes/create", {
            method: "POST",
            body: formData,
        })
    }
</script>

<style>
    h3 {
        width: 100%;
    }

    #buttons {
        display: flex;
        width: 100%;
        justify-content: space-between;
    }

    #files {
        display: flex;
    }

    #files button {
        margin-right: 5rem;
        cursor: pointer;
    }

    input[type=file] {
        display: none;
    }
</style>