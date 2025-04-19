type ToastNotificationType = "ok" | "info" | "warn" | "error";

interface Window {
    /** Automatically updates localStorage when changed */
    globalTheme: "latte" | "frappe" | "macchiato" | "mocha" | "black";

    /** Notifications must be manually closed if no timeout is provided */
    toastNotification: {
        config: {
            /** Animation speed in milliseconds */
            animationSpeed: number;
        };
        container: HTMLDivElement;
        add: ({
            type,
            text,
            timeout,
        }: {
            type: ToastNotificationType;
            text: string;
            timeout?: number;
        }) => void;
        remove: ({ element }: { element: HTMLDivElement }) => void;
    };
}
