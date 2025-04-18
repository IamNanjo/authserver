Object.defineProperty(window, "globalTheme", {
    get() {
        return localStorage.getItem("theme") || "mocha";
    },
    set(value) {
        localStorage.setItem("theme", value);
        document.body.className = `theme-${value}`;
    },
});

window.addEventListener(
    "DOMContentLoaded",
    () => (window.globalTheme = window.globalTheme),
);
