/** @type Partial<CSSStyleDeclaration> */
const containerStyle = {
    position: "absolute",
    top: "0.5em",
    right: "0.5em",
    display: "flex",
    flexDirection: "column",
    gap: "0.5em",
    width: "min(80%, 14em)",
    height: "min(80%, 25em)",
    zIndex: "1",
    overflow: "auto",
    scrollbarWidth: "none",
};

/** @type Partial<CSSStyleDeclaration> */
const notificationBaseStyle = {
    position: "relative",
    color: "var(--crust)",
    width: "100%",
    padding: "0.5em",
    borderRadius: "4px",
    fontWeight: "500",
    transform: "translateX(150%)",
    userSelect: "none",
};

/** @type Record<ToastNotificationType, Partial<CSSStyleDeclaration>> */
const notificationStyles = {
    ok: {
        ...notificationBaseStyle,
        backgroundColor: "var(--green)",
    },
    info: {
        ...notificationBaseStyle,
        backgroundColor: "var(--blue)",
    },
    warn: {
        ...notificationBaseStyle,
        backgroundColor: "var(--peach)",
    },
    error: {
        ...notificationBaseStyle,
        backgroundColor: "var(--red)",
    },
};

window.toastNotification = {
    config: { animationSpeed: 200 },
    container: document.createElement("div"),
    add(notification) {
        const el = document.createElement("div");
        el.innerText = notification.text;
        Object.assign(el.style, notificationStyles[notification.type]);
        el.style.transition = `transform ${this.config.animationSpeed}ms ease-in-out`;

        this.container.prepend(el);

        setTimeout(() => {
            el.style.transform = "translateX(0%)";
        }, 0);

        el.addEventListener("click", () => {
            this.remove({ element: el });
        });

        if ("timeout" in notification && !Number.isNaN(notification.timeout)) {
            setTimeout(() => {
                this.remove({ element: el });
            }, notification.timeout + this.config.animationSpeed);
        }
    },
    remove({ element }) {
        element.style.transform = "translateX(150%)";
        setTimeout(() => element.remove(), this.config.animationSpeed);
    },
};

Object.assign(window.toastNotification.container.style, containerStyle);

window.addEventListener("DOMContentLoaded", () => {
    document.body.appendChild(window.toastNotification.container);
});
