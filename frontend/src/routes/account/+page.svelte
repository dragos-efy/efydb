<svelte:head>
    <title>Account</title>
</svelte:head>

<section>
    {#if !loaded}
    <div class="spin"></div>
    {:else if userInfo}
    <div id="user-actions">
        <i efy_icon="user" id="avatar" title="Avatar" />
        <h2>{userInfo.name}</h2>
        <p>{userInfo.bio}</p>
        <div class="user-buttons">
            <button on:click={logout}>Logout</button>
            <button on:click={deleteAccount}>Delete Account</button>
        </div>
    </div>
    {:else}
    <div id="signUp">
        <h1>Account</h1>
        <input type="text" bind:value={username} placeholder="Username">
        <input type="password" bind:value={password} placeholder="Password">
        <input type="text" bind:value={bio} placeholder="Bio">
        <div class="sign-buttons">
            <button on:click={signIn}>Sign In</button>
            <button on:click={signUp}>Sign Up</button>
        </div>
    </div>
    {/if}
    {#if dialogShown}
    <Modal title={dialogTitle} message={dialogMsg} onClose={onDialogClose} onConfirm={onDialogConfirm} />
    {/if}
</section>

<script type="ts">
	import fetchJson from "$lib/fetchjs";
	import { getToken, logoutUser, setUserInfo } from "$lib/token";
	import Modal from "../../components/Modal.svelte";
	import { onMount } from "svelte";

    let loaded = false;
    let userInfo: any;

    let username: string;
    let password: string;
    let bio: string;

    let dialogShown: boolean = false;
    let dialogTitle: string;
    let dialogMsg: string;
    const onDialogClose = () => {
        dialogShown = false;
    }
    let onDialogConfirm: (() => void) | null = null;

    const showDialog = (title: string, message: string, onConfirm: (() => void) | null = null) => {
        dialogTitle = title;
        dialogMsg = message;
        onDialogConfirm = onConfirm;
        dialogShown = true;
    }

    onMount(async () => {
        let token = getToken();
        if (token) {
            userInfo = await fetchJson("/users/account", {});
        }
        loaded = true;
    })

    const getReqJson = () => {
        return {
            method: "POST",
            body: JSON.stringify({
                name: username,
                password: password,
                bio: bio,
            })
        }
    }

    const signUp = async () => {
        if (!password || !username) {
            showDialog("Invalid request", "Please fill out all the info!");
            return
        }

        const response = await fetchJson("/users/register", getReqJson())

        if (response.message) {
            showDialog("Server error", response.message);
            return;
        }

        userInfo = response;
        setUserInfo(response);
    }

    const signIn = async () => {
        if (!password || !username) {
            showDialog("Invalid request", "Please fill out all the info!");
            return
        }

        const response = await fetchJson("/users/login", getReqJson())

        if (response.message) {
            showDialog("Server error", response.message);
            return;
        }

        userInfo = response;
        setUserInfo(response);
    }

    const logout = () => {
        userInfo = null,
        logoutUser();
    }

    const deleteAccount = async () => {
        showDialog("Delete account", "Are you sure? This can't be undone!", async () => {
            await fetchJson("/users/delete", {
                method: "DELETE"
            });
            logout();
        });
    }
</script>

<style>
    #user-actions {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
    }

    #user-actions > div {
        margin-top: 30rem;
        display: flex;
        flex-wrap: wrap;
        gap: var(--efy_gap0);
    }

    :is(.user-buttons, .sign-buttons) button {
        margin: 0;
    }

    .sign-buttons {
        display: flex;
        gap: var(--efy_gap0);
    }

    #avatar {
        padding: 20rem;
        line-height: 1;
        font-size: 80rem;
        border-radius: calc(var(--efy_radius2) * 2);
        border: 8rem solid var(--efy_text);
        margin: 0 0 25rem 0;
    }

    #avatar:before {
        color: var(--efy_text);
    }

    #signUp {
        max-width: calc(var(--efy_100vh) * 0.7);
    }
</style>