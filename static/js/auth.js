document.body.addEventListener("htmx:afterRequest", (e) => {
    /** @type {XMLHttpRequest} */
    const xhr = e.detail.xhr;
    let err = "Unknown error";
    try {
        const data = JSON.parse(xhr.response);
        console.debug(data);

        err = data.error;
    } catch {}

    window.toastNotification.add({
        type: "error",
        text: err,
        timeout: 5000,
    });
});

async function initializePasskeys() {
    if (
        !window.PublicKeyCredential ||
        !PublicKeyCredential.isConditionalMediationAvailable
    ) {
        return;
    }

    if (!(await PublicKeyCredential.isConditionalMediationAvailable())) return;

    document.getElementById("or").style.display = "flex";
    document.getElementById("passkey").style.display = "flex";

    document
        .getElementById("passkey-button")
        .addEventListener("click", async () => {
            /** @type {PublicKeyCredentialRequestOptions} */
            const publicKeyCredentialCreationOptions = await fetch(
                "/api/auth/passkey/begin",
            )
                .then((res) => res.json())
                .then(PublicKeyCredential.parseRequestOptionsFromJSON)
                .catch((err) => console.error(err));

            const assertion = await navigator.credentials.get({
                publicKey: publicKeyCredentialCreationOptions,
            });
            assertion;

            const response = await fetch("/api/auth/passkey/finish");
            response;
        });
}

initializePasskeys();
