/** @type {HTMLInputElement} */
const emailInput = document.getElementById("email");

/** @type {HTMLInputElement} */
const usernameInput = document.getElementById("username");

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
            const email = emailInput.value;
            const username = usernameInput.value;

            if (!email && !username) {
                usernameInput.reportValidity();
                return;
            }

            const userExistsRes = await fetch(
                `/api/user/exists?email=${email}&username=${username}`
            );

            if (userExistsRes.status !== 200) {
                /** @type {{reason?: "email" | "username"; error?: string;}} */
                const data = await userExistsRes.json();

                switch (data.reason) {
                    case "email":
                        emailInput.setCustomValidity(data.error);
                        emailInput.reportValidity();
                        break;
                    case "username":
                        usernameInput.setCustomValidity(data.error);
                        usernameInput.reportValidity();
                        break;
                }

                return;
            }

            /** @type {PublicKeyCredentialCreationOptions} */
            const publicKeyCredentialCreationOptions = await fetch(
                "/api/auth/passkey-begin-register",
                {
                    method: "POST",
                    body: JSON.stringify({ email, username }),
                }
            )
                .then((res) => res.json())
                .catch((err) => console.error(err.message));

            const assertion = await navigator.credentials.get({
                publicKey: publicKeyCredentialCreationOptions,
            });

            if (!assertion) {
                return;
            }

            const response = await fetch("/api/auth/passkey");
            response;
        });
}

initializePasskeys();
