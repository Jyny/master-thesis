import abstractview from "./abstractview.js";

export default class extends abstractview {
    constructor(params) {
        super(params);
        this.setTitle("Register");
    }

    async getHtml() {
        return `
            <h1>Welcome to New Session</h1>
            <p>
                If you want to register as an Session Owner. <br/> Please click the button below.
            </p>
            <div class="input-group mb-5">
                <div class="input-group-prepend">
                    <span class="input-group-text">Session ID</span>
                </div>
                <input type="text" class="form-control" disabled="disabled">
            </div>
            <p>
                <a class="btn btn-lg btn-secondary">Register As Owner</a>
            </p>
        `;
    }
}