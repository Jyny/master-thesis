import abstractview from "./abstractview.js";

export default class extends abstractview {
    constructor(params) {
        super(params);
        this.setTitle("Seesion Owner");
    }

    async getHtml() {
        return `
            <h1>Seesion Owner</h1>
            <p>
                If you want to Unseal the Session Record <br/> Click the button below to get Unseal Challenge.
            </p>
            <div class="input-group mt-5 mb-3">
                <div class="input-group-prepend">
                    <span class="input-group-text">Session ID</span>
                </div>
                <input type="text" class="form-control" value="${this.params.session_id}" readonly>
            </div>
            <div class="input-group mb-3">
                <div class="input-group-prepend">
                    <span class="input-group-text">Owner ID</span>
                </div>
                <input type="text" class="form-control" value="${this.params.owner_id}" readonly>
            </div>
            <div class="accordion mb-3">
                <div class="accordion-item">
                    <div class="accordion-header" id="acchead">
                        <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#accbody">Owner Key</button>
                    </div>
                    <div id="accbody" class="accordion-collapse collapse">
                        <div class="accordion-body">
                            <code>${this.params.owner_key}</code>
                        </div>
                    </div>
                </div>
            </div>
            <p>
                <a class="btn btn-lg btn-secondary">Unseal</a>
            </p>
        `;
    }

    async listener() {
    }
}