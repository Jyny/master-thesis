import abstractview from "./abstractview.js";

export default class extends abstractview {
    constructor(params) {
        super(params);
        this.setTitle("Welcome");
    }

    async getHtml() {
        return `
            <h1>Meeting Box</h1>
            <p>
                Scan QR Code on the Meeing Box.
            </p>
        `;
    }
}